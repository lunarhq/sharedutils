package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	pb "cloud.google.com/go/pubsub"
	"github.com/lunarhq/sharedutils/env"
	"github.com/lunarhq/sharedutils/types"
)

type Writer struct {
	ctx    context.Context
	client *pb.Client
	topics map[string]*pb.Topic
	mu     sync.RWMutex
}

func NewWriter(ctx context.Context) (*Writer, error) {
	projectID := env.Get("PROJECT_ID", "")
	client, err := pb.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &Writer{ctx: ctx,
		client: client,
		topics: map[string]*pb.Topic{},
		mu:     sync.RWMutex{}}, nil
}

func (p *Writer) getTopic(topic string) *pb.Topic {
	p.mu.RLock()
	t, ok := p.topics[topic]
	p.mu.RUnlock()
	if !ok {
		t = p.client.Topic(topic)
		p.mu.Lock()
		p.topics[topic] = t
		p.mu.Unlock()
	}
	return t
}

func (p *Writer) Write(topic string, data interface{}) error {
	t := p.getTopic(topic)

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	res := t.Publish(p.ctx, &pb.Message{Data: bytes})
	log.Println("got resp:", topic, res)
	//@Todo sync/async??
	r, err = res.Get(p.ctx)
	log.Println("Write result", r)
	log.Println("Write err:", err)
	return err
	// return nil
}

func (p *Writer) WriteErr(msg string) error {
	data := types.Error{msg}
	return p.Write(TopicError, data)
}

func (p *Writer) Close() {
	for _, t := range p.topics {
		t.Stop()
	}
}
