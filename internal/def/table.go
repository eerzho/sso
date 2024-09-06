package def

type TableName string

const (
	TableUsers TableName = "users"
)

func (tn TableName) String() string {
	return string(tn)
}
