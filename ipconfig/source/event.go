package source

import "github.com/oim/common/discovery"

var eventChan chan *Event

// EventChan 从chan读取到这里
func EventChan() <-chan *Event {
	return eventChan
}

type EventType string

const (
	// AddNodeEvent 添加网关节点
	AddNodeEvent EventType = "addNode"
	// DelNodeEvent 删除网关节点
	DelNodeEvent EventType = "delNode"
)

type Event struct {
	Type         EventType
	Ip           string
	Port         string
	ConnectNum   float64
	MessageBytes float64
}

func NewEvent(ed *discovery.EndpointInfoModel) *Event {
	if ed == nil || ed.MetaData == nil {
		return nil
	}
	var connNum, msgBytes float64
	if data, ok := ed.MetaData["connect_num"]; ok {
		connNum = data.(float64) // 如果出错，此处应该panic 暴露错误
	}
	if data, ok := ed.MetaData["message_bytes"]; ok {
		msgBytes = data.(float64) // 如果出错，此处应该panic 暴露错误
	}
	return &Event{
		Type:         AddNodeEvent,
		Ip:           ed.IP,
		Port:         ed.Port,
		ConnectNum:   connNum,
		MessageBytes: msgBytes,
	}
}
