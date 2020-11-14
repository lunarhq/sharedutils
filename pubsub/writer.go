package pubsub

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/segmentio/kafka-go"
)

type Writer struct {
	writer *kafka.Writer
}

func NewWriter(topic string) *Writer {
	w := &Writer{}
	w.writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  strings.Split(BROKERS, ","),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return w
}

func (w *Writer) Write(key string, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := w.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(bytes),
		},
	); err != nil {
		return err
	}
	return nil
}
func (w *Writer) Close() error {
	return w.writer.Close()
}
