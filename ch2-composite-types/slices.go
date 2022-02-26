//SLICES
//ARE NOT COMPARABLE WITH ==
//Slicing slice share mem sometimes
//(sometimes please view func slicesShareMem2)
//Slices en funciones permiten modificar
//contenido pero no permiten modificar el tamanio por ello no toma el append
package main

import "fmt"

func main() {
	x := []int{11, 12, 13, 14}
	y := x[:2] //[11 12]
	//la funcion es inicio:fin
	//si no se pone inicio toma 0
	//sino se pone fin toma el lenght
	//pos inicio es tomada en slide,
	//pos final no es tomada en slide
	z := x[1:]  //[12 13 14]
	d := x[1:3] //[12 13]
	e := x[:]   //[11 12 13 14]
	fmt.Println("0->", x, y, z, d, e)
	//0-> [11 12 13 14] [11 12] [12 13 14] [12 13] [11 12 13 14]

	slicesShareMem()
	fmt.Println("3->", x, y)
	//3-> [11 12 13 14] [11 12]

	slicesShareMem2(x)
	fmt.Println("7->", x, y)
	//7-> [1111 12 2000 14] [1111 12]

	//si no quieres que un slice comparta memoria usa el metodo copy
	num := copy(e, x)
	fmt.Println("8->", e, x, num)
	//8-> [1111 12 2000 14] [1111 12 2000 14] 4
}

func slicesShareMem() {
	x := []int{1, 2, 3, 4}
	y := x[:2] //[1 2]
	//si modifico la pos 0 de y //tambien modifica x
	y[0] = 11
	fmt.Println("1->", x, y)
	//1->[11 2 3 4] [11 2] ambos mod
	y = append(y, 1000)
	fmt.Println("2->", x, y)
	//2->[11 2 1000 4] [11 2 1000] //ambos mod adiciono 1000 a ambos fmt.Println(x, y)
}

func slicesShareMem2(x []int) {
	y := x[:2]
	//si modifico la pos 0 de y //tambien modifica x
	y[0] = 1111
	fmt.Println("4->", x, y)
	//4-> [1111 12 13 14] [1111 12] ambos mod
	x = append(x, 1000)
	fmt.Println("5->", x, y)
	//5-> [1111 12 13 14 1000] [1111 12] //solo lo adiciono a x porque tienen tamanios dif

	y = append(y, 2000)
	fmt.Println("6->", x, y)
	//6-> [1111 12 13 14 1000] [1111 12 2000] //solo lo adiciono a y
	//ojo pero cuando lo retorna modifica la pos 2 que tiene el 2000 de y en x
	//no envia el append.
	//7-> [1111 12 2000 14] [1111 12]

}
