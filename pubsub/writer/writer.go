package writer

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/ksuid"
)

const (
	//@Todo move to env
	BROKERS = "kafka-0.kafka-hs:9093,kafka-1.kafka-hs:9093,kafka-2.kafka-hs:9093"
	// BROKERS = "localhost:9093,localhost:9093,localhost:9093"
)

type PubWriter struct {
	w map[string]*kafka.Writer
}

func New() (*PubWriter, error) {
	p := &PubWriter{}
	return p, nil
}

func (p *PubWriter) getWriter(topic string) *kafka.Writer {
	w, found := p.w[topic]
	if !found {
		w = kafka.NewWriter(kafka.WriterConfig{
			Brokers:  strings.Split(BROKERS, ","),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		})
		p.w[topic] = w
	}
	return w
}

func (p *PubWriter) write(topic string, data interface{}) error {
	w := p.getWriter(topic)
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	key := topic + "_" + ksuid.New().String()
	//@Todo better context
	return w.WriteMessages(context.Background(), kafka.Message{Key: []byte(key), Value: bytes})
}
