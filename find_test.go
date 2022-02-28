package garrays

import (
	"log"
	"testing"
)

func TestFind(t *testing.T) {
	log.Println(users)
	foundName := *Find(users, func(user User) bool {
		return user.Age > 50
	})
	log.Println(foundName)
}
