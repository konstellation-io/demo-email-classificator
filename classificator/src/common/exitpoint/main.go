package main

import (
	"exitpoint/proto"
	"fmt"

	"github.com/konstellation-io/kre-runners/kre-go/v4"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func handlerInit(ctx *kre.HandlerContext) {
	ctx.Logger.Info("[handler init]")
}

// Default handler for the node, every message that do not match a custom handler will be managed by this function.
func defaultHandler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[executing default handler]")

	// Redirect the message to entrypoint
	// This method is useful when you don't need to do any change to the request data,
	// and you just want to send it to the next node.
	ctx.SendAny(data)
	return nil
}

func etlHandler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[executing etl handler]")

	// When the message is an early reply, redirect the message to entrypoint.
	if ctx.IsMessageEarlyReply() {
		ctx.SendAny(data)
	}

	return nil
}

func statsStorerHandler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[executing stats-storer handler]")

	// Unpack the message payload
	statsStorerOutput := &proto.StatsStorerOutput{}
	err := anypb.UnmarshalTo(data, statsStorerOutput, protobuf.UnmarshalOptions{})
	if err != nil {
		return fmt.Errorf("invalid request: %s", err)
	}

	// Do nothing, end of the pipeline
	return nil
}

func main() {
	// Custom handlers definition for multiple subscriptions
	handlers := map[string]kre.Handler{
		"etl":          etlHandler,
		"stats-storer": statsStorerHandler,
	}

	kre.Start(handlerInit, defaultHandler, handlers)
}
