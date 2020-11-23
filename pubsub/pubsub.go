package pubsub

const (
	//@Todo move to env
	PubsubBrokers = "kafka-0.kafka-hs:9093,kafka-1.kafka-hs:9093,kafka-2.kafka-hs:9093"
)

const (
	TopicAccountCreated = "account.created"
	TopicAccountUpdated = "account.updated"
	TopicAccountDeleted = "account.deleted"

	TopicKeyCreated = "key.created"
	TopicKeyUpdated = "key.updated"
	TopicKeyDeleted = "key.deleted"

	TopicPaymentMethodCreated = "paymentmethod.created"
	TopicPaymentMethodUpdated = "paymentmethod.updated"
	TopicPaymentMethodDeleted = "paymentmethod.deleted"

	TopicError = "error"
)
