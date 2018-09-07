package units

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Bytes2Intaface(value [] byte) (v [] interface{}) {
	errShal := json.Unmarshal(value, &v)
	if errShal != nil {
		fmt.Println(errShal)
	}
	return v
}
func Bytes2Map(value [] byte) (v map[string]string) {
	errShal := json.Unmarshal(value,&v)
	if errShal != nil {
		fmt.Println(errShal)
	}
	return v
}

func HttpPostForm(url string, params map[string]interface{},header map[string]string ){
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
