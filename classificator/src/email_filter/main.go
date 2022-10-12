package main

import (
	"email_filter/proto"
	"fmt"

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

	err = storeEmail(ctx, req.Email)
	if err != nil {
		ctx.Logger.Errorf("error storing email: %w", err)
	}
	//storeMetrics(ctx)

	return nil
}

func storeEmail(ctx *kre.HandlerContext, email *proto.Email) error {
	return ctx.DB.Save("repair_emails", email)
}

// storeMetrics is a helper function to save influxdb metrics for the email_filter node.
//func storeMetrics(ctx *kre.HandlerContext) {
//	tags := map[string]string{}
//
//	fields := map[string]interface{}{
//		"called_node": "stats_storer",
//	}
//
//	ctx.Measurement.Save("number_of_calls", fields, tags)
//}

func main() {
	kre.Start(handlerInit, handler)
}
