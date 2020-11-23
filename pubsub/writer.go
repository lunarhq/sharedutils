package pubsub

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/ksuid"
)

type Writer struct {
	w map[string]*kafka.Writer
}

func NewWriter() *Writer {
	return &Writer{w: make(map[string]*kafka.Writer)}
}

func (p *Writer) getWriter(topic string) *kafka.Writer {
	w, found := p.w[topic]
	if !found {
		w = kafka.NewWriter(kafka.WriterConfig{
			Brokers:  strings.Split(PubsubBrokers, ","),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		})
		p.w[topic] = w
	}
	return w
}

func (p *Writer) Write(topic string, data interface{}) error {
	w := p.getWriter(topic)
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	//@Todo Overridable key
	key := topic + "_" + ksuid.New().String()
	//@Todo better context
	return w.WriteMessages(context.Background(), kafka.Message{Key: []byte(key), Value: bytes})
}
