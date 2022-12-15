package discorduser

import (
	"encoding/json"
	"fmt"
)

func (s *Session) GetCurrentUser() (user *User, err error) {
	result, err := s.RequestWithAuth("/users/@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal(result, &user)
	return
}
func (s *Session) GetUser(id string) (user *User, err error) {
	result, err := s.RequestWithAuth("/users/" + id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal(result, &user)
	return
}

func (s *Session) GetUserGuilds() (guilds *[]Guild, err error) {
	result, err := s.RequestWithAuth("/users/@me/guilds")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal(result, &guilds)
	return
}
