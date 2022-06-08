import (
    "errors"
    "net/smtp"
)

package github.com/absentbird/loginauth

type login_auth struct {
    username string
    password string
}
func (s *login_auth) Start(server *smtp.ServerInfo) (string, []byte, error) {
    return "LOGIN", []byte(s.username), nil
}
func (s *login_auth) Next(fromServer []byte, more bool) ([]byte, error) {
    if more {
        switch string(fromServer) {
        case "Username:":
            return []byte(s.username), nil
        case "Password:":
            return []byte(s.password), nil
        default:
            return nil, errors.New("Unknown from server")
        }
    }
    return nil, nil
}

func Auth(username, password string) smtp.Auth {
    return &login_auth{username, password}
}
