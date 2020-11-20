To create new topic
     kubectl exec --stdin --tty kafka-0 -- kafka-topics.sh --zookeeper zk-0.zk-hs:2181 --create --topic user.created --replication-factor 1 --partition 1
