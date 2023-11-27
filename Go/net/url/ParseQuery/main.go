package main

import "net/url"

func main() {
	query := "q=go+url+parsequery&oq=go+url+parsequery&aqs=chrome..69i57j0l7.10159j0j7&sourceid=chrome&ie=UTF-8"
	values, err := url.ParseQuery(query)
	if err != nil {
		panic(err)
	}
	println(values.Get("q"))
	println(values.Get("oq"))
	println(values.Get("aqs"))
	println(values.Get("sourceid"))
	println(values.Get("ie"))
}
