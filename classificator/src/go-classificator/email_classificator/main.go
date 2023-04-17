package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/konstellation-io/kre-runners/kre-go/v4"
	protobuf "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"math/rand"

	"etl/proto"
)

type Email struct {
	Title        string `csv:"title"`
	Body         string `csv:"body"`
	Author       string `csv:"author"`
	CreationDate string `csv:"date"`
}

func handlerInit(ctx *kre.HandlerContext) {
	ctx.Logger.Info("[handler init]")
}

func handler(ctx *kre.HandlerContext, data *anypb.Any) error {
	ctx.Logger.Info("[handler invoked]")

	if ctx.IsMessageEarlyReply() {
		return nil
	}

	req := &proto.EtlOutput{}
	err := anypb.UnmarshalTo(data, req, protobuf.UnmarshalOptions{})
	if err != nil {
		return fmt.Errorf("invalid request: %s", err)
	}

	val, err := ctx.Configuration.Get("emails_processed", kre.ProjectScope)
	if err != nil {
		return fmt.Errorf("error getting project config: %w", err)
	}
	ctx.Logger.Info(val)

	val, err = ctx.Configuration.Get("emails_processed", kre.WorkflowScope)
	if err != nil {
		return fmt.Errorf("error getting workflow config: %w", err)
	}
	ctx.Logger.Info(val)

	val, err = ctx.Configuration.Get("emails_processed", kre.NodeScope)
	if err != nil {
		ctx.Logger.Infof("node config not found: %s", err)
	}

	ctx.Logger.Infof("Classifying batch %q", req.EmailsKey)

	emailsBatch, err := ctx.ObjectStore.Get(req.EmailsKey)
	if err != nil {
		return fmt.Errorf("error getting emails batch from object store: %w", err)
	}

	var emails []*Email
	if err := gocsv.UnmarshalBytes(emailsBatch, &emails); err != nil {
		return fmt.Errorf("error unmarshaling emails: %w", err)
	}

	for _, email := range emails {
		category := clasifyEmail(email)
		output := &proto.ClassificatorOutput{
			Email: &proto.Email{
				Title:        email.Title,
				Body:         email.Body,
				Author:       email.Author,
				CreationDate: email.CreationDate,
			},
			Category: category,
		}

		if category == proto.EmailCategory_CATEGORY_REPARATIONS {
			err := ctx.SendOutput(output, "repairs")
			if err != nil {
				return fmt.Errorf("error sending output to repairs: %w", err)
			}
		}

		err = ctx.SendOutput(output)
		if err != nil {
			return fmt.Errorf("error sending output: %w", err)
		}
	}

	return nil
}

func clasifyEmail(_ *Email) proto.EmailCategory {
	categories := []proto.EmailCategory{
		proto.EmailCategory_CATEGORY_REPARATIONS,
		proto.EmailCategory_CATEGORY_ADMINISTRATION,
		proto.EmailCategory_CATEGORY_BILLING,
		proto.EmailCategory_CATEGORY_SPAM,
	}

	return categories[rand.Intn(len(categories)-1)]
}

func main() {
	kre.Start(handlerInit, handler)
}
