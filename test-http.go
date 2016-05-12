package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

func formatErrResponse(response *http.Response) error {
	body, _ := ioutil.ReadAll(response.Body)
	formattedBody := ""
	if len(body) > 0 {
		formattedBody = fmt.Sprintf(", %s", string(body))
	}
	return errors.New(fmt.Sprintf("Agent error: %d%s", response.StatusCode, formattedBody))
}
func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	wg.Add(2)
	resMap := make(map[string]interface{})
	var i int = 0
	for i = 0; i < 2; i++ {
		fmt.Println("hello -------------------------")
		fmt.Println(i)
		go func(i int) {
			defer wg.Done()
			fmt.Println("in go routine" + strconv.Itoa(i))
			request, _ := http.NewRequest("GET", "https://www.baidu.com", nil)
			httpClient := http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			}
			response, err := httpClient.Do(request)
			fmt.Println(response)
			lock.Lock()
			resMap[strconv.Itoa(i)] = formatErrResponse(response)
			lock.Unlock()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("go routine done -----------------------------")
		}(i)
	}
	wg.Wait()
	fmt.Println(len(resMap))
	for _, value := range resMap {
		fmt.Println(value)
		fmt.Println("")
	}
}
