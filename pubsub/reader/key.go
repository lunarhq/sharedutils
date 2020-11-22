package reader

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/lunarhq/sharedutils/types"
	"github.com/segmentio/kafka-go"
)

type KeyReader struct {
	r *kafka.Reader
}

func NewKeyReader(topic, group string) *KeyReader {
	kr := &KeyReader{}
	kr.r = kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(BROKERS, ","),
		Topic:   topic,
		GroupID: group,
	})
	return kr
}

func (kr *KeyReader) Close() error {
	return kr.r.Close()
}

func (kr *KeyReader) Read() (*types.Key, error) {
	m, err := kr.r.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}
	var result types.Key
	if err := json.Unmarshal(m.Value, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
