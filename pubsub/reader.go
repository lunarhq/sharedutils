package pubsub

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/segmentio/kafka-go"
)

type Reader struct {
	*kafka.Reader
}

func NewReader(topic, group string) (*Reader, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(PubsubBrokers, ","),
		Topic:   topic,
		GroupID: group,
	})
	return &Reader{r}, nil
}

func (r *Reader) Read(out *interface{}) error {
	m, err := r.ReadMessage(context.Background())
	if err != nil {
		return err
	}
	return json.Unmarshal(m.Value, out)
}
