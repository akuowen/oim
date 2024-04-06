package sdk

type connect struct {
	serverAddr string
	sendChan   chan *MessageBody
	revChan    chan *MessageBody
}

func newConnect(serverAddr string) *connect {
	return &connect{
		serverAddr: serverAddr,
		sendChan:   make(chan *MessageBody),
		revChan:    make(chan *MessageBody),
	}
}

func (c *connect) sendMessage(body *MessageBody) {
	c.revChan <- body
}

func (c *connect) rev() <-chan *MessageBody {
	return c.revChan
}
