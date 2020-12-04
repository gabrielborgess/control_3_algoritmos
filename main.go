package main

import(
	"fmt"
	)

func main(){
	var lista []int
	lista = desglosar(9,lista)//Se desglosa la lista
	lista=invertir(lista)//Se invierte la lista desglosada

	fmt.Println(len(lista))

}

func desglosar(n int,lista[]int) []int{
	if n==0{
		return lista
	}else{
		lista=append(lista,n)
		return desglosar(n-1,lista)
	}
}


//func comprobar(contador int,lista[]int,listaaux[]int,k int) []int{
//	//El contador lleva la condicion de recursividad, la lista es la lista procesada del numero n dado, el k es el numero objetivo a ocmprobar y listaux es la lista donde se almacenara los conjuntos encontrados
//	largo:=len(lista)
//	if contador==largo-1{
//		return listaaux
//	}else{
//		for i := 1; i < largo-1; i++ {
//			temp:=lista[contador]+lista[i]+
//			if
//		}
//	}
//
//
//}
//



func comprobar(lista []int,numero int)[]int{
	largo:=len(lista)
	var listatemp []int
	contador:=0
	contadoraux:=1
	for contador< largo { 
		for contadoraux<largo{
			if lista[contador]+lista[contadoraux]==numero{
				lista=make([]int, 2)
				lista=append(lista,)
			}
		}
		contador += contador
	}
}

func invertir(lista []int) []int {
	if len(lista) == 0 {
		return lista
	}
	return append(invertir(lista[1:]), lista[0])
}