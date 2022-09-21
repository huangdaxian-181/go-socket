package event

/**
   推送消息给指定的客户端
   消息插入广播
   初始化一个struct
   消息写入struct
**/

type Event struct {
	EventExample chan *EventExample
}

type EventExample struct {
	WorkId int
	Data   []byte
}

var EventExamples *Event = NewEvent()

func NewEvent() *Event {
	return &Event{
		EventExample: make(chan *EventExample),
	}
}

func NewEventExample(WorkId int, Data []byte) *EventExample {
	return &EventExample{
		WorkId: WorkId,
		Data:   Data,
	}
}

func (e *Event) PushMsg(n *EventExample) {
	e.EventExample <- n
}

func (e *Event) Chan() chan *EventExample {
	return e.EventExample
}
