package controller

import (
	"fmt"
	"im-go/service"
)

var emailService service.EmailService

func emailHelloWorld() {
	fmt.Println(emailService.AddFriend())
}
