package main

import (
	"fmt"
	"math"
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

	// Unpack the message payload
	req := &proto.Request{}
	err := anypb.UnmarshalTo(data, req, protobuf.UnmarshalOptions{})
	if err != nil {
		return fmt.Errorf("invalid request: %s", err)
	}

	// Read emails from CSV file from the filename received in the request
	file, err := os.Open(ctx.Path(req.Filename))
	if err != nil {
		return fmt.Errorf("error reading emails file: %w", err)
	}
	defer file.Close()

	var emails []*Email
	if err := gocsv.UnmarshalFile(file, &emails); err != nil {
		return fmt.Errorf("error unmarshaling emails: %w", err)
	}

	batchSize := int(req.BatchSize)

	// Send early reply to the client
	err = ctx.SendEarlyReply(&proto.Response{
		Message: fmt.Sprintf("Processing %d emails", len(emails)),
	})
	if err != nil {
		return fmt.Errorf("error sending early reply: %w", err)
	}

	amountOfBatches := int(math.Ceil(float64(len(emails)) / float64(batchSize)))

	// Send batches of emails to the next step in the pipeline using the object store
	for i := 0; i < amountOfBatches; i++ {
		lastItemInBatch := i*batchSize + batchSize
		if lastItemInBatch > len(emails) {
			lastItemInBatch = len(emails) - 1
		}

		batchKey := fmt.Sprintf("emails-%s-%d", ctx.GetRequestID(), i)

		err := sendBatch(ctx, emails[i*batchSize:lastItemInBatch], batchKey)
		if err != nil {
			return fmt.Errorf("error sending batch: %w", err)
		}
	}

	return nil
}

func sendBatch(ctx *kre.HandlerContext, batch []*Email, batchKey string) error {
	ctx.Logger.Infof("Processing batch %q", batchKey)

	parsedBatch, err := gocsv.MarshalBytes(batch)
	if err != nil {
		return fmt.Errorf("error parsing emails batch: %w ", err)
	}

	// Save the batch in the object store with the key generated from the request ID and the batch number
	err = ctx.ObjectStore.Save(batchKey, parsedBatch)
	if err != nil {
		return fmt.Errorf("error saving emails in object store: %w", err)
	}

	// Send the batch key to the next step in the workflow
	err = ctx.SendOutput(&proto.EtlOutput{
		EmailsKey: batchKey,
	})
	if err != nil {
		return fmt.Errorf("error sending output: %w", err)
	}

	return nil
}

func main() {
	kre.Start(handlerInit, handler)
}
