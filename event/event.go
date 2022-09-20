package event

/**
   推送消息给指定的客户端
   消息插入广播
   初始化一个struct
   消息写入struct
**/

type Event struct {
	EventExample chan interface{}
}

type EventExample struct {
	WorkId int

	Data byte
}

var EventExamples *Event = NewEvent()

func NewEvent() *Event {
	return &Event{
		EventExample: make(chan interface{}, 200),
	}
}

func NewEventExample(WorkId int, Data byte) *EventExample {
	return &EventExample{
		WorkId: WorkId,
		Data:   Data,
	}
}

//发送消息
func (e *Event) SubEvent(ev *EventExample) {

	go func() {
		e.EventExample <- ev

	}()
}

func (e *Event) GetEvent() chan interface{} {
	return e.EventExample
}
