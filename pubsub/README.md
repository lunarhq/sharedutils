To create new topic
     kubectl exec --stdin --tty kafka-0 -- kafka-topics.sh --zookeeper zk-0.zk-hs:2181 --create --topic user.created --replication-factor 1 --partition 1



How to use
=============

	import (
		"github.com/lunarhq/sharedutils/pubsub"
		"github.com/lunarhq/sharedutils/pubsub/writer"
		"github.com/lunarhq/sharedutils/pubsub/reader"
	)

	//Publish messages
	writer := writer.New(pubsub.TopicKeyCreated)
	writer.KeyCreated(pubsub.Key{})

	//Publish messages
	pub := pubsub.NewWriter()
	pub.Write(pubsub.TopicKeyCreated, &key)


	//Read messages
	reader := reader.New(pubsub.TopicKeyCreated, "group-id")
	defer reader.Close()

	for {
		var result types.Key
		err := reader.Read(&result)
		//Do stuff with result and err
	}
