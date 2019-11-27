package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func body(response *http.Response) []byte {
	page, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return page
}
func response(response *http.Response) {
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Proto)
	fmt.Println(response.ProtoMajor)
	fmt.Println(response.ProtoMinor)

	headers := response.Header
	for key, header := range headers {
		fmt.Println(key, header)
	}
	//fmt.Println(headers)
	//fmt.Println(headers.Get("Cache-Control"))
	//fmt.Println(headers.Get("Content-Type"))

}
func main() {
	res, err := http.Get("http://www.baidu.com/")
	if err != nil {
		log.Fatal(err)
	}

	response(res)
	//fmt.Printf("%s", body(res))
}
