package pubsub

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/lunarhq/sharedutils/env"
	"github.com/lunarhq/sharedutils/types"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	"github.com/segmentio/ksuid"
)

type Writer struct {
	w map[string]*kafka.Writer
	d *kafka.Dialer
}

func NewWriter() (*Writer, error) {
	pwd := env.Get("client-passwords", "supersecret")
	mechanism, err := scram.Mechanism(scram.SHA256, "user", pwd)
	if err != nil {
		return nil, err
	}
	dialer := &kafka.Dialer{
		Timeout:       3 * time.Second,
		SASLMechanism: mechanism,
		// DualStack:     true,
	}
	return &Writer{w: make(map[string]*kafka.Writer), d: dialer}, nil
}

func (p *Writer) getWriter(topic string) *kafka.Writer {
	w, found := p.w[topic]
	if !found {
		w = kafka.NewWriter(kafka.WriterConfig{
			Brokers:      strings.Split(PubsubBrokers, ","),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			Dialer:       p.d,
			BatchTimeout: 10 * time.Millisecond,
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

func (p *Writer) WriteErr(msg string) error {
	topic := TopicError
	w := p.getWriter(topic)

	errorType := types.Error{msg}
	bytes, err := json.Marshal(errorType)
	if err != nil {
		return err
	}
	//@Todo Overridable key
	key := topic + "_" + ksuid.New().String()
	//@Todo better context
	return w.WriteMessages(context.Background(), kafka.Message{Key: []byte(key), Value: bytes})
}
