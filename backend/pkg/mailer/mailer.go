package mailer

import (
	"backend/pkg/logger"
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"text/template"
)

func Mailer(recipient string, sender, password, domain, protocol, host, port string, data interface{}) {

	address := fmt.Sprintf("%s:%s", host, port)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	auth := smtp.PlainAuth("", sender, password, host)

	var body bytes.Buffer
	t, err := template.ParseFiles("email.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	t.Execute(&body, data)

	// Setup message
	var message = ""
	var delimeter = "simple boundary"

	// fmt.Println(recipient)

	//basic email headers
	message += fmt.Sprintf("From: %s\r\n", (&mail.Address{Name: "Granulex", Address: sender}).String())
	message += "Subject: Ссылки для скачивания\r\n"
	message += "MIME-Version: 1.0\r\n"
	message += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", delimeter)

	//place HTML message
	message += fmt.Sprintf("\r\n--%s\r\n", delimeter)
	message += "Content-Type: text/html; charset=\"utf-8\"\r\n"
	message += "Content-Transfer-Encoding: 7bit\r\n"
	message += "\r\n" + body.String()

	//place file
	message += fmt.Sprintf("\r\n--%s\r\n", delimeter)
	message += "Content-Type: text/plain; charset=\"utf-8\"\r\n"
	message += "Content-Transfer-Encoding: base64\r\n"
	message += "Content-Disposition: attachment;filename=ttt\r\n"
	// rawFile, fileErr := os.ReadFile("Example.txt")
	// if fileErr != nil {
	// 	log.Panic(fileErr)
	// }
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte("Example Text"))

	// fmt.Println(message)

	c, err := smtp.Dial(address)
	if err != nil {
		logger.ERROR(err.Error())
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		logger.ERROR(err.Error())
		c.Close()
	}

	if err = c.Auth(auth); err != nil {
		logger.ERROR(err.Error())
	}

	if err = c.Mail(sender); err != nil {
		logger.ERROR(err.Error())
		c.Close()
	}

	if err = c.Rcpt(recipient); err != nil {
		logger.ERROR(err.Error())
		c.Close()
	}

	w, err := c.Data()

	if err != nil {
		logger.ERROR(err.Error())
		c.Close()
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		logger.ERROR(err.Error())
		c.Close()
	}

	err = w.Close()

	if err != nil {
		logger.ERROR(err.Error())
		c.Close()
	}

	logger.INFO("Send mail success!")
	c.Quit()
}
