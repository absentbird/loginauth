# LoginAuth
Login Authorization option for sending email over SMTP with Go. Particularly useful for Office365 accounts such as Outlook.com or Exchange Online.

## EXAMPLE

```go
receiver := "example2@example.com"
sender   := "example1@example.com"
password := "hunter1"
message  := []byte("Hello.")

auth := loginauth.Auth(sender, password)

err := smtp.SendMail("smtp.office365.com:587", auth, sender, receiver, message)
if err != nil {
    log.Fatal(err)
    return
}
log.Println("Email sent successfully!")
```
