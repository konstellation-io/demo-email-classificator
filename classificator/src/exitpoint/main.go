package main

import (
	"github.com/konstellation-io/kre-runners/kre-go"
	"google.golang.org/protobuf/types/known/anypb"
)

func handlerInit(ctx *kre.HandlerContext) {
	ctx.Logger.Info("[handler init]")
}

func defaultHandler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[handler invoked]")

	if ctx.GetRequestMessageType() == kre.MessageType_EARLY_REPLY {
		return ctx.SendAny(data)
	}

	return nil
}

func main() {
	handlers := map[string]kre.Handler{
		"etl": defaultHandler,
	}

	kre.Start(handlerInit, defaultHandler, handlers)
}
