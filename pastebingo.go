package main

import(
	"fmt"
	"os"	
	"net/http"
	"net/url"
	"io/ioutil"
)



var APIKEY string = "HERE_IS_YOUR_PERSONNAL_DEV_KEY"
var API_URL string = "https://pastebin.com/api/api_post.php"


func extractContent(filepath string) string {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	return string(data)
}

func reqApi(code string) string {
	resp, err := http.PostForm(API_URL,	url.Values{
		"api_dev_key"    :   {APIKEY},
		"api_option"     :   {"paste"},
		"api_paste_code" :	 {code},
	})
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {
	content := extractContent(os.Args[1]) 
	fmt.Println(reqApi(content))
}
