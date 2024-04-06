package client

import "github.com/oim/client/sdk"

func RunCui() {
	chat := sdk.NewChat("127.0.0.1", "akuo", "111", "222")
	go doRev(chat)
}

func doRev(chat *sdk.Chat) {
	rev := chat.Rev()
	for body := range rev {
		//chat.Send(body)
		print(body)
		// TODO 处理消息
	}
}
