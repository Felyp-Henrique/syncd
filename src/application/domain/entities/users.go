package entities

type User struct {
	ID       int        `db:"id" json:"id"`
	Name     string     `db:"name" json:"name"`
	Email    string     `db:"email" json:"email"`
	Password string     `db:"password" json:"-"`
	Role     UserRole   `db:"type" json:"type"`
	Status   UserStatus `db:"status" json:"status"`
}

func NewUser() *User {
	return &User{
		ID:       0,
		Name:     "",
		Email:    "",
		Password: "",
		Role:     UserRoleOperator,
		Status:   UserStatusInactived,
	}
}
