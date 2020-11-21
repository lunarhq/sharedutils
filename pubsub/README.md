To create new topic
     kubectl exec --stdin --tty kafka-0 -- kafka-topics.sh --zookeeper zk-0.zk-hs:2181 --create --topic user.created --replication-factor 1 --partition 1



How to use
=============

	import (
		"github.com/lunarhq/sharedutils/pubsub"
		"github.com/lunarhq/sharedutils/pubsub/writer"
		"github.com/lunarhq/sharedutils/pubsub/reader"
	)

	//Publishing messages
	writer := writer.NewWriter(pubsub.TopicKeyCreated)
	writer.KeyCreated(pubsub.Key{})


	//Reader messages
	reader := reader.NewKeyReader(pubsub.TopicKeyCreated, "group-id")
	for {
		key, err := reader.Read()
	}
	reader.Close()
