package reader

import (
	"strings"

	"github.com/segmentio/kafka-go"
)

const (
	//@Todo move to env
	BROKERS = "kafka-0.kafka-hs:9093,kafka-1.kafka-hs:9093,kafka-2.kafka-hs:9093"
)

type PubReader struct {
	r map[string]*kafka.Reader
}

func New() (*PubReader, error) {
	p := &PubReader{}
	return p, nil
}

//@Todo this doesn't return err atm.
func (p *PubReader) Close() {
	for _, v := range p.r {
		v.Close()
	}
}

func (p *PubReader) getReader(topic, group string) *kafka.Reader {
	uniqKey := topic + group
	r, found := p.r[uniqKey]
	if !found {
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers: strings.Split(BROKERS, ","),
			Topic:   topic,
			GroupID: group,
		})
		p.r[uniqKey] = r
	}
	return w
}
