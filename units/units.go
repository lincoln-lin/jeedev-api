package units

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"strconv"
)

func GetKey(args ...interface{}) (key string)  {
	s := ""
	for _,i := range args {
		switch i.(type) {
		case int:
			s += strconv.Itoa(i.(int))+":" //interface 转 int,int 再转string
			break
		case int64:
			s += strconv.FormatInt(i.(int64), 10) + ":"
			break
		case map[string]string:
			v := i.(map[string]string)
			s += Map2String(v) + ":"
			break
		case []string: //数组类型
			v := i.([]string)
			s += Array2String(v) + ":"
		default: //默认是字符串
			s += i.(string) + ":"
		}
	}
	//fmt.Println(s)
	return GetMd5(s)
}

func GetMd5(str string) (md5str string){
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x",has)
}
func Map2String(m map[string]string) (v string){
	mjson,_ := json.Marshal(m)
	return string(mjson)
}
func Array2String(a []string) (v string)  {
	ajson,_ := json.Marshal(a)
	return string(ajson)
}

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
