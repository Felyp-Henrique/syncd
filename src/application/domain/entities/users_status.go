package entities

const (
	UserStatusDeleted   UserStatus = "deleted"
	UserStatusInactived UserStatus = "inactived"
	UserStatusActived   UserStatus = "actived"
)

type UserStatus string

func (u UserStatus) Value() string {
	return string(u)
}

func (u UserStatus) IsDeleted() bool {
	return u == UserStatusDeleted
}

func (u UserStatus) IsInactived() bool {
	return u == UserStatusInactived
}

func (u UserStatus) IsActived() bool {
	return u == UserStatusActived
}
