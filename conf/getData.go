package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetData(path string) map[string]interface{} {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var result interface{}
	errJ := json.Unmarshal(data, &result)
	if errJ != nil {
		panic(errJ)
	}
	return result.(map[string]interface{})
}
