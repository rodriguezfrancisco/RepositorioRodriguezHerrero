package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Función para crear una rutina personalizada.
func crearRutinaPersonalizada() Rutina {
	var rutina Rutina

	scanner := bufio.NewScanner(os.Stdin)
	continuar := true

	greenPrintf("\n\n\n\n\n\n\nIngrese un nombre para la rutina:")
	bluePrintf("(ej: Rutina para Elias)\n\n\n\n\n___________________________________\n")

	scanner.Scan()
	rutina.NombreDeRutina = scanner.Text()
	rutina.NombreDeRutina = strings.ToLower(rutina.NombreDeRutina)
	rutina.NombreDeRutina = strings.ReplaceAll(rutina.NombreDeRutina, " ", "")

	greenPrintf("\n\n\n\n\n\n\nCategorías disponibles:")
	for categoria, ejercicios := range categorias {
		bluePrintf("\n- %s ", strings.Title(categoria))
		redPrintf("(%d ejercicios)\n", len(ejercicios))
	}

	for continuar {
		greenPrintf("\n\n\n\n\n\n\nCategorías disponibles:")
		for categoria, ejercicios := range categorias {
			bluePrintf("\n- %s ", strings.Title(categoria))
			redPrintf("(%d ejercicios)\n", len(ejercicios))
		}

		greenPrintf("\n\n\nSeleccione una categoría para agregar a su rutina")
		bluePrintf("(ej: 'fuerza')")
		greenPrintf(" o escriba ")
		bluePrintf("'listo' ")
		greenPrintf("para finalizar: ")
		bluePrintf("\n\n\n___________________________________\n")
		scanner.Scan()
		categoria := scanner.Text()
		categoria = strings.ReplaceAll(categoria, " ", "")
		categoria = strings.ToLower(categoria)

		if categoria == "listo" {
			fmt.Println("\n\n\n\n\n\n\n\n\n\n\n.")
			continuar = false
			continue
		}

		ejercicios, ok := categorias[categoria]
		if !ok {
			redPrintf("\n\n\n\nLa categoría seleccionada no existe.\n\n Las categorias disponibles son:")
			for categoria, ejercicios := range categorias {
				bluePrintf("\n- %s ", strings.Title(categoria))
				redPrintf("(%d ejercicios)", len(ejercicios))
			}
			continue
		}

		greenPrintf("\n\n\n\n\nOrganizar la lista de ejercicios por:\n 1.Nombre \n 2.Duracion\n 3.Calorias\n\n\n___________________________________\n")
		scanner.Scan()
		filtrado := scanner.Text()
		filtrado = strings.ReplaceAll(filtrado, " ", "")

		switch strings.ToLower(filtrado) {
		case "1":
			greenPrintf("\n\n\n\n\nEjercicios disponibles en la categoría %s:\n\n", categoria)
			for i, ejercicio := range ejercicios {
				bluePrintf("%d.", i+1)
				fmt.Printf(" %s", ejercicio.Nombre)
				redPrintf(" (Duración: %d minutos, Intensidad: %s)\n", ejercicio.Duracion, ejercicio.Intensidad)
			}
		case "2":
			greenPrintf("\n\n\n\n\nEjercicios disponibles en la categoría %s:\n\n", categoria)
			for i, ejercicio := range ejercicios {
				bluePrintf("%d.", i+1)
				fmt.Printf(" %s", ejercicio.Nombre)
				redPrintf(" (Duración: %d minutos, Intensidad: %s)\n", ejercicio.Duracion, ejercicio.Intensidad)
			}
		case "3":
			greenPrintf("\n\n\n\n\nEjercicios disponibles en la categoría %s:\n\n", categoria)
			for i, ejercicio := range ejercicios {
				bluePrintf("%d.", i+1)
				fmt.Printf(" %s", ejercicio.Nombre)
				redPrintf(" (Duración: %d minutos, Intensidad: %s)\n", ejercicio.Duracion, ejercicio.Intensidad)
			}
		case "nombre":
			greenPrintf("\n\n\n\n\nEjercicios disponibles en la categoría %s:\n\n", categoria)
			for i, ejercicio := range ejercicios {
				bluePrintf("%d.", i+1)
				fmt.Printf(" %s", ejercicio.Nombre)
				redPrintf(" (Duración: %d minutos, Intensidad: %s)\n", ejercicio.Duracion, ejercicio.Intensidad)
			}
		case "duracion":
			greenPrintf("\n\n\n\n\nEjercicios disponibles en la categoría %s:\n\n", categoria)
			for i, ejercicio := range ejercicios {
				bluePrintf("%d.", i+1)
				fmt.Printf(" %s", ejercicio.Nombre)
				redPrintf(" (Duración: %d minutos, Intensidad: %s)\n", ejercicio.Duracion, ejercicio.Intensidad)
			}
		case "calorias":
			greenPrintf("\n\n\n\n\nEjercicios disponibles en la categoría %s:\n\n", categoria)
			for i, ejercicio := range ejercicios {
				bluePrintf("%d.", i+1)
				fmt.Printf(" %s", ejercicio.Nombre)
				redPrintf(" (Duración: %d minutos, Intensidad: %s)\n", ejercicio.Duracion, ejercicio.Intensidad)
			}
		}

		greenPrintf("\n\nSeleccione el ejercicio que desea agregar a su rutina o escriba ")
		bluePrintf("(0)")
		greenPrintf(" para cambiar de categoría: ")
		bluePrintf("\n\n\n___________________________________\n")
		scanner.Scan()
		ejercicioIndexStr := scanner.Text()
		ejercicioIndex, err := strconv.Atoi(ejercicioIndexStr)
		if err != nil || ejercicioIndex < 0 || ejercicioIndex > len(ejercicios) {
			fmt.Println("Número de ejercicio inválido.")
			continue
		}
		if ejercicioIndex == 0 {
			continue
		}

		ejercicioSeleccionado := ejercicios[ejercicioIndex-1]
		rutina.Ejercicios = append(rutina.Ejercicios, ejercicioSeleccionado)
		rutina.DuracionTotal += ejercicioSeleccionado.Duracion

		fmt.Printf("\n\n\n\n\nSe ha agregado '%s' a su rutina.\n", ejercicioSeleccionado.Nombre)
	}

	fmt.Printf("La duración total de su rutina es de %d minutos.\n", rutina.DuracionTotal)
	rutinasL = append(rutinasL, rutina)
	//fmt.Printf("\n\n\n\n RutinasL %+v", rutinasL)
	return rutina
}
