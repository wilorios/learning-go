//MAPS
//ARE NOT COMPARABLE WITH == (only nil)
package main

import "fmt"

func main() {
	//map String keys and int values var nilMap map[string]int
	teams := map[string][]string{
		"America": {"Ramos", "Graterol"}, "Real": {"Ronaldo", "Casemiro"},
		"PSG": {"Neymar", "Mbape"}}
	//teams map[America:[Ramos Graterol] PSG:[Neymar Mbape Marquinos] ...
	//if you ask for unexisting key = empty
	fmt.Println("Millonarios", teams["Millonarios"])
	//Millonarios []
	//how to know if a map has a key?
	v, ok := teams["America"]
	fmt.Println(v, ok)
	//[Ramos Graterol] true
	teams2 := teams
	fmt.Println(teams2)
	teams2["Arsenal"] = []string{"Henry", "Lumberg"}
	//ADD Arsenal a ambos, comparten mem
	fmt.Println(teams)
	fmt.Println(teams2)
}
