package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	agregarEjercicioACategoria("Flexiones", "Fuerza", 5, "Media", 10, "brazos a 90")
	agregarEjercicioACategoria("Sentadillas", "Cardio", 20, "Baja", 15, "bajar")

	bluePrintf("\n\n\n\n\n\n\nBienvenido a tu gestor de rutinas de entrenamiento personalizado")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		greenPrintf("\n¿Qué te gustaría hacer?\n")
		bluePrintf("1.")
		fmt.Print("Nueva Rutina\n")
		bluePrintf("2.")
		fmt.Print("Agregar Ejercicio\n")
		bluePrintf("3.")
		fmt.Print("Listar Categorias\n")
		bluePrintf("4.")
		fmt.Print("Consultar Rutina Creada\n")
		bluePrintf("5.")
		fmt.Print("Modificar Rutina Creada\n")
		bluePrintf("6.")
		fmt.Print("Salir\n\n\n")
		scanner.Scan()
		comando := scanner.Text()
		comando = strings.ReplaceAll(comando, " ", "")

		switch strings.ToLower(comando) {
		case "nuevarutina":
			rutina := crearRutinaPersonalizada()
			fmt.Printf("Rutina creada:\n %+v\n", rutina)
		case "agregarejercicio":
			solicitarYAgregarEjercicio()
		case "listarcategorias":
			for categoria, ejercicios := range categorias {
				fmt.Printf("Categoría: %s, Ejercicios disponibles: %d\n", categoria, len(ejercicios))
				for _, ejercicio := range ejercicios {
					fmt.Printf("    - %s\n", ejercicio.Nombre)
				}
			}
		case "consultarrutinacreada":
			if len(rutinasL) == 0 {
				redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
				continue
			}
			greenPrintf("Selecciona la rutina que queres ver:\n")
			for i, v := range rutinasL {
				rutinaActual := v
				bluePrintf("%v. ", i+1)
				fmt.Printf("%v\n", rutinaActual.NombreDeRutina)
			}
			scanner.Scan()
			seleccionIn := scanner.Text()
			seleccion, _ := strconv.Atoi(seleccionIn)
			if len(rutinasL) < seleccion {
				redPrintf("La rutina seleccionada no existe")
				continue
			}
			var rutina Rutina
			if seleccion == 0 {
				for i, nombre := range rutinasL {
					if nombre.NombreDeRutina == seleccionIn {
						rutina = consultaRutinaCreada(i)
					}
				}
			} else {
				rutina = consultaRutinaCreada(seleccion - 1)
			}
			bluePrintf("\nTitulo: ")
			fmt.Print(rutina.NombreDeRutina)
			bluePrintf("\nEjercicios: ")
			for i, v := range rutina.Ejercicios {
				redPrintf("%v: {", i)
				greenPrintf("Nombre: %v, Duracion: %v, Categoria: %v, Intensidad: %v, Descripcion: %v, Calorias: %v", v.Nombre, v.Duracion, v.Tipo, v.Intensidad, v.Calorias, v.Descripcion)
				redPrintf("}\n")
			}
			bluePrintf("\nDuracion total: ")
			fmt.Print(rutina.DuracionTotal)
			greenPrintf(" Minutos\n")

		case "modificarrutinacreada":
			if len(rutinasL) == 0 {
				redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
				continue
			}
			greenPrintf("Selecciona la rutina que queres modificar:\n")
			for i, v := range rutinasL {
				rutinaActual := v
				bluePrintf("%v. ", i+1)
				fmt.Printf("%v\n", rutinaActual.NombreDeRutina)
			}

			scanner.Scan()
			seleccionIn := scanner.Text()
			seleccion, _ := strconv.Atoi(seleccionIn)
			if len(rutinasL) < seleccion {
				redPrintf("La rutina seleccionada no existe")
				continue
			}
			if seleccion == 0 {
				for i, nombre := range rutinasL {
					if nombre.NombreDeRutina == seleccionIn {
						modificarRutina(i)
					}
				}

			} else {
				modificarRutina(seleccion - 1)
			}
		case "salir":
			fmt.Println("Gracias por usar el gestor de rutinas. ¡Hasta pronto!")
			return
		case "1":
			rutina := crearRutinaPersonalizada()
			greenPrintf("Rutina creada: %+v\n", rutina)
			bluePrintf("Para consultar esta y otras rutinas seleccione 4\n")
		case "2":
			solicitarYAgregarEjercicio()
		case "3":
			for categoria, ejercicios := range categorias {
				fmt.Printf("Categoría: %s, Ejercicios disponibles: %d\n", categoria, len(ejercicios))
				for _, ejercicio := range ejercicios {
					fmt.Printf("    - %s\n", ejercicio.Nombre)
				}
			}
		case "4":
			if len(rutinasL) == 0 {
				redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
				continue
			}
			greenPrintf("Selecciona la rutina que queres ver:\n")
			for i, v := range rutinasL {
				rutinaActual := v
				bluePrintf("%v. ", i+1)
				fmt.Printf("%v\n", rutinaActual.NombreDeRutina)
			}
			scanner.Scan()
			seleccionIn := scanner.Text()
			seleccion, _ := strconv.Atoi(seleccionIn)
			if len(rutinasL) < seleccion {
				redPrintf("La rutina seleccionada no existe")
				continue
			}
			var rutina Rutina
			if seleccion == 0 {
				for i, nombre := range rutinasL {
					if nombre.NombreDeRutina == seleccionIn {
						rutina = consultaRutinaCreada(i)
					}
				}
			} else {
				rutina = consultaRutinaCreada(seleccion - 1)
			}
			bluePrintf("\nTitulo: ")
			fmt.Print(rutina.NombreDeRutina)
			bluePrintf("\nEjercicios: ")
			for i, v := range rutina.Ejercicios {
				redPrintf("%v: {", i)
				greenPrintf("Nombre: %v, Duracion: %v, Categoria: %v, Intensidad: %v, Descripcion: %v, Calorias: %v", v.Nombre, v.Duracion, v.Tipo, v.Intensidad, v.Calorias, v.Descripcion)
				redPrintf("}\n")
			}
			bluePrintf("\nDuracion total: ")
			fmt.Print(rutina.DuracionTotal)
			greenPrintf(" Minutos\n")
		case "5":
			if len(rutinasL) == 0 {
				redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
				continue
			}
			greenPrintf("Selecciona la rutina que queres modificar:\n")
			for i, v := range rutinasL {
				rutinaActual := v
				bluePrintf("%v. ", i+1)
				fmt.Printf("%v\n", rutinaActual.NombreDeRutina)
			}

			scanner.Scan()
			seleccionIn := scanner.Text()
			seleccion, _ := strconv.Atoi(seleccionIn)
			if len(rutinasL) < seleccion {
				redPrintf("La rutina seleccionada no existe")
				continue
			}
			if seleccion == 0 {
				for i, nombre := range rutinasL {
					if nombre.NombreDeRutina == seleccionIn {
						modificarRutina(i)
					}
				}

			} else {
				modificarRutina(seleccion - 1)
			}
		case "6":
			greenPrintf(" \n\n\n\n\n\n\n________________________________________\n")
			greenPrintf("/ Gracias por usar el gestor de rutinas. \\\n")
			greenPrintf("\\ ¡Hasta pronto!                          /\n")
			greenPrintf(" ----------------------------------------\n")
			bluePrintf("        \\   ^__^\n")
			bluePrintf("         \\  (oo)\\_______\n")
			bluePrintf("            (__)\\       )\\/\\\n")
			bluePrintf("                ||----w |\n")
			bluePrintf("                ||     ||\n")
			return
		default:
			fmt.Println("Comando no reconocido.")
		}
	}
}
