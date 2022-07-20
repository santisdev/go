package main

import (
	"fmt"
)

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name != "" && mToSend.sender.number != 0 &&
		mToSend.recipient.name != "" && mToSend.recipient.number != 0 {
		return true
	}

	return false
}

func main() {
	mToSends := []messageToSend{
		{
			message: "test message ",
			sender: user{
				name:   "user 0",
				number: 16545550987,
			},
			recipient: user{
				name:   "user1",
				number: 4256686131,
			},
		},
		{
			message: "test message 1",
			sender: user{
				number: 939999418,
			},
			recipient: user{
				name:   "User2",
				number: 0,
			},
		},
		{
			message: "test message 2",
			sender: user{
				name:   "user3",
				number: 3232325141,
			},
			recipient: user{
				name:   "user4",
				number: 42425426426,
			},
		},
		{
			message: "test message 3",
			sender: user{
				name:   "user5",
				number: 0,
			},
			recipient: user{
				name:   "user6",
				number: 425624662,
			},
		},
		{
			message:   "test message 4",
			sender:    user{},
			recipient: user{},
		},
	}
	for _, mToSend := range mToSends {
		if canSendMessage(mToSend) {
			fmt.Printf(`sending "%s" from %s (%v) to %s (%v)`,
				mToSend.message,
				mToSend.sender.name,
				mToSend.sender.number,
				mToSend.recipient.name,
				mToSend.recipient.number,
			)
			fmt.Println()
		} else {
			fmt.Println("can't send message")
		}
	}
}
