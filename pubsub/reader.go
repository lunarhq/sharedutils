package pubsub

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/segmentio/kafka-go"
)

const (
	// BROKERS = "kafka-0.kafka-hs:9093,kafka-1.kafka-hs:9093,kafka-2.kafka-hs:9093"
	BROKERS = "localhost:9093,localhost:9093,localhost:9093"
)

type Reader struct {
	reader *kafka.Reader
}

func NewReader(topic, group string) *Reader {
	r := &Reader{}
	r.reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(BROKERS, ","),
		Topic:   topic,
		GroupID: group,
	})
	return r
}

func (r *Reader) Read(out *interface{}) error {
	m, err := r.reader.ReadMessage(context.Background())
	if err != nil {
		return err
	}
	if err := json.Unmarshal(m.Value, out); err != nil {
		return err
	}
	return nil
}

func (r *Reader) Close() error {
	return r.reader.Close()
}
