package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// expected output now
// {"address":[{"city":"C*t*","line_1":"d*a*k*d*s*","line_2":"d*a*k*j*s*"},{"city":"C*t* *e*t","line_1":"a*d*e*s*1","line_2":"a*d*e*s*2"}],"address2":[{"city":"P*n*a*","line_1":"a*d*e*s* *h*s* *","line_2":"a*d*e*s* *o*e*a*d*e*s"},{"city":"c*d*n* *i*y","line_1":"d*m* *d*r*s*","line_2":"t*s*i*g*a*d*e*s"}],"email":"t*s*@*m*i*.*o*","message":"G*t*P*o*i*e*D*t*i*","name":"T*s*"}
var data = `{
	"name": "Test",
	"email": "test@gmail.com",
	"address": [{
		"line_1": "dhaskjdask",
		"line_2": "djaskdjask",
		"city": "City"
	},{
		"line_1": "address 1",
		"line_2": "address 2",
		"city": "City Test"
	}],
	"address2": [{
		"line_1": "address2 phase 7",
		"line_2": "address2 some address",
		"city": "Punjab"
	},{
		"line_1": "demo address",
		"line_2": "testing address",
		"city": "coding city"
	}],	
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
	resp := partialMasking(unmarshalVar)
	bytest, _ = json.Marshal(resp)
	fmt.Println(string(bytest))
}

func partialMasking(unmarshalVar map[string]interface{}) map[string]interface{} {
	responseJson := map[string]interface{}{}
	nestedJson := []map[string]interface{}{}
	maskedString := ""
	for k, v := range unmarshalVar {
		switch v.(type) {
		// handling string data
		case string:
			arrayStr := strings.Split(v.(string), "")
			maskedString = strMask(arrayStr)
			responseJson[k] = maskedString
		// handling if there is any multiple data
		case []interface{}:
			for _, va := range v.([]interface{}) {
				resp := partialMasking(va.(map[string]interface{}))
				nestedJson = append(nestedJson, resp)
				responseJson[k] = nestedJson
			}
			nestedJson = nil
		}
	}
	return responseJson
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
