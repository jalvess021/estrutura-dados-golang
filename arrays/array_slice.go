package arrays

import (
	"fmt"
	"reflect"
)

//Fixo
var Arrays = [4]int{563, 251, 22, 63458}

//Sem pre-definicao de tamanho
var Slice = []string{"valor negativo", "sem valor", "empty", "null"}


func PrintArryAndSlice()  {
	fmt.Println("------------Slice--------------")
	checkArrayOrSlice(Slice)
	fmt.Println("------------Array--------------")
	checkArrayOrSlice(Arrays)
}

func PrintEvenNumbers(list []int){
	for _, v := range  removeOddNumbers(list) {
		fmt.Printf("%v\t\n", v)
	}
}

func checkArrayOrSlice(list interface{}){
	typeValue := reflect.ValueOf(list)
	if typeValue.Kind() == reflect.Array || typeValue.Kind() == reflect.Slice {
		for i := 0; i < typeValue.Len(); i++ {
			fmt.Printf("Chave: %d => valor: %v\n", i, typeValue.Index(i))
		}
	}
}

func removeOddNumbers(list []int) (evenNumbers []int)  {
	for _, v := range list {
		if v % 2 == 0{
			evenNumbers = append(evenNumbers, v)
		}
	}
	return
}



