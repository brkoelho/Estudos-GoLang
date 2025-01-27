package main

import (
	"fmt"
)

func main() {
	// Exemplo Padrão - ele adota, como não foi definido, ZeroValue de 0 em K
	var k int
	if k <= 0 {
		fmt.Println("Com K sendo ZeroValue=0 teremos o retorno:", k)
	} else {
		fmt.Println("Com K sendo ZeroValue=0 teremos o retorno:", -1*k)
	}
	// Agora adotando valores negativos
	z := -11
	if z <= 0 {
		fmt.Println("Com Z sendo de valor igual a:", z, "O retorno será:", z)
	} else {
		fmt.Println("Com Z sendo de valor igual a:", z, "O retorno será:", -1*z)
	}
	// Agora adotando valores positivos
	w := 22
	if w <= 0 {
		fmt.Println("Com W sendo de valor igual a:", w, "O retorno será:", w)
	} else {
		fmt.Println("Com W sendo de valor igual a:", w, "O retorno será:", -1*w)
	}
}

//// No site do codewars ficou assim:

// package kata

// func MakeNegative(x int) int {
//   if x<=0{
//   return x
//   } else {
//   return -x
//     }
// }
