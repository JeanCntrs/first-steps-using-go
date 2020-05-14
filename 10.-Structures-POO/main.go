package main

import (
	"fmt"
	"time"

	us "./user"
)

// Esta estructura hereda el objeto User importado con sus variables y métodos
type marcela struct {
	us.User
}

func main() {
	newUser := new(marcela)
	newUser.UserInfo(1, "Marcela S.", time.Now(), true)
	fmt.Println(newUser.User)
}
