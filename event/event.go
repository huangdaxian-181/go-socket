package event

/**
   推送消息给指定的客户端
   消息插入广播
   初始化一个struct
   消息写入struct
**/

type Event struct {
	E chan *EventExample
}

type EventExample struct {
	WorkId int

	Data byte
}

func NewEvent() *Event {
	return &Event{
		E: make(chan *EventExample),
	}
}

func NewEventExample(WorkId int, Data byte) *EventExample {
	return &EventExample{
		WorkId: WorkId,
		Data:   Data,
	}
}

func SubEvent(E chan *EventExample) *Event {
	return &Event{
		E: E,
	}
}
