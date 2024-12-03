package main

import (
	//"log"
	"encoding/json"
	"fmt"
	"os"
)

// ///
type Course struct {
	Name           string `json:"name"`
	StudentsNumber int    `json: studentsNumber`
	Chapters       string "json: Chapters"
}

func main() {
	file, err := os.ReadFile("example.json") //////////чтение Json файла
	if err != nil {
		fmt.Println("error openin file", err)
	}

	parseJSON(file)
	//defer file.Close()
}

func parseJSON(fileJson []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(fileJson, &data)

	if err != nil {
		fmt.Printf("Couldnt open %s\n", err)
		return
	}
	fmt.Printf("Course %v\n", data)

	// // Преобразуем map в структуру
	// var result Course
	// for k, v := range data {
	// 	fmt.Println(k, v)
	// }

	// if k, ok := data["name"].(string); ok {
	// 	result.Name = k
	// 	fmt.Println(result)
	// }

}

///
