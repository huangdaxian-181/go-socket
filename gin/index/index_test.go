package index

import (
	"fmt"
	"go-socket/event"
	"testing"
)

func Test2(t *testing.T) {
}

func Test1(t *testing.T) {
	e := event.EventExamples

	e.SubEvent(event.NewEventExample(2131, byte(12)))

	str := e.GetEvent()

	fmt.Println(str)
}
