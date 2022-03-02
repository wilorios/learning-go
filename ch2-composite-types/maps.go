//MAPS
//ARE NOT COMPARABLE WITH == (only nil)
//Maps dont share mem
//Maps,  the function change the 100% data
//hence we dont need to use the pointer reference
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
	//prints 1-> map[America:[Ramos Graterol] PSG:[Neymar Mbape]]
	teams2["Arsenal"] = []string{"Henry", "Lumberg"}
	//ADD Arsenal a ambos, comparten mem
	fmt.Println("2->", teams)
	//prints 2-> map[America:[Ramos Graterol] Arsenal:[Henry Lumberg]...
	fmt.Println("3->", teams2)
	//prints 2-> map[America:[Ramos Graterol] Arsenal:[Henry Lumberg]...
	addToMap(teams)
	fmt.Println("5->main ", teams)
	//prints 5->main  map[America:[Ramos Graterol] Arsenal:[Henry Lumberg] Liverpool:[Lucho Salah] PSG:[Neymar Mbape]]
	fmt.Println("6->main ", teams2)
	//prints 6->main  map[America:[Ramos Graterol] Arsenal:[Henry Lumberg] Liverpool:[Lucho Salah] PSG:[Neymar Mbape]]

	//this is th way you can delete a key of a map
	delete(teams, "America")
	howInitializeMap()
}

func addToMap(teams map[string][]string) {
	teams["Liverpool"] = []string{"Lucho", "Salah"}
	fmt.Println("4->in func ", teams)
	//prints ->in func  map[America:[Ramos Graterol] Arsenal:[Henry Lumberg] Liverpool:[Lucho Salah] PSG:[Neymar Mbape]]
}

func howInitializeMap() {
	m := map[string]int{}
	m["key1"] = 2
	fmt.Println(m)
}
