package compare

import (
	"encoding/json"

	"github.com/valyala/fastjson"
)

type User struct {
	Nama string
	Umur int
}

func UnmarshallSTD(data string) {
	var hasil User
	json.Unmarshal([]byte(data), &hasil)
}

func UnmarshallFast(data string) {
	var p fastjson.Parser
	p.Parse(data)
}
