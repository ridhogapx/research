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
	mailer.SetBody("text/html", "Hello world! <a href='http://ec2-3-113-2-184.ap-northeast-1.compute.amazonaws.com/api/v1/auth/verify/v1.local.LVPgUDrC189Y6GkcYTc-2Jfts4KXrhMb7-paBpIfZwg7mxKKWtmzoe419PwDcJzfFRbtu_bRb43DTfyrOtW0-UWljAV9c2l-gb3qSPVXu76LNqEtOylapK6sDFqqul815zgTcy6qhwm5cKtwUUv0M25Ghf5ZIBcxkn8hsL-crIbtBG5C42FIKL3qF6DO80_Cu1Frlp-SIyI6qQ6aXQP6V0nnNfFNihtnrb1Uy9fGIAuckXIh64tOHV5AJlk5lch3onOu-zEE5zL6Lz00KcE9ASMdm8Dal8GXPk15ufWDjxLZX_0NsV9oS5-qxNdNdcJvgmesKLdU8_B9Pzbm8lxz3xpYKs7YAXQoowlF-wTzlMksGpOKz_8m1qDIxlX1aAH9c6Ji6V4-fZSeCCKk2hES_M9QDDz0YFN7xQzsiAEON-NRg215Lh8.Q29weXJpZ2h0IDIwMjMgQ1YgWmFtYW4gTm93'>Click me!</a>")

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "devgrumpycat@gmail.com", "ivcrcydpapcpakqz")

	err := dialer.DialAndSend(mailer)

	if err != nil {
		panic(err)
	}

	fmt.Println("Mail has been sended!")



}
