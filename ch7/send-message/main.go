package main

import "fmt"

type Messenger interface {
	Send(m string)
}

type SMS struct {
	PhoneNumber string
}

func (s SMS) Send(body string) {
	fmt.Printf("send message to %s  -  %s\n", s.PhoneNumber, body)
}

type Email struct {
	Address string
}

func (e Email) Send(body string) {
	fmt.Printf("send message to %s  -  %s\n", e.Address, body)
}

func SendNotification(m Messenger, body string) {
	m.Send(body)
}

type Slack struct {
	person string
}

func main() {
	sms := SMS{PhoneNumber: "98383838383"}
	email := Email{Address: "bhanu@gmail.com"}

	SendNotification(sms, "Hey")
	SendNotification(email, "Hey! There")

}
