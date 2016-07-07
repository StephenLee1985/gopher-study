package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Secret struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

func main() {
	fmt.Println("birth cry")

	activeUrl := url.URL{
		Host:   "192.168.17.139:8080",
		Scheme: "http",
		Path:   "/postMan",
	}

	fmt.Println("active url: ", activeUrl.String())

	in := Secret{Cert: "aaaa", Key: "bbbb"}

	body, err := json.Marshal(in)

	fmt.Println("active post body: ", string(body))

	request, err := http.NewRequest("POST", activeUrl.String(), strings.NewReader(string(body)))
	if err != nil {
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("response is not 200")
		return
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

	return
}
