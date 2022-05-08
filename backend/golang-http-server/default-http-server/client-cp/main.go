package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Dari contoh yang diberikan, cobalah untuk mengimplementasikan sebuah http client sederhana.
// Buatlah sebuah http client dan lakukan request GET ke API https://www.metaweather.com/api/.
// Buatlah request get ke endpoint /api/location/(woeid)/(date)/ dengan nilai woeid 1047378.
// Untuk date, gunakan format YYYY/MM/dd

func main() {
	// TODO: answer here
	client := http.Client{}

	resp, err := client.Get(fmt.Sprintf("https://www.metaweather.com/api/location/%s/%s", "1047378", time.Date(2022, time.May, 7, 0, 0, 0, 0, time.Local).Format("YYY/MM/dd")))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(res))
}
