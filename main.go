package main

import (
	"fmt"

	"github.com/Xavier2920093/TareasGO/matrices"
)

func main() {
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matriz:")
	matrices.MostrarMatriz(matriz)

	sumaTotalFilas := matrices.SumarFilas(matriz)
	fmt.Println("Suma total de las filas es:", sumaTotalFilas)

	sumaTotalColumnas := matrices.SumarColumnas(matriz)
	fmt.Println("Suma total de las columnas es:", sumaTotalColumnas)

	fmt.Printf("El resultado es %d x %d : %d\n", sumaTotalFilas, sumaTotalColumnas, sumaTotalFilas*sumaTotalColumnas)
}
