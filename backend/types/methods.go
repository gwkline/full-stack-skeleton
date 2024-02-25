package types

func (u *User) IsAdmin() bool {
	isAdmin := false
	if u.Role == RoleAdmin {
		isAdmin = true
	}

	return isAdmin
}

func (s BaseModel) GetID() uint {
	return s.ID
}
