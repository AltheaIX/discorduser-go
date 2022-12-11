package discorduser

import (
	"fmt"
	"testing"
)

func TestGetCurrentUser(t *testing.T) {
	s, _ := Auth("")
	user, err := s.GetCurrentUser()
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(user.Verified)
}
