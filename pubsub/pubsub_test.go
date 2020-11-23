package pubsub

import (
	"log"
	"testing"
)

func TestWriter(t *testing.T) {
	w := NewWriter()
	err := w.Write(TopicKeyCreated, "test")
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestReader(t *testing.T) {
	_, err := NewReader(TopicKeyCreated, "group")
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	// var out interface{}
	// if err := r.Read(&out); err != nil {
	// 	log.Println(err)
	// 	t.Fail()
	// }
}
