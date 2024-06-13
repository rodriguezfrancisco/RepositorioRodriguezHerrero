package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	gocsv "github.com/gocarina/gocsv"
)

func main() {
	agregarEjercicioACategoria("Flexiones", "Fuerza", 30, "Media", 10, "brazos a 90", 30, 5, 80)
	agregarEjercicioACategoria("Salto estrella", "Cardio", 35, "Baja", 15, "saltar haciendo estrellas", 70, 20, 5)
	agregarEjercicioACategoria("Estiramiento de triceps", "Flexibilidad", 20, "Baja", 15, "poner un brazo atras de la cabeza con el codo apuntando al techo y tirarlo para abajo", 0, 100, 0)
	agregarEjercicioACategoria("Biceps", "Fuerza", 20, "Media", 15, "levantar las pesas", 15, 0, 80)
	agregarEjercicioACategoria("Sentadilla", "Fuerza", 15, "Media", 15, "bajar", 40, 10, 70)
	agregarEjercicioACategoria("Press banca", "Fuerza", 25, "Alta", 28, "acostarse boca arriba y levantar las pesas", 20, 0, 90)
	agregarEjercicioACategoria("Estocadas", "Cardio", 25, "Baja", 15, "poner una pierna adelante y bajar", 65, 60, 40)
	agregarEjercicioACategoria("Espagat", "Flexibilidad", 20, "Alta", 15, "Tijeras en el piso", 10, 100, 0)

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
		fmt.Print("Agregar Ejercicios Máximos a una Rutina\n")
		/*
			bluePrintf("7.")
			fmt.Print("Agregar Ejercicios Minimos a una Rutina\n")
		*/
		bluePrintf("7.")
		fmt.Print("Agregar Ejercicios Eficientes a una Rutina\n") // Nueva opción del menú
		bluePrintf("8.")
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
			rutinasFile, err := os.OpenFile("rutinas.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				panic(err)
			}
			defer rutinasFile.Close()
			rutinasCsv := []*RutinaCsv{}
			if err := gocsv.UnmarshalFile(rutinasFile, &rutinasCsv); err != nil { // Load rutinas from file
				panic(err)
			}
			if len(rutinasCsv) == 0 {
				redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
				continue
			}
			greenPrintf("Selecciona la rutina que queres ver:\n")
			for i, v := range rutinasCsv {
				rutinaActual := v
				bluePrintf("%v. ", i+1)
				fmt.Printf("%v\n", rutinaActual)
			}
			scanner.Scan()
			seleccionIn := scanner.Text()
			seleccion, _ := strconv.Atoi(seleccionIn)
			if len(rutinasCsv) < seleccion {
				redPrintf("La rutina seleccionada no existe")
				continue
			}
			var rutina string
			if seleccion == 0 {
				for i, nombre := range rutinasL {
					if nombre.NombreDeRutina == seleccionIn {
						rutina = consultaRutinaCreada(i)
					}
				}
			} else {
				rutina = consultaRutinaCreada(seleccion - 1)
			}
			bluePrintf("\nRutina: \n\n")
			greenPrintf(rutina)
			bluePrintf("\n")
		case "modificarrutinacreada":
			rutinasFile, err := os.OpenFile("rutinas.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				panic(err)
			}
			defer rutinasFile.Close()
			rutinasCsv := []*RutinaCsv{}
			if err := gocsv.UnmarshalFile(rutinasFile, &rutinasCsv); err != nil { // Load rutinas from file
				panic(err)
			}
			if len(rutinasCsv) == 0 {
				redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
				continue
			}
			greenPrintf("Selecciona la rutina que queres modificar:\n")
			for i, v := range rutinasCsv {
				rutinaActual := v
				bluePrintf("%v. ", i+1)
				fmt.Printf("%v\n", rutinaActual)
			}

			scanner.Scan()
			seleccionIn := scanner.Text()
			seleccion, _ := strconv.Atoi(seleccionIn)
			if len(rutinasCsv) < seleccion {
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

		case "agregarejerciciosmaximosarutina":
			// if len(rutinasL) == 0 {
			// 	redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
			// 	continue
			// }
			// greenPrintf("Selecciona la rutina a la que queres agregar ejercicios:\n")
			// for i, v := range rutinasL {
			// 	rutinaActual := v
			// 	bluePrintf("%v. ", i+1)
			// 	fmt.Printf("%v\n", rutinaActual.NombreDeRutina)
			// }
			// scanner.Scan()
			// seleccionIn := scanner.Text()
			// seleccion, err := strconv.Atoi(seleccionIn)
			// if err != nil || seleccion < 1 || seleccion > len(rutinasL) {
			// 	redPrintf("La rutina seleccionada no existe")
			// 	continue
			// }

			scanner := bufio.NewScanner(os.Stdin)
			// continuar := true

			greenPrintf("\n\n\n\n\n\n\nIngrese un nombre para la rutina:")
			bluePrintf("(ej: Rutina para Elias)\n\n\n\n\n___________________________________\n")

			scanner.Scan()
			nombre := scanner.Text()
			nombre = strings.ToLower(nombre)
			nombre = strings.ReplaceAll(nombre, " ", "")
			fmt.Println("Ingrese la categoría de ejercicios:")
			scanner.Scan()
			categoria := scanner.Text()

			fmt.Println("Ingrese el tiempo disponible en segundos:")
			scanner.Scan()
			tiempoIn := scanner.Text()
			tiempoDisponible, err := strconv.Atoi(tiempoIn)
			if err != nil || tiempoDisponible <= 0 {
				redPrintf("Tiempo no válido.")
				continue
			}

			agregarEjerciciosMaximosARutina(nombre /*seleccion-1,*/, categoria, tiempoDisponible)
			/*
				case "agregarejerciciosminimosarutina": // Nueva opción del menú
					if len(rutinasL) == 0 {
						redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
						continue
					}
					greenPrintf("Selecciona la rutina a la que queres agregar ejercicios:\n")
					for i, v := range rutinasL {
						rutinaActual := v
						bluePrintf("%v. ", i+1)
						fmt.Printf("%v\n", rutinaActual.NombreDeRutina)
					}
					scanner.Scan()
					seleccionIn := scanner.Text()
					seleccion, err := strconv.Atoi(seleccionIn)
					if err != nil || seleccion < 1 || seleccion > len(rutinasL) {
						redPrintf("La rutina seleccionada no existe")
						continue
					}

					fmt.Println("Ingrese la categoría de ejercicios:")
					scanner.Scan()
					categoria := scanner.Text()

					fmt.Println("Ingrese el tiempo disponible en segundos:")
					scanner.Scan()
					tiempoIn := scanner.Text()
					tiempoDisponible, err := strconv.Atoi(tiempoIn)
					if err != nil || tiempoDisponible <= 0 {
						redPrintf("Tiempo no válido.")
						continue
					}

					ejerciciosMin := seleccionarEjerciciosMinDuracion(rutinasL[seleccion-1], categoria, tiempoDisponible, "", "")
					if len(ejerciciosMin) == 0 {
						fmt.Println("No se encontraron ejercicios que satisfagan los criterios.")
					} else {
						fmt.Println("Ejercicios seleccionados:")
						for _, ejercicio := range ejerciciosMin {
							fmt.Printf("    - %s\n", ejercicio.Nombre)
						}
					}
			*/
		case "agregarejercicioseficientes":
			// Solicitar datos para agregar ejercicios eficientes a una rutina
			fmt.Print("Nombre de la rutina: ")
			scanner.Scan()
			nombreDeRutina := scanner.Text()

			fmt.Print("Categoría: ")
			scanner.Scan()
			categoria := scanner.Text()

			fmt.Print("Calorías deseadas: ")
			scanner.Scan()
			caloriasDeseadasStr := scanner.Text()
			caloriasDeseadas, err := strconv.Atoi(caloriasDeseadasStr)
			if err != nil {
				fmt.Println("Entrada no válida para calorías.")
				continue
			}

			agregarEjerciciosEficientesARutina(nombreDeRutina, categoria, caloriasDeseadas)
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
			rutinasFile, err := os.OpenFile("rutinas.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				panic(err)
			}
			defer rutinasFile.Close()
			rutinasCsv := []*RutinaCsv{}
			if err := gocsv.UnmarshalFile(rutinasFile, &rutinasCsv); err != nil { // Load rutinas from file
				panic(err)
			}
			if len(rutinasCsv) == 0 {
				redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
				continue
			}
			greenPrintf("Selecciona la rutina que queres ver:\n")
			for i, v := range rutinasCsv {
				rutinaActual := v
				bluePrintf("%v. ", i+1)
				fmt.Printf("%v\n", rutinaActual)
			}
			scanner.Scan()
			seleccionIn := scanner.Text()
			seleccion, _ := strconv.Atoi(seleccionIn)
			if len(rutinasCsv) < seleccion {
				redPrintf("La rutina seleccionada no existe")
				continue
			}
			var rutina string
			if seleccion == 0 {
				for i, nombre := range rutinasL {
					if nombre.NombreDeRutina == seleccionIn {
						rutina = consultaRutinaCreada(i)
					}
				}
			} else {
				rutina = consultaRutinaCreada(seleccion - 1)
			}
			bluePrintf("\nRutina: \n\n")
			greenPrintf(rutina)
			bluePrintf("\n")
		case "5":
			rutinasFile, err := os.OpenFile("rutinas.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				panic(err)
			}
			defer rutinasFile.Close()
			rutinasCsv := []*RutinaCsv{}
			if err := gocsv.UnmarshalFile(rutinasFile, &rutinasCsv); err != nil { // Load rutinas from file
				panic(err)
			}
			if len(rutinasCsv) == 0 {
				redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
				continue
			}
			greenPrintf("Selecciona la rutina que queres modificar:\n")
			for i, v := range rutinasCsv {
				rutinaActual := v
				bluePrintf("%v. ", i+1)
				fmt.Printf("%v\n", rutinaActual)
			}

			scanner.Scan()
			seleccionIn := scanner.Text()
			seleccion, _ := strconv.Atoi(seleccionIn)
			if len(rutinasCsv) < seleccion {
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
			// if len(rutinasL) == 0 {
			// 	redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
			// 	continue
			// }
			// greenPrintf("Selecciona la rutina a la que queres agregar ejercicios:\n")
			// for i, v := range rutinasL {
			// 	rutinaActual := v
			// 	bluePrintf("%v. ", i+1)
			// 	fmt.Printf("%v\n", rutinaActual.NombreDeRutina)
			// }
			// scanner.Scan()
			// seleccionIn := scanner.Text()
			// seleccion, err := strconv.Atoi(seleccionIn)
			// if err != nil || seleccion < 1 || seleccion > len(rutinasL) {
			// 	redPrintf("La rutina seleccionada no existe")
			// 	continue
			// }

			scanner := bufio.NewScanner(os.Stdin)
			// continuar := true

			greenPrintf("\n\n\n\n\n\n\nIngrese un nombre para la rutina:")
			bluePrintf("(ej: Rutina para Elias)\n\n\n\n\n___________________________________\n")

			scanner.Scan()
			nombre := scanner.Text()
			nombre = strings.ToLower(nombre)
			nombre = strings.ReplaceAll(nombre, " ", "")
			fmt.Println("Ingrese la categoría de ejercicios:")
			scanner.Scan()
			categoria := scanner.Text()

			fmt.Println("Ingrese el tiempo disponible en segundos:")
			scanner.Scan()
			tiempoIn := scanner.Text()
			tiempoDisponible, err := strconv.Atoi(tiempoIn)
			if err != nil || tiempoDisponible <= 0 {
				redPrintf("Tiempo no válido.")
				continue
			}

			agregarEjerciciosMaximosARutina(nombre /*seleccion-1,*/, categoria, tiempoDisponible)
		/*
			case "7": // Nueva opción del menú
				if len(rutinasL) == 0 {
					redPrintf("\n\n\nAun no agregaste ninguna rutina\n\n")
					continue
				}
				greenPrintf("Selecciona la rutina a la que queres agregar ejercicios:\n")
				for i, v := range rutinasL {
					rutinaActual := v
					bluePrintf("%v. ", i+1)
					fmt.Printf("%v\n", rutinaActual.NombreDeRutina)
				}
				scanner.Scan()
				seleccionIn := scanner.Text()
				seleccion, err := strconv.Atoi(seleccionIn)
				if err != nil || seleccion < 1 || seleccion > len(rutinasL) {
					redPrintf("La rutina seleccionada no existe")
					continue
				}

				fmt.Println("Ingrese la categoría de ejercicios:")
				scanner.Scan()
				categoria := scanner.Text()

				fmt.Println("Ingrese el tiempo disponible en segundos:")
				scanner.Scan()
				tiempoIn := scanner.Text()
				tiempoDisponible, err := strconv.Atoi(tiempoIn)
				if err != nil || tiempoDisponible <= 0 {
					redPrintf("Tiempo no válido.")
					continue
				}

				ejerciciosMin := seleccionarEjerciciosMinDuracion(rutinasL[seleccion-1], categoria, tiempoDisponible, "", "")
				if len(ejerciciosMin) == 0 {
					fmt.Println("No se encontraron ejercicios que satisfagan los criterios.")
				} else {
					fmt.Println("Ejercicios seleccionados:")
					for _, ejercicio := range ejerciciosMin {
						fmt.Printf("    - %s\n", ejercicio.Nombre)
					}
				}
		*/
		case "7": // Nueva opción en el menú

			// Solicitar datos para agregar ejercicios eficientes a una rutina
			fmt.Print("Nombre de la rutina: ")
			scanner.Scan()
			nombreDeRutina := scanner.Text()

			fmt.Print("Categoría: ")
			scanner.Scan()
			categoria := scanner.Text()

			fmt.Print("Calorías deseadas: ")
			scanner.Scan()
			caloriasDeseadasStr := scanner.Text()
			caloriasDeseadas, err := strconv.Atoi(caloriasDeseadasStr)
			if err != nil {
				fmt.Println("Entrada no válida para calorías.")
				continue
			}

			agregarEjerciciosEficientesARutina(nombreDeRutina, categoria, caloriasDeseadas)
		case "8":
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
