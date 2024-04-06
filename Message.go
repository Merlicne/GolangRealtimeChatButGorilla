package main

type MessageList []*Message

type Message struct {
	msg string
	wt  *Client
}

func newMessage(msg string, wt *Client) *Message {
	return &Message{
		msg: msg,
		wt:  wt,
	}
}
