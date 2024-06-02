package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func consultaRutinaCreada(n int) Rutina {
	return rutinasL[n]
}

func modificarRutina(nombreDeRut int) {
	for {
		greenPrintf("¿Qué modificación querés hacerle?\n")
		bluePrintf("1.")
		fmt.Print("Agregar ejercicio\n")
		bluePrintf("2.")
		fmt.Print("Editar ejercicio\n")
		bluePrintf("3.")
		fmt.Print("Eliminar ejercicio\n")
		bluePrintf("4.")
		fmt.Print("Volver\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		comando := scanner.Text()
		comando = strings.ReplaceAll(comando, " ", "")

		switch strings.ToLower(comando) {
		case "agregarejercicio":
			agregarEjercicioARutinaExistente(nombreDeRut)
		case "1":
			agregarEjercicioARutinaExistente(nombreDeRut)
		case "editarejercicio":
			fmt.Print("todavia no disponible")
			continue
		case "eliminarejercicio":
			greenPrintf("¿Qué ejercicio querés eliminar?(insertar el numero)\n")
			for i, rut := range rutinasL[nombreDeRut].Ejercicios {
				bluePrintf("%v.", i+1)
				fmt.Printf("%v\n", rut)
			}
			scanner.Scan()
			seleccionIn := scanner.Text()
			seleccion, _ := strconv.Atoi(seleccionIn)
			if len(rutinasL[nombreDeRut].Ejercicios) < seleccion || len(rutinasL[nombreDeRut].Ejercicios) == 0 {
				fmt.Println("El ejercicio seleccionado no existe")
				continue
			}
			if seleccion == 0 {
				redPrintf("Introduzca correctamente el numero de ejercicio\n")
				continue
			}
			rutinasL[nombreDeRut].DuracionTotal -= rutinasL[nombreDeRut].Ejercicios[seleccion-1].Duracion
			j := rutinasL[nombreDeRut].Ejercicios[:seleccion-1]
			j = append(j, rutinasL[nombreDeRut].Ejercicios[seleccion:]...)
			rutinasL[nombreDeRut].Ejercicios = j
			bluePrintf("Ejercicio eliminado\n")
			continue
		case "3":
			greenPrintf("¿Qué ejercicio querés eliminar?(insertar el numero)\n")
			for i, rut := range rutinasL[nombreDeRut].Ejercicios {
				bluePrintf("%v.", i+1)
				fmt.Printf("%v\n", rut)
			}
			scanner.Scan()
			seleccionIn := scanner.Text()
			seleccion, _ := strconv.Atoi(seleccionIn)
			if len(rutinasL[nombreDeRut].Ejercicios) < seleccion || len(rutinasL[nombreDeRut].Ejercicios) == 0 {
				fmt.Println("El ejercicio seleccionado no existe")
				continue
			}
			if seleccion == 0 {
				redPrintf("Introduzca correctamente el numero de ejercicio\n")
				continue
			}
			rutinasL[nombreDeRut].DuracionTotal -= rutinasL[nombreDeRut].Ejercicios[seleccion-1].Duracion
			j := rutinasL[nombreDeRut].Ejercicios[:seleccion-1]
			j = append(j, rutinasL[nombreDeRut].Ejercicios[seleccion:]...)
			rutinasL[nombreDeRut].Ejercicios = j
			bluePrintf("Ejercicio eliminado\n")
			continue
		case "volver":
			bluePrintf("Cambios guardados\n")
			return

		case "4":
			bluePrintf("Cambios guardados\n")
			return
		}
		return
	}
}
