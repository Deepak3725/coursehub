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

	//step 1: smtp.Dial creates a plain unencrypted connection to port 587 first
	//Port 587 expects plain connection first, then upgrade to TLS (STARTTLS)
	//tls.Dial was wrong here because port 587 server expects plain "hello" first
	//tls.Dial sends encrypted bytes immediately → server gets confused → timeout
	//smtp.Dial sends plain first → then we upgrade via StartTLS → correct handshake ✅
	client, err := smtp.Dial("smtp.gmail.com:587")
	if err != nil {
		return err
	}
	defer client.Quit()

	//step 2: upgrade plain connection to encrypted using STARTTLS
	//this is the TLS upgrade that port 587 expects after plain connection
	if err = client.StartTLS(&tls.Config{
		ServerName:         "smtp.gmail.com", //verifying server name where we will send
		InsecureSkipVerify: true,             //Disabling some certificates verification...
	}); err != nil {
		return err
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