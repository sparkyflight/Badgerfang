package main

import (
	"context"
	"fmt"
	"net/http"
)

type CoreVariables struct {
	Name    string
	Version string
	APIURL  string
	Token   string
	Context context.Context
}

func CallHTTP(c *CoreVariables, endpoint string, method string) (*http.Response, error) {
	url := c.APIURL + endpoint

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", c.Name+"/"+c.Version)
	request.Header.Set("Authorization", c.Token)

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return response, nil
}

func main() {
	// Example usage:
	coreVars := &CoreVariables{
		Name:    "Sparkyflight",
		Version: "0.1",
		APIURL:  "https://api.sparkyflight.xyz/",
		Token:   "",
		Context: context.Background(),
	}

	p, err := CallHTTP(coreVars, "posts/list", "GET")
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}
