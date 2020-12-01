package pubsub

const (
	//@Todo move to env
	KafkaEndpoint = "shared-kafka.default.svc.cluster.local"
	PubsubBrokers = "shared-kafka-0.shared-kafka-headless.default.svc.cluster.local:9092,shared-kafka-1.shared-kafka-headless.default.svc.cluster.local:9092,shared-kafka-2.shared-kafka-headless.default.svc.cluster.local:9092"
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

	TopicError        = "error"
	TopicApiRequested = "api.requested"
)
