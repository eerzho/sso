package def

type HeaderKey string

const (
	RequestID   HeaderKey = "X-Request-Id"
	ContentType HeaderKey = "Content-Type"
)
