package main

import (
	"fmt"
	"stats_storer/proto"

	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/konstellation-io/kre-runners/kre-go"
)

const messageCounterKey = "messageCounter"

func handlerInit(ctx *kre.HandlerContext) {
	ctx.Logger.Info("[handler init]")
	ctx.Set(messageCounterKey, 0)
}

func handler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[handler invoked]")

	req := &proto.ClassificatorOutput{}
	err := anypb.UnmarshalTo(data, req, protobuf.UnmarshalOptions{})
	if err != nil {
		return fmt.Errorf("invalid request: %s", err)
	}

	storeMetrics(ctx, req.Category.String())
	messageCounter := ctx.Get(messageCounterKey).(int)
	messageCounter = (messageCounter + 1) % 50
	if messageCounter == 0 {
		res := &proto.StatsStorerOutput{Message: "50 messages processed"}
		err := ctx.SendOutput(res)
		if err != nil {
			ctx.Logger.Errorf("error publishing message: %w", err)
		}
	}
	ctx.Set(messageCounterKey, messageCounter)
	return nil
}

// storeMetrics is a helper function to save influxdb metrics for the stats_storer node.
func storeMetrics(ctx *kre.HandlerContext, emailClass string) {
	tags := map[string]string{}

	fields := map[string]interface{}{
		"email_class": emailClass,
	}

	ctx.Measurement.Save("classified_emails", fields, tags)
}

func main() {
	kre.Start(handlerInit, handler)
}
