package util

import (
	"encoding/json"
	"fmt"
)

func PrettyP(i interface{}) {
	s, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		fmt.Printf("failed pretty-printing: %v\n", err)
	}
	fmt.Println(string(s))
}
