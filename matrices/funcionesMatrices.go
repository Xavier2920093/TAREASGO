package matrices

import "fmt"

func sumarFilas(matriz [][]int) int {
	sumaTotalFilas := 0
	for i := 0; i < len(matriz); i++ {
		sumaFilas := 0
		for j := 0; j < len(matriz[i]); j++ {
			sumaFilas += matriz[i][j]
		}
		sumaTotalFilas += sumaFilas
		fmt.Printf("la suma en la fila %d es : %d\n", (i + 1), sumaFilas)
	}
	return sumaTotalFilas
}

func sumarColumnas(matriz [][]int) int {
	sumatTotalesColumnas := 0
	for i := 0; i < len(matriz); i++ {
		sumacolumnaz := 0
		for j := 0; j < len(matriz[i]); j++ {
			sumacolumnaz += matriz[j][i]
		}
		sumatTotalesColumnas += sumacolumnaz
		fmt.Printf("la suma en la columna %d es : %d\n", (i + 1), sumacolumnaz)
	}
	return sumatTotalesColumnas
}
