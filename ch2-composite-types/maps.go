//MAPS
//ARE NOT COMPARABLE WITH == (only nil)

package main

import "fmt"

func main() {
	//map String keys and int values var nilMap map[string]int
	teams := map[string][]string{
		"America": {"Ramos", "Graterol"},
		"PSG":     {"Neymar", "Mbape"}}
	//teams map[America:[Ramos Graterol] PSG:[Neymar Mbape Marquinos]
	//if you ask for unexisting key = empty
	fmt.Println("Millonarios", teams["Millonarios"])
	//Millonarios []
	//how to know if a map has a key?
	v, ok := teams["America"]
	fmt.Println(v, ok)
	//[Ramos Graterol] true
	teams2 := teams
	fmt.Println("1->", teams2)
	teams2["Arsenal"] = []string{"Henry", "Lumberg"}
	//ADD Arsenal a ambos, comparten mem
	fmt.Println("2->", teams)
	fmt.Println("3->", teams2)

	addToMap(teams)
	fmt.Println("5->main ", teams)
	//prints
	fmt.Println("6->main ", teams2)

}

func addToMap(teams map[string][]string) {
	teams["Liverpool"] = []string{"Lucho", "Salah"}
	fmt.Println("4->in func ", teams)

}
