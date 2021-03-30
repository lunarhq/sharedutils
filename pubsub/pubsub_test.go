package pubsub

import (
	"context"
	"log"
	"testing"
	"time"

	pb "cloud.google.com/go/pubsub"
)

func TestWriter(t *testing.T) {
	ctx := context.Background()
	w, err := NewWriter(ctx)
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	payload := struct {
		Msg  string
		Time time.Time
	}{
		Msg:  "test:" + time.Now().String(),
		Time: time.Now(),
	}
	err = w.Write(TopicTest, payload)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestReader(t *testing.T) {
	ctx := context.Background()
	reader, err := NewReader(ctx, "test-sub2")
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	err = reader.Sub.Receive(ctx, func(c context.Context, msg *pb.Message) {
		log.Println("[msg]:", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		log.Println("err receiving:", err)
		t.Fail()
	}
	log.Println("Finished")
}
