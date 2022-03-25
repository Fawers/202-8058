package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	Id       int
	Teacher  string
	Students []string
}

type Course struct {
	Id      int
	Name    string
	Classes []Class
}

func main() {
	goCourse := Course{
		Id:   202,
		Name: "Go",
		Classes: []Class{
			{
				Id:      8058,
				Teacher: "Fabricio",
				Students: []string{
					"Jandson",
					"Elimar",
				},
			},
		},
	}

	jsonDoc, err := json.Marshal(&goCourse)
	if err != nil {
		fmt.Printf(
			"não foi possível serializar %v para json\nerro: %s\n", goCourse, err)
		return
	}

	fmt.Println(string(jsonDoc))

	var courseObj Course

	fmt.Printf("obj antes do unmarshal: %#v\n", courseObj)
	json.Unmarshal(jsonDoc, &courseObj)
	fmt.Printf("obj depois do unmarshal: %#v\n", courseObj)
}
