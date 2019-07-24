package controller

import (
	"fmt"
	"web/service"
)

var emailService service.EmailService

func emailHelloWorld() {
	fmt.Println(emailService.AddFriend())
}
