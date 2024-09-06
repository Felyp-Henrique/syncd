package entities

const (
	USER_TYPE_OPERATOR int = iota
	USER_TYPE_ADMIN
)

const (
	USER_STATUS_DELETED int = iota
	USER_STATUS_INACTIVED
	USER_STATUS_ACTIVED
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
