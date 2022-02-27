package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}
type Manager struct {
	Employee //EMBEDDED FIELD
	Reports  []Employee
}

func main() {
	m := Manager{Employee: Employee{
		Name: "Bob Bobson",
		ID:   "12345"},
		Reports: []Employee{}}
	fmt.Println(m.ID)
	// prints 12345
	fmt.Println(m.Description())
	//prints Bob Bobson (12345)
}

//Description works for employee and manager
func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}
