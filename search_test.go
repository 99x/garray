package garrays

import (
	"log"
	"testing"
)

func TestSearch(t *testing.T) {
	log.Println(users)
	log.Printf("Count : %d", len(users))
	searchData := Search(users, func(user User, _ int) bool {
		return user.Age > 50
	})
	log.Println(searchData)
	log.Printf("Count : %d", len(searchData))
}
