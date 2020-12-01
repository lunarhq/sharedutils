package pubsub

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lunarhq/sharedutils/env"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type Reader struct {
	*kafka.Reader
}

func NewReader(topic, group string) (*Reader, error) {
	pwd := env.Get("client-passwords", "supersecret")
	mechanism, err := scram.Mechanism(scram.SHA256, "user", pwd)
	if err != nil {
		return nil, err
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{KafkaEndpoint},
		Topic:   topic,
		GroupID: group,
		Dialer: &kafka.Dialer{
			Timeout:       3 * time.Second,
			SASLMechanism: mechanism,
			// DualStack:     true, //@Todo is this needed?
		},
		MaxWait: time.Second,
	})
	return &Reader{r}, nil
}

func (r *Reader) Read(out interface{}) error {
	m, err := r.ReadMessage(context.Background())
	if err != nil {
		return err
	}
	return json.Unmarshal(m.Value, out)
}
