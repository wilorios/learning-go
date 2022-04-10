package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := "1 2 3"
	// obtiene un slice de un string separado por espacios
	s := strings.Fields(in)
	fmt.Printf("Example 1: %v\n", s)

	//CONVIERTE DE STRING A ENTERO
	ci, _ := strconv.Atoi("10")
	fmt.Printf("Example 2: %d\n", ci)

	//ORDENA SLICES
	x := []int{11, 8, 9}
	x = append(x, 1)
	sort.Ints(x)
	fmt.Printf("Example 3: %v\n", x)

	//ORDENA STRING Y QUITA DUPLICADOS
	fmt.Printf("Example 4: %s\n ", TwoToOne("zzzzaaaaxxxxbbb", "wwwwkkkmiugswascfgyvfdaaa"))
	//prints abcdfgikmsuvwxyz
	fmt.Printf(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
	//prints abigailtheta

	DigitalRoot(589)
}

func TwoToOne(s1 string, s2 string) string {
	var r []string
	//convierte un string en slices
	slices := strings.Split(s1, "")
	//adiciona valores a un slice
	slices = append(slices, strings.Split(s2, "")...)

	//inicializa un map de booleans
	m := make(map[string]bool)

	//itero los slices
	for _, v := range slices {
		//verifico si la key del mapa existe
		if _, ok := m[v]; !ok {
			//si no existe adiciono key + bool al mapa
			m[v] = true
			//adiciono valores al return
			r = append(r, v)
		}
	}
	//ordeno el array de strings
	sort.Strings(r)
	//convierto el slice a string
	return strings.Join(r, "")
}

func LongestConsec(strarr []string, k int) string {
	var a, r string
	for i := 0; i < len(strarr)-k+1; i++ {
		a = strings.Join(strarr[i:i+k], "")
		if len(a) > len(r) {
			r = a
		}
	}
	return r
}

func DigitalRoot(n int) int {
	a := strings.Split(strconv.Itoa(n), "")
	var r int
	for i := 0; i < len(a); i++ {
		ci, _ := strconv.Atoi(a[i])
		r = r + ci
		if i == (len(a)-1) && r > 9 {
			i = -1
			a = strings.Split(strconv.Itoa(r), "")
			r = 0
		}
	}
	return r
}
