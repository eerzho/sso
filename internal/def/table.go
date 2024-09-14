package def

type TableName string

const (
	TableUsers         TableName = "users"
	TableRefreshTokens TableName = "refresh_tokens"
	TableRoles         TableName = "roles"
	TablePermissions   TableName = "permissions"
)

func (tn TableName) String() string {
	return string(tn)
}
