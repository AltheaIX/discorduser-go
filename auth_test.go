package discorduser

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	s, err := Auth("")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(s.Authorization)
}
