package discorduser

import (
	"encoding/json"
	"fmt"
)

func (s *Session) GetCurrentUser() (user *User, err error) {
	result, err := s.RequestWithAuth("https://discord.com/api/users/@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal(result, &user)
	return
}
