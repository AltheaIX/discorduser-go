package discorduser

import (
	"fmt"
	"testing"
)

var s, _ = Auth("")

func TestGetCurrentUser(t *testing.T) {
	user, err := s.GetCurrentUser()
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(user)
}

func TestGetUser(t *testing.T) {
	user, err := s.GetUser("475468428300386315")
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(user)
}

func TestGetUserGuilds(t *testing.T) {
	guilds, err := s.GetUserGuilds()
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println((*guilds)[1])
}
