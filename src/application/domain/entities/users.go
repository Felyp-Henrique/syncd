package entities

const (
	UserTypeOperator int = iota
	UserTypeAdmin
)

const (
	UserStatusDeleted int = iota
	UserStatusInactived
	UserStatusActived
)

type User struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
	Type     int    `db:"type" json:"type"`
	Status   int    `db:"status" json:"status"`
}

func NewUser() *User {
	return &User{}
}
