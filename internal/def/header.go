package def

type HeaderKey string

const (
	HeaderRequestID   HeaderKey = "X-Request-Id"
	HeaderContentType HeaderKey = "Content-Type"
)

func (hk HeaderKey) String() string {
	return string(hk)
}
