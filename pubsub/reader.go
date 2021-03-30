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

func NewReader(ctx context.Context) (*Reader, error) {
	projectID := env.Get("PROJECT_ID", "")
	client, err := pb.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &Reader{ctx: ctx, client: client}, nil
}

func (r *Reader) Subscription(sub string) *pb.Subscription {
	return r.client.Subscription(sub)
}
