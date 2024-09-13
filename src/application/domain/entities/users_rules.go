package entities

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleOperator UserRole = "operator"
)

type UserRole string

func (u UserRole) Value() string {
	return string(u)
}

func (u UserRole) IsAdmin() bool {
	return u == UserRoleAdmin
}

func (u UserRole) IsOperator() bool {
	return u == UserRoleOperator
}
