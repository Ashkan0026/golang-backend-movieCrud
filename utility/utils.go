package utility

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseRequestBody(r *http.Request, x interface{}) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, x)
	if err != nil {
		panic(err)
	}
}
