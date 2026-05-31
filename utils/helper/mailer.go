package helper

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// SendGrid API → HTTPS port 443 → Render never blocks this!
// Earlier we used tls.Dial/smtp.Dial → Render blocks SMTP ports 465 and 587
// SendGrid uses HTTP request instead of SMTP connection → works on all cloud providers ✅

// SendGrid API → Send Mail → Check Response
func Send_Mail(to, sub, body string) error {
	from := "deepakluitel679@gmail.com" //server mail this is used to send mail
	apiKey := "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" //sendgrid api key used to verify before sending...

	// actual email payload in JSON format for SendGrid
	payload := fmt.Sprintf(`{
		"personalizations": [{"to": [{"email": "%s"}]}],
		"from": {"email": "%s", "name": "CourseHub"},
		"subject": "%s",
		"content": [{"type": "text/plain", "value": "%s"}]
	}`, to, from, sub, body)

	//step 1: create HTTP POST request to SendGrid API
	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", strings.NewReader(payload))
	if err != nil {
		return err
	}

	//step 2: set headers - Authorization with API key and content type
	req.Header.Set("Authorization", "Bearer "+apiKey) //sendgrid verifies using this api key before sending...
	req.Header.Set("Content-Type", "application/json")

	//step 3: execute the request - send to SendGrid server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //closing response body because it occupies memory

	//step 4: check if request was successful
	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("sendgrid error: %s", string(b))
	}

	return nil
}