package main

import (
	"gopkg.in/gomail.v2"
	"fmt"
)

func main() {
	mailer := gomail.NewMessage()
	from := "devgrumpycat@gmail.com"
	to := "rneko2006@gmail.com"

	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Test from research again")
	mailer.SetBody("text/html", "Hello world! <a href='https://www.google.com'>Click me!</a>")

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "devgrumpycat@gmail.com", "ivcrcydpapcpakqz")

	err := dialer.DialAndSend(mailer)

	if err != nil {
		panic(err)
	}

	fmt.Println("Mail has been sended!")



}
