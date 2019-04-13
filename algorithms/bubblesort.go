package main

import (
	"fmt"
)

func main() {

	var toBeSorted []int = []int{1, 3, 2, 4, 8, 0, 7, 2, 3, 6}
	toBeSorted = bubble(toBeSorted)

	fmt.Println(toBeSorted)

}

func bubble(arr []int) []int {
	swapped := true

	j := len(arr) - 1

	for swapped {
		swapped = false //flag para saber si se ha hecho al menos un swap, sino es que está ya ordenado

		i := 0
		for i = 0; i < j; i++ {
			if arr[i] > arr[i+1] {
				fmt.Println("SWAP", arr[i], arr[i+1])
				arr[i], arr[i+1] = arr[i+1], arr[i]

				swapped = true
			}
		}
		fmt.Println("ITERATIONS", i)
		j-- //como el valor mas alto SUBE (por eso "burbujea" hacia arriba) no es necesario iterar siempre toda la coleccion. Cada vez que se haga una iteracion completa por toda la coleccion, podremos omitir la ultima subiteracion
	}

	return arr
}
