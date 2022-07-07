package main

import (
	"fmt"
	"log"
	"net/smtp"
	// "os"
)

func sendMail() {
	// sender data
	from := "John.Doe@gmail.com"
	// os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	password := "ieiemcjdkejspqz"
	// os.Getenv("SMTPpwd")   // ex: "ieiemcjdkejspqz"
	// receiver address
	toEmail := "Jane.Smith@yahoo.com"
	// os.Getenv("ToEmailAddr") // ex: "Jane.Smith@yahoo.com"
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, host)
	msg := []byte(
		"From: Golang <" + from + ">\r\n" +
			"To: " + toEmail + "\r\n" +
			"Subject: Let's Go\r\n" +
			"\r\n" +
			"This is the email body.\r\n")
	err := smtp.SendMail(address, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email is send, check your inbox...")
}
