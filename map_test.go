package garrays

import (
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	log.Println(users)
	names := *Map(users, func(users User, _ int) string {
		return users.Name
	})

	log.Println(names)
}
