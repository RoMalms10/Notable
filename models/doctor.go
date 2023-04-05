package models

type Doctor struct {
	UUID     int           `json:"id"`
	Name     string        `json:"name"`
	Calendar []Appointment `json:"calendar"`
}

//serialize the Person struct as JSON
// jsonBytes, err := json.Marshal(p)
// if err != nil {
//     // handle error
// }
// fmt.Println(string(jsonBytes)) // {"name":"Alice","age":30}
