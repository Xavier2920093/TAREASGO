package main

import "fmt"

func main() {
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	//mostrar matriz
	fmt.Print(matriz)
	sumaTotalFilas := sumarFilas(matriz)
	fmt.Println("Suma total de las filas es :", sumaTotalFilas)

	fmt.Println()

	sumatTotalesColumnas := sumarColumnas(matriz)
	fmt.Println("Suma total de las columnas es :", sumatTotalesColumnas)

	fmt.Printf("El resultado es %d x %d : %d\n", sumaTotalFilas, sumatTotalesColumnas, sumaTotalFilas*sumatTotalesColumnas)
}
