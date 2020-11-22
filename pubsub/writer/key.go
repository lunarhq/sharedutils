package writer

import (
	"github.com/lunarhq/sharedutils/pubsub"
	"github.com/lunarhq/sharedutils/types"
)

func (p *PubWriter) KeyCreated(key types.Key) error {
	return p.write(pubsub.TopicKeyCreated, key)
}
func (p *PubWriter) KeyUpdated(key types.Key) error {
	return p.write(pubsub.TopicKeyUpdated, key)
}
func (p *PubWriter) KeyDeleted(key types.Key) error {
	return p.write(pubsub.TopicKeyDeleted, key)
}
