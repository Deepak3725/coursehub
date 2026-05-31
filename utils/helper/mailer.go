package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Send_Mail sends an email using SendGrid API
// Parameters:
//   - to: recipient email address (e.g. "buyer@gmail.com")
//   - sub: email subject (e.g. "Payment Verification")
//   - body: email body text (e.g. "Your payment was successful")
// Returns:
//   - error: nil if email sent successfully, error message if failed

// Flow: Build Payload → Create Request → Set Headers → Execute Request → Check Response

func Send_Mail(to, sub, body string) error {
	from := "deepakluitel679@gmail.com" // sender email - must match verified sender in SendGrid!
	apiKey := "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"

	// Step 1: Build payload using proper JSON encoding
	// json.Marshal properly escapes special characters like \n, ", % etc
	// fmt.Sprintf was breaking JSON when body contained special characters
	payload := map[string]interface{}{
		"personalizations": []map[string]interface{}{
			{"to": []map[string]string{{"email": to}}},
		},
		"from":    map[string]string{"email": from, "name": "CourseHub"},
		"subject": sub,
		"content": []map[string]string{
			{"type": "text/plain", "value": body},
		},
	}

	// json.Marshal converts map to proper JSON string handling all special characters
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Step 2: Create HTTP POST request
	// bytes.NewReader converts JSON bytes to readable stream
	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", bytes.NewReader(jsonPayload))
	if err != nil {
		return err
	}

	// Step 3: Set headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Step 4: Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Step 5: Check response
	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("sendgrid error: %s", string(b))
	}

	return nil
}