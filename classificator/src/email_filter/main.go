package main

import (
	"github.com/konstellation-io/kre-runners/kre-go"
	"google.golang.org/protobuf/types/known/anypb"
)

func handlerInit(ctx *kre.HandlerContext) {
	ctx.Logger.Info("[handler init]")
}

func handler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[handler invoked]")
	return nil
}

// storeMetrics is a helper function to save influxdb metrics for the email_filter node.
func storeMetrics(ctx *kre.HandlerContext, component string, lastTag string) {
	tags := map[string]string{}

	fields := map[string]interface{}{}

	ctx.Measurement.Save("results", fields, tags)
}

func main() {
	kre.Start(handlerInit, handler)
}
