package units

import (
	"encoding/json"
	"fmt"
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
