package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// Ejercicio representa un ejercicio individual en la rutina.
type Ejercicio struct {
	Nombre      string
	Duracion    int // Duración del ejercicio en minutos
	Tipo        string
	Intensidad  string
	Calorias    int
	Descripcion string
}

// Rutina representa una serie de ejercicios.
type Rutina struct {
	NombreDeRutina string
	Ejercicios     []Ejercicio
	DuracionTotal  int
}

// Categorías de ejercicios por tipo usando un mapa.
var categorias = make(map[string][]Ejercicio)
var rutinasL []Rutina
var redPrintf = color.New(color.FgRed).PrintfFunc()
var greenPrintf = color.New(color.FgGreen).PrintfFunc()
var bluePrintf = color.New(color.FgBlue).PrintfFunc()

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

// Función para agregar un ejercicio a una categoría.
func agregarEjercicioACategoria(nombre, tipo string, duracion int, intensidad string, calorias int, descripcion string) {
	nombre, tipo = strings.ToLower(nombre), strings.ToLower(tipo)
	nombre, tipo = strings.ReplaceAll(nombre, " ", ""), strings.ReplaceAll(tipo, " ", "")
	ej := Ejercicio{Nombre: nombre, Duracion: duracion, Tipo: tipo, Intensidad: intensidad, Calorias: calorias, Descripcion: descripcion}
	categorias[tipo] = append(categorias[tipo], ej)
	// fmt.Println(": ", categorias)
}

// Función para solicitar los detalles de un nuevo ejercicio y agregarlo a una categoría.
func solicitarYAgregarEjercicio() {
	scanner := bufio.NewScanner(os.Stdin)

	// Solicitar detalles del ejercicio
	fmt.Print("Ingrese el nombre del ejercicio: ")
	scanner.Scan()
	nombre := scanner.Text()

	fmt.Print("Ingrese la duración del ejercicio (en minutos): ")
	scanner.Scan()
	duracionStr := scanner.Text()
	duracion, err := strconv.Atoi(duracionStr)
	if err != nil {
		fmt.Println("Duración inválida. Por favor, ingrese un número válido.")
		return
	}

	fmt.Print("Ingrese el tipo de ejercicio (por ejemplo, Cardio, Fuerza): ")
	scanner.Scan()
	tipo := scanner.Text()

	fmt.Print("Ingrese la intensidad del ejercicio (por ejemplo, Baja, Media, Alta): ")
	scanner.Scan()
	intensidad := scanner.Text()

	fmt.Print("Ingrese la cantidad de calorias quemadas: ")
	scanner.Scan()
	caloriasStr := scanner.Text()
	calorias, err := strconv.Atoi(caloriasStr)
	if err != nil {
		fmt.Println("Descripcion invalida")
		return
	}

	fmt.Print("Ingrese la descripción del ejercicio: ")
	scanner.Scan()
	descripcion := scanner.Text()

	// Llamar a agregarEjercicioACategoria para agregar el ejercicio
	agregarEjercicioACategoria(nombre, tipo, duracion, intensidad, calorias, descripcion)

	// Confirmación
	fmt.Printf("Ejercicio '%s' agregado correctamente a la categoría '%s'.\n", nombre, tipo)
}

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

func main() {
	agregarEjercicioACategoria("Flexiones", "Fuerza", 5, "Media", 10, "brazos a 90")
	agregarEjercicioACategoria("Sentadillas", "Cardio", 20, "Baja", 15, "bajar")

	bluePrintf("\n\n\n\n\n\n\nBienvenido a tu gestor de rutinas de entrenamiento personalizado.")

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
