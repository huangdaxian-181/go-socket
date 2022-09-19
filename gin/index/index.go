package index

import "go-socket/event"

func test() {

	event.SubEvent(event.EventExample(10, byte("你好")))
}
