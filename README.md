# LoginAuth
Login Authorization option for sending email over SMTP with Go. Particularly useful for Office365 accounts such as Outlook.com or Exchange Online.

If you've been trying to send email over SMTP and keep getting an error like `"Error tls: first record does not look like a TLS handshake"` this package might help.

## EXAMPLE

```go
server   := "smtp.office365.com:587"
sender   := "example1@example.com"
password := "hunter2"
subject  := "Test"
message  := "Hello world."
receiver := []string{"example2@example.com"}

auth := loginauth.Auth(sender, password)

body := []byte("Subject:"+subject+"\r\n\r\n"+message)
err := smtp.SendMail(server, auth, sender, receiver, body)
if err != nil {
    log.Fatal(err)
}
log.Println("Email sent successfully.")
```
