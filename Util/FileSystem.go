package Util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Httpget() int {
	resp, _ := http.Get("http://c-lap.cn/")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Body : %s", body)
	return len(body)
}
