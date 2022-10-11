package main

import (
	"github.com/konstellation-io/kre-runners/kre-go"
	"google.golang.org/protobuf/types/known/anypb"
)

func handlerInit(ctx *kre.HandlerContext) {
	ctx.Logger.Info("[handler init]")
}

func defaultHandler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[default handler invoked]")
	storeMetrics(ctx)
	return nil
}

func handler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[handler invoked]")

	if ctx.GetRequestMessageType() == kre.MessageType_EARLY_REPLY {
		ctx.SendAny(data)
	}

	storeMetrics(ctx)

	return nil
}

// storeMetrics is a helper function to save influxdb metrics for the exitpoint node.
func storeMetrics(ctx *kre.HandlerContext) {
	tags := map[string]string{}

	fields := map[string]interface{}{
		"called_node": "exitpoint",
	}

	ctx.Measurement.Save("number_of_calls", fields, tags)
}

func main() {
	handlers := map[string]kre.Handler{
		"etl": handler,
	}

	kre.Start(handlerInit, handler, handlers)
}
