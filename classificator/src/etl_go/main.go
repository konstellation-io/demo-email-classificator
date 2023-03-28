package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/konstellation-io/kre-runners/kre-go/v4"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"etl/proto"
)

type Email struct {
	Title        string `csv:"title"`
	Body         string `csv:"body"`
	Author       string `csv:"author"`
	CreationDate string `csv:"date"`
}

func handlerInit(ctx *kre.HandlerContext) {
	ctx.Logger.Info("[handler init]")
}

func handler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[handler invoked]")

	req := &proto.BatchClassificatorRequest{}
	err := anypb.UnmarshalTo(data, req, protobuf.UnmarshalOptions{})
	if err != nil {
		return fmt.Errorf("invalid request: %s", err)
	}

	file, err := os.Open(req.Filename)
	if err != nil {
		return fmt.Errorf("error reading emails file: %w", err)
	}
	defer file.Close()

	var emails []*Email
	if err := gocsv.UnmarshalFile(file, &emails); err != nil {
		return fmt.Errorf("error unmarshaling emails: %w", err)
	}

	batchSize := int(req.BatchSize)

	for i := 0; i < len(emails)/batchSize; i++ {
		batch := emails[i*batchSize : i*batchSize+batchSize]

		parsedBatch, err := gocsv.MarshalBytes(batch)
		if err != nil {
			//ctx.Logger.Errorf("error parsing emails batch: %s", err)
			return fmt.Errorf("error parsing emails batch: %s", err)
		}

		batchKey := fmt.Sprintf("emails-%s-%d", ctx.GetRequestID(), i)

		ctx.Logger.Infof("Sending batch: %s", batchKey)
		fmt.Println(string(parsedBatch))

		err = ctx.ObjectStore.Save(batchKey, parsedBatch)
		if err != nil {
			return fmt.Errorf("error saving emails in object store: %w", err)
		}

		err = ctx.SendOutput(&proto.BatchEtlOutput{
			ObjectKey: batchKey,
		})
		if err != nil {
			return fmt.Errorf("error sending output: %w", err)
		}
	}

	return nil
}

func main() {
	kre.Start(handlerInit, handler)
}
