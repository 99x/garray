package garrays

import (
	"log"
	"testing"
)

func TestRemove(t *testing.T) {
	localUsers := *ReadUserData()

	log.Printf("Count: %d", len(localUsers))
	log.Printf("1st user's name: %s", localUsers[0].Name)
	log.Printf("2nd user's name: %s", localUsers[1].Name)

	localUsers = Remove(localUsers, 0)
	log.Println("First element removed")

	log.Printf("Count: %d", len(localUsers))
	log.Printf("1st user's name: %s", localUsers[0].Name)

}

func TestSplice(t *testing.T) {
	localUsers := *ReadUserData()

	log.Printf("Count: %d", len(localUsers))
	log.Printf("1st user's name: %s", localUsers[0].Name)
	log.Printf("5th user's name: %s", localUsers[4].Name)

	localUsers = Splice(localUsers, 1, 4)
	log.Println("First element removed")
	log.Printf("Count: %d", len(localUsers))
	log.Printf("1st user's name: %s", localUsers[0].Name)
	log.Printf("2nd user's name: %s", localUsers[0].Name)

}

func TestFilter(t *testing.T) {
	localUsers := *ReadUserData()

	log.Printf("Count: %d", len(localUsers))

	localUsers = Filter(localUsers, func(user User, _ int) bool {
		return user.Age > 50
	})
	log.Println("Array filtered if age > 50")
	log.Printf("Count: %d", len(localUsers))

}
