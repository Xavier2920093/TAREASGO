package matrices

import "fmt"

func MostrarMatriz(matriz [][]int) {
	for _, fila := range matriz {
		fmt.Println(fila)
	}
}

func SumarFilas(matriz [][]int) int {
	sumaTotalFilas := 0
	for i := 0; i < len(matriz); i++ {
		sumaFilas := 0
		for j := 0; j < len(matriz[i]); j++ {
			sumaFilas += matriz[i][j]
		}
		sumaTotalFilas += sumaFilas
		fmt.Printf("La suma en la fila %d es: %d\n", (i + 1), sumaFilas)
	}
	return sumaTotalFilas
}

func SumarColumnas(matriz [][]int) int {
	sumatTotalesColumnas := 0
	for i := 0; i < len(matriz); i++ {
		sumacolumnaz := 0
		for j := 0; j < len(matriz[i]); j++ {
			sumacolumnaz += matriz[j][i]
		}
		sumatTotalesColumnas += sumacolumnaz
		fmt.Printf("La suma en la columna %d es: %d\n", (i + 1), sumacolumnaz)
	}
	return sumatTotalesColumnas
}
