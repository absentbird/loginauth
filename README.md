# LoginAuth
Login Authorization option for sending email over SMTP with Go. Particularly useful for Office365 accounts such as Outlook.com or Exchange Online.

## EXAMPLE

```go
server   := "smtp.office365.com:587"
sender   := "example1@example.com"
password := "hunter1"
message  := []byte("Hello.")
receiver := "example2@example.com"

auth := loginauth.Auth(sender, password)

err := smtp.SendMail(server, auth, sender, receiver, message)
if err != nil {
    log.Fatal(err)
}
log.Println("Email sent successfully.")
```
