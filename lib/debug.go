package lib

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(o any) {
	b, _ := json.Marshal(o)
	fmt.Println(string(b))
}
