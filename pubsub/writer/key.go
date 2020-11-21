package writer

import (
	"github.com/lunarhq/sharedutils/pubsub"
)

func (p *PubWriter) KeyCreated(key pubsub.Key) error {
	return p.write(pubsub.TopicKeyCreated, key)
}
func (p *PubWriter) KeyUpdated(key pubsub.Key) error {
	return p.write(pubsub.TopicKeyUpdated, key)
}
func (p *PubWriter) KeyDeleted(key pubsub.Key) error {
	return p.write(pubsub.TopicKeyDeleted, key)
}
