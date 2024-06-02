package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func agregarEjercicioARutina(rutina *Rutina, ejercicio Ejercicio) {
	rutina.Ejercicios = append(rutina.Ejercicios, ejercicio)
	rutina.DuracionTotal += ejercicio.Duracion
}

// Función para agregar un ejercicio a una rutina existente
func agregarEjercicioARutinaExistente(nombreDeRut int) {
	// // Verifica si hay rutinas creadas
	// if len(rutinasL) == 0 {
	// 	redPrintf("\n\nNo hay rutinas existentes para agregar ejercicios. Crea una rutina primero.\n\n")
	// 	return
	// }

	// // Listar rutinas disponibles para seleccionar
	// fmt.Println("\nSelecciona una rutina para agregar un ejercicio:")
	// for i, rutina := range rutinasL {
	// 	fmt.Printf("%d. %s (Duración total: %d minutos)\n", i+1, rutina.NombreDeRutina, rutina.DuracionTotal)
	// }

	// // Escanear la selección de rutina
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// seleccionStr := scanner.Text()
	// seleccion, err := strconv.Atoi(seleccionStr)
	// if err != nil || seleccion < 1 || seleccion > len(rutinasL) {
	// 	fmt.Println("Selección inválida.")
	// 	return
	// }

	// Obtener la rutina seleccionada
	rutinaSeleccionada := &rutinasL[nombreDeRut]

	// Listar categorías disponibles para seleccionar un ejercicio
	greenPrintf("\nSeleccione una categoría para elegir un ejercicio:\n")
	for categoria, ejercicios := range categorias {
		bluePrintf("- %s (%d ejercicios)\n", categoria, len(ejercicios))
	}

	// Escanear la selección de categoría
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	categoriaSeleccionada := scanner.Text()
	categoriaSeleccionada = strings.ToLower(categoriaSeleccionada)
	categoriaSeleccionada = strings.ReplaceAll(categoriaSeleccionada, " ", "")

	// Verificar si la categoría existe
	ejercicios, ok := categorias[categoriaSeleccionada]
	if !ok {
		redPrintf("\nCategoría no válida.\n")
		return
	}

	// Listar ejercicios en la categoría seleccionada
	greenPrintf("\nSeleccione un ejercicio de la categoría %v : (inserte el numero de ejercicio)\n", categoriaSeleccionada)
	for i, ejercicio := range ejercicios {
		bluePrintf("%d.", i+1)
		fmt.Printf("%s (Duración: %d minutos, Intensidad: %s)\n", ejercicio.Nombre, ejercicio.Duracion, ejercicio.Intensidad)
	}

	// Escanear la selección de ejercicio
	scanner.Scan()
	seleccionEjercicioStr := scanner.Text()
	seleccionEjercicio, err := strconv.Atoi(seleccionEjercicioStr)
	if err != nil || seleccionEjercicio < 1 || seleccionEjercicio > len(ejercicios) {
		redPrintf("\nSelección de ejercicio inválida. Introduzca el numero de ejercicio correctamente\n")
		return
	}

	// Obtener el ejercicio seleccionado
	ejercicioSeleccionado := ejercicios[seleccionEjercicio-1]

	// Agregar el ejercicio a la rutina seleccionada
	agregarEjercicioARutina(rutinaSeleccionada, ejercicioSeleccionado)

	fmt.Printf("\nEjercicio '%s' agregado a la rutina '%s'. Duración total de la rutina: %d minutos.\n",
		ejercicioSeleccionado.Nombre, rutinaSeleccionada.NombreDeRutina, rutinaSeleccionada.DuracionTotal)
}
