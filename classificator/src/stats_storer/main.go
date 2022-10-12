package main

import (
	"fmt"
	"stats_storer/proto"

	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/konstellation-io/kre-runners/kre-go"
)

func handlerInit(ctx *kre.HandlerContext) {
	ctx.Logger.Info("[handler init]")
}

func handler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[handler invoked]")

	req := &proto.ClassificatorOutput{}
	err := anypb.UnmarshalTo(data, req, protobuf.UnmarshalOptions{})
	if err != nil {
		return fmt.Errorf("invalid request: %s", err)
	}

	storeMetrics(ctx, req.Category.String())
	return nil
}

// storeMetrics is a helper function to save influxdb metrics for the stats_storer node.
func storeMetrics(ctx *kre.HandlerContext, emailClass string) {
	tags := map[string]string{}

	fields := map[string]interface{}{
		"email_class": emailClass,
	}

	ctx.Measurement.Save("classified_emails", fields, tags)

	fields = map[string]interface{}{
		"called_node": "stats_storer",
	}

	ctx.Measurement.Save("number_of_calls", fields, tags)
}

func main() {
	kre.Start(handlerInit, handler)
}
