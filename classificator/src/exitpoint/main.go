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

func defaultHandler(ctx *kre.HandlerContext, _ *anypb.Any) error {
	ctx.Logger.Info("[executing default handler]")
	return nil
}

func etlHandler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[executing etl handler]")

	if ctx.IsMessageEarlyReply() {
		ctx.SendAny(data)
	}

	return nil
}

func statsStorerHandler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[executing stats-storer handler]")

	statsStorerOutput := &proto.StatsStorerOutput{}
	err := anypb.UnmarshalTo(data, statsStorerOutput, protobuf.UnmarshalOptions{})
	if err != nil {
		return fmt.Errorf("invalid request: %s", err)
	}
	//ctx.Logger.Info(statsStorerOutput.Message)

	obj, err := ctx.ObjectStore.Get("emails")
	if err != nil {
		ctx.Logger.Errorf("error getting object: %s", err)
	}

	ctx.Logger.Infof("object found: %s", obj)
	ctx.Logger.Infof("object found: %s", string(obj))

	return nil
}

func main() {
	handlers := map[string]kre.Handler{
		"etl":          etlHandler,
		"stats-storer": statsStorerHandler,
	}

	kre.Start(handlerInit, defaultHandler, handlers)
}
