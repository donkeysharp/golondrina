package store

const (
	GITHUB_PROVIDER = "github"
)

type NotificationEvent struct {
	Id       string
	Title    string
	Url      string
	Provider string
}
