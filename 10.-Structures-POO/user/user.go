package user

import "time"

// User : Esta estructura sería el equivalente a una clase en POO
type User struct {
	ID       int
	Name     string
	CreateAt time.Time
	Status   bool
}

// UserInfo : Este método asigna valores a las variables de tipo User
func (u *User) UserInfo(id int, name string, createAt time.Time, status bool) {
	u.ID = id
	u.Name = name
	u.CreateAt = createAt
	u.Status = status
}
