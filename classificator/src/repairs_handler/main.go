package main

import (
	"fmt"
	"repairs_handler/proto"

	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/konstellation-io/kre-runners/kre-go/v4"
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

	err = storeEmail(ctx, req.Email)
	if err != nil {
		ctx.Logger.Errorf("error storing email: %w", err)
	}

	return nil
}

func storeEmail(ctx *kre.HandlerContext, email *proto.Email) error {
	return ctx.DB.Save("repair_emails", email)
}

func main() {
	kre.Start(handlerInit, handler)
}
