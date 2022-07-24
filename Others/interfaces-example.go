package main

import (
	"fmt"
	"time"
)

type message interface {
	getMessage() string
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main() {
	msgs := []message{
		sendingReport{
			reportName:    "First Report",
			numberOfSends: 10,
		},
		birthdayMessage{
			recipientName: "John Doe",
			birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
		},
		sendingReport{
			reportName:    "First Report",
			numberOfSends: 10,
		},
		birthdayMessage{
			recipientName: "Bill Deer",
			birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, msg := range msgs {
		sendMessage(msg)
	}
}
