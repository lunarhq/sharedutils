package pubsub

import (
	"context"

	pb "cloud.google.com/go/pubsub"
	"github.com/lunarhq/sharedutils/env"
)

type Reader struct {
	ctx    context.Context
	client *pb.Client
	Sub    *pb.Subscription
}

func NewReader(ctx context.Context, subscription string) (*Reader, error) {
	projectID := env.Get("PROJECT_ID", "")
	client, err := pb.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	sub := client.Subscription(subscription)

	return &Reader{ctx: ctx, client: client, Sub: sub}, nil
}
