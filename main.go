package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// "address": [{
// 	"line_1": "dhaskjdask",
// 	"line_2": "djaskdjask",
// 	"city": "Bangalore"
// }],

// output now
// {"email":"t*s*@*m*i*.*o*","message":"G*t*P*o*i*e*D*t*i*","name":"T*s*"}
var data = `{
	"name": "Test",
	"email": "test@gmail.com",
	
	"message": "Got Profile Detail"
}
`
var bytest []byte

func main() {
	var unmarshalVar map[string]interface{}
	err := json.Unmarshal([]byte(data), &unmarshalVar)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(partialMasking(unmarshalVar)))
}

func partialMasking(unmarshalVar map[string]interface{}) []byte {
	responseJson := map[string]interface{}{}
	for k, v := range unmarshalVar {
		switch v.(type) {
		case string:
			arrayStr := strings.Split(v.(string), "")
			result := strMask(arrayStr)
			responseJson[k] = result
			bytest, _ = json.Marshal(responseJson)
		}
	}
	return bytest
}

func strMask(s []string) string {
	var conditionMasking string
	var reformatStrArr []string
	for k := range s {
		if (k+1)%2 == 0 {
			conditionMasking = "*"
		} else {
			conditionMasking = s[k]
		}
		reformatStrArr = append(reformatStrArr, conditionMasking)
		conditionMasking = ""
	}
	reformatStr := strings.Join(reformatStrArr, "")
	return reformatStr
}
