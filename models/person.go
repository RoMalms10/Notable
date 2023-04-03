package models

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//serialize the Person struct as JSON
// jsonBytes, err := json.Marshal(p)
// if err != nil {
//     // handle error
// }
// fmt.Println(string(jsonBytes)) // {"name":"Alice","age":30}
