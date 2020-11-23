package slack

import (
	"testing"
)

func TestSlack(t *testing.T) {
	Post(":tada: Account created", Field{"title", "value"}, Field{"title2", "value2"})
}
