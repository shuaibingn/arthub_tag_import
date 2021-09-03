package test

import (
	"fmt"
	"testing"

	"arthub_tag_import/internal"
)

func TestGet(t *testing.T) {
	url := "https://httpbin.org/get"
	resp, err := internal.Get(url, nil)
	if err != nil {
		panic(err)
	}

	res := string(resp)
	fmt.Println(res)
}

func TestPost(t *testing.T) {
	url := "https://httpbin.org/post"
	resp, err := internal.Post(url, nil, nil)
	if err != nil {
		panic(err)
	}

	res := string(resp)
	fmt.Println(res)
}
