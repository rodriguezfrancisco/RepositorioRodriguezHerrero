package main

import (
	"fmt"
	"sort"
)

func seleccionarEjerciciosParaTiempoMaximo(categoria string, tiempoDisponible int) []Ejercicio {
	ejercicios, ok := categorias[categoria]
	if !ok {
		fmt.Println("Categoría no válida.")
		return nil
	}

	// Ordenar los ejercicios por duración de menor a mayor
	sort.Slice(ejercicios, func(i, j int) bool {
		return ejercicios[i].Duracion < ejercicios[j].Duracion
	})

	var ejerciciosSeleccionados []Ejercicio
	var tiempoTotal int

	for _, ejercicio := range ejercicios {
		if tiempoTotal+ejercicio.Duracion <= tiempoDisponible {
			ejerciciosSeleccionados = append(ejerciciosSeleccionados, ejercicio)
			tiempoTotal += ejercicio.Duracion
		} else {
			break
		}
	}

	return ejerciciosSeleccionados
}

func agregarEjerciciosMaximosARutina(nombreDeRut int, categoria string, tiempoDisponible int) {
	rutinaSeleccionada := &rutinasL[nombreDeRut]

	ejerciciosSeleccionados := seleccionarEjerciciosParaTiempoMaximo(categoria, tiempoDisponible)
	if ejerciciosSeleccionados == nil {
		return
	}

	for _, ejercicio := range ejerciciosSeleccionados {
		agregarEjercicioARutina(rutinaSeleccionada, ejercicio)
	}

	fmt.Printf("\nSe han agregado %d ejercicios a la rutina '%s' de la categoría '%s' con un tiempo total de %d segundos.\n",
		len(ejerciciosSeleccionados), rutinaSeleccionada.NombreDeRutina, categoria, rutinaSeleccionada.DuracionTotal)
}
