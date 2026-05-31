package helper

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)


//TLS Dial → NewClient → Auth → Mail(from) → Rcpt(to) → Data() → Write → Close → Quit

func Send_Mail(to, sub, body string) error {
	from := "deepakluitel679@gmail.com" //server mail this is used to send mail
	password := "nuqrkjtmxqjwnary"      //app password using this gmail server verifies before sending...

	msg := "From: CourseHub<deepakluitel679@gmail.com>\r\nSubject: " + sub + "\r\n\r\n" + body //actual email format...

	//step 1: create encrypted path to gmail server on internet..conn is the path blueprint our client will use when sending

	conn, err := tls.Dial("tcp", "smtp.gmail.com:465", &tls.Config{ //takes network protocol "tcp",smtp.gmail.com(sending email server):465 is tunneled port
		ServerName:         "smtp.gmail.com", //verifying server name where we will send
		InsecureSkipVerify: true,             //Disabling some certificates verification...
	})
	if err != nil {
		return err
	}
	//setting up client

	client, clErr := smtp.NewClient(conn, "smtp.gmail.com") //creates new client that will send msg to smtp server using conn path

	if clErr != nil {
		return clErr
	}
	//client login verifying the client..
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com") //returns auth that will be used to verify..

	if Err := client.Auth(auth); Err != nil { //Auth will authenticate client using above plain authentication mechanism..if failed connection closes
		return Err
	}
	//setting up sender
	if Serr := client.Mail(from); Serr != nil {
		return Serr
	}
	//setting up reciever
	if Rerr := client.Rcpt(to); Rerr != nil {
		return Rerr
	}

	//writing message..

	writter, Werr := client.Data() //this method provides us with writer to write message
	if Werr != nil {
		return Werr
	}
	fmt.Fprintf(writter, msg) //Fprintf format string accordingly and writes it in writter
	writter.Close()           //closing writter because it occupies memory
	client.Quit()             //after we are done sending msg we close connection of client with mail server

	return nil
}
