package users

import (
	"fmt"
	"time"

	"github.com/fswebcoder/goland/models"
)

func AltaUsuarios() {
	u := new(models.User)
	u.AddUser(1, "Fernando", time.Now(), time.Now(), true)
	fmt.Println(u)
}
