package main

import (
	"fmt"
	base62 "github.com/eknkc/basex"
	"strings"
)

func main() {

	enc, _ := base62.NewEncoding("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	decodeParam, _ := enc.Decode(url)
	rawQueryParams := strings.Split(fmt.Sprintf("%s", decodeParam), "=")

	var location string
	var contId string
	for i, val := range rawQueryParams {
		if i == 1 {
			location = strings.Replace(val, "&cont_id", "", -1)
		} else if i == 2 {
			contId = val
		}
	}
	fmt.Println(location)
	fmt.Println(contId)
	return location, contId, nil
	location, contId, _ := util.UrlBase62Decoding("81KeKwGaJbr99HdhEHtxyST0MjpXXd6GexPfwZR3aNv1yKXQk3")

}

type Attt struct {
	Name string `json:"name"`
}
