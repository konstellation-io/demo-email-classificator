package main

import (
	"fmt"
	"math/rand"
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

	for i := rand.Intn(10); i != 0; i-- {
		category := pickupRandCategory()
		storeMetrics(ctx, category.String())
	}

	return nil
}

func pickupRandCategory() proto.EmailCategory {
	number := rand.Intn(4)
	switch number {
	case 0:
		return proto.EmailCategory_CATEGORY_REPARATIONS
	case 1:
		return proto.EmailCategory_CATEGORY_ADMINISTRATION
	case 2:
		return proto.EmailCategory_CATEGORY_BILLING
	default:
		return proto.EmailCategory_CATEGORY_SPAM
	}
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
