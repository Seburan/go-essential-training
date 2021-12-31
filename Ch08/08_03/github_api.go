/*Challenge : Call the Github API
- Write a function that queries a GitHub User's API for a given login
- it should return a struct of User, with a name, and a number of public
repositories
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const githubUserAPIUrl string = "https://api.github.com/users";

type User struct {
	Name	string `json:"name"`
	PublicRepos int	`json:"public_repos"`
	Login string `json:"login"`
	Company string `json:"company"`
}


func userInfo(login string) (*User, error) {
	// build target api url string
	httpUrl := fmt.Sprintf("%s/%s",githubUserAPIUrl, login);

	resp, err := http.Get(httpUrl);
	if err != nil {
		return nil, err;
	}

	defer resp.Body.Close();

	// Decode response
	respDecoder := json.NewDecoder(resp.Body);

	// prepare a user struct to load data into
	userData := &User{};

	// decode the response and load data into struct
	if err := respDecoder.Decode(userData); err != nil {
		return nil, err;
	}

	return userData, nil;

}

func main() {

	// lookup seburan's info
	login := "seburan";
	myUserInfo, err := userInfo(login);
	if err != nil {
		log.Fatalf("Error : cannot fetch %s'info - %s", login, err);
	}

	fmt.Printf("Got :\n %+v\n", *myUserInfo);
}
