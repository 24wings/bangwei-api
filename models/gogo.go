// gogo
package models

import (
	"encoding/json"
	//	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func connectToDb() {
	var result []bson.M
	conn, err := mgo.Dial("118.31.72.227/tutu")
	defer conn.Close()
	if err == nil {
		conn.DB("tutu").C("fenxiao-users").Find(&result)
		//		json.Unmarshal()
		if err == nil {
			strBytes, _ := bson.Marshal(result)
			str, _ := json.Marshal(&strBytes)
			fmt.Println("results:", result, string(str))
		} else {
			fmt.Println("result:")
		}
	} else {
		fmt.Println("err", err.Error())
	}
}
