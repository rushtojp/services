package example

import (
	"fmt"
	"github.com/micro/services/clients/go/twitter"
	"os"
)

// Get a user's twitter profile
func GetAusersTwitterProfile() {
	twitterService := twitter.NewTwitterService(os.Getenv("MICRO_API_TOKEN"))
	rsp, err := twitterService.User(&twitter.UserRequest{
		Username: "crufter",
	})
	fmt.Println(rsp, err)
}
