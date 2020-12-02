package main

import(
	"fmt"
	)


func main(){
	var lista []int
	lista = desglosar(9,lista)//Se desglosa la lista
	lista=invertir(lista)//Se invierte la lista desglosada

	fmt.Println(lista)

}


func desglosar(n int,lista[]int) []int{

	if n==0{
		return lista
	}else{
		lista=append(lista,n)
		return desglosar(n-1,lista)
	}
}


func invertir(lista []int) []int {
	if len(lista) == 0 {
		return lista
	}
	return append(invertir(lista[1:]), lista[0])
}