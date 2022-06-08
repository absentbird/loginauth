# LoginAuth
Login Authorization option for sending email over SMTP with Go. Particularly useful for Office365 accounts such as Outlook.com or Exchange Online.

If you've been trying to send email over SMTP and keep getting an error like `"Error tls: first record does not look like a TLS handshake"` this package should help.

## EXAMPLE

```go
server   := "smtp.office365.com:587"
sender   := "example1@example.com"
password := "hunter1"
message  := "Hello."
receiver := "example2@example.com"

auth := loginauth.Auth(sender, password)

err := smtp.SendMail(server, auth, sender, []string{receiver}, []byte(message))
if err != nil {
    log.Fatal(err)
}
log.Println("Email sent successfully.")
```
