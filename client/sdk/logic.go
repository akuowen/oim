package sdk

const (
	MsgTypeText = "text"
)

type MessageBody struct {
	Type    string
	Name    string
	From    string
	To      string
	Content string
	Session string
}
