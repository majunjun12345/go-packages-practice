package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

// hawegppdiaiobdag
func main() {
	auth := smtp.PlainAuth("", "598959594@qq.com", "nqetiwlrqasgbbca", "smtp.qq.com")
	to := []string{"15527254815@163@qq.com"}
	nickname := "test"
	user := "953637695@qq.com"
	subject := "test mail"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "This is the email body."
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v\n", err)
	}
}
