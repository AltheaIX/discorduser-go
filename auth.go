package discorduser

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const VER = "0.01.0"

func Auth(token string) (s *Session, err error) {
	s = &Session{
		Authorization: token,
	}

	fmt.Println("You're using the " + VER + " version of Discord")
	return
}

func (s *Session) RequestWithAuth(url string) (result []byte, err error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://discord.com/api"+url, nil)
	req.Header.Add("authorization", s.Authorization)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
