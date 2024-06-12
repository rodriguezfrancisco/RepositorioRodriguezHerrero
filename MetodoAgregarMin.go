package main

import (
	"fmt"
	"sort"
)

func seleccionarEjerciciosMinDuracion(rutina Rutina, categoria string, tiempoDisponible int, dificultad string, tipo string) []Ejercicio {
	ejercicios, ok := categorias[categoria]
	if !ok {
		fmt.Println("Categoría no válida.")
		return nil
	}

	// Filtrar ejercicios por tipo y dificultad
	var ejerciciosFiltrados []Ejercicio
	for _, ejercicio := range ejercicios {
		if ejercicio.Dificultad == dificultad && ejercicio.Tipo == tipo {
			ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
		}
	}

	// Ordenar los ejercicios por duración de menor a mayor
	sort.Slice(ejerciciosFiltrados, func(i, j int) bool {
		return ejerciciosFiltrados[i].Duracion < ejerciciosFiltrados[j].Duracion
	})

	var ejerciciosSeleccionados []Ejercicio
	var tiempoTotal int

	for _, ejercicio := range ejerciciosFiltrados {
		if tiempoTotal+ejercicio.Duracion <= tiempoDisponible {
			// Verificar si el ejercicio ya está en la rutina para evitar repeticiones
			if !ejercicioEnRutina(rutina, ejercicio) {
				ejerciciosSeleccionados = append(ejerciciosSeleccionados, ejercicio)
				tiempoTotal += ejercicio.Duracion
			}
		} else {
			break
		}
	}

	return ejerciciosSeleccionados
}

// Función auxiliar para verificar si un ejercicio ya está en la rutina
func ejercicioEnRutina(rutina Rutina, ejercicio Ejercicio) bool {
	for _, ej := range rutina.Ejercicios {
		if ej.Nombre == ejercicio.Nombre {
			return true
		}
	}
	return false
}
