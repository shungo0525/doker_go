package mail

import (
  "fmt"
  "net/smtp"
  "os"
  "strings"
)

var (
  hostname = "mailhog"
  port     = 1025
  username = "user@example.com"
  password = "password"
  from     = "from@example.net"
  subject  = "hello"
  body     = "Hello World!"
  receiver = []string{"receiver@example.com"}
)

func SendEmail() {
  smtpServer := fmt.Sprintf("%s:%d", hostname, port)
  auth := smtp.CRAMMD5Auth(username, password)
  msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(receiver, ","), subject, body))

  if err := smtp.SendMail(smtpServer, auth, from, receiver, msg); err != nil {
    fmt.Fprintln(os.Stderr, err)
  }
}