package pubsub

const (
	//@Todo move to env
	PubsubBrokers = "kafka-0.kafka-hs:9093,kafka-1.kafka-hs:9093,kafka-2.kafka-hs:9093"
)

const (
	TopicKeyCreated = "key.created"
	TopicKeyUpdated = "key.updated"
	TopicKeyDeleted = "key.deleted"

	TopicError = "error"
)
