package mailer

const (
	FromName = "GoSocial"
)

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) error
}
