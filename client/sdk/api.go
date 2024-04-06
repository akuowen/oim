package sdk

type Chat struct {
	NickName  string
	SessionId string
	UserId    string
	conn      *connect
}

func NewChat(serverAddr, nickName, sessionId, userId string) *Chat {
	return &Chat{
		NickName:  nickName,
		SessionId: sessionId,
		UserId:    userId,
		conn:      newConnect(serverAddr),
	}
}

func (c *Chat) Send(body *MessageBody) {
	c.conn.sendMessage(body)
}

func (c *Chat) Rev() <-chan *MessageBody {
	return c.conn.rev()
}
