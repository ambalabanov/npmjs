package npmjs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

//Pkg := `{"total_rows":1115763,"offset":0,"rows":[{"id":"0","key":"0","value":{"rev":"1-5fbff37e48e1dd03ce6e7ffd17b98998"}}]}`
type Pkg struct {
	Total  int `json:"total_rows"`
	Offset int `json:"offset"`
	Row    []struct {
		ID  string `json:"id"`
		Key string `json:"key"`
		Val struct {
			Rev string `json:"rev"`
		} `json:"value"`
	} `json:"rows"`
}

// GetRegistry for node.js packages registry
func GetRegistry() *Pkg {
	var p = &Pkg{}
	url := "https://replicate.npmjs.com/_all_docs"

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	jerr := json.Unmarshal([]byte(body), &p)
	if jerr != nil {
		panic(jerr)
	}
	return p
}

//WriteFile for write file
func WriteFile(p *Pkg, n string) {
	file, err := os.Create(n)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	for i := 0; i < len(p.Row); i++ {
		file.WriteString(p.Row[i].ID + "\n")
	}
}

// Get - main func
func Get(filename string) {
	p := GetRegistry()
	WriteFile(p, filename)
}
