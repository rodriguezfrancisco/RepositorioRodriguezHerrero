package main

import (
	"bufio"
	//"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/gocarina/gocsv"
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
	Id             string      `csv:"client_id"`
	NombreDeRutina string      `csv:"client_nombre_de_rutina"`
	Ejercicios     []Ejercicio `csv:"client_ejercicios"`
	DuracionTotal  int         `csv:"client_duracion_total"`
}

// Categorías de ejercicios por tipo usando un mapa.
var categorias = make(map[string][]Ejercicio)
var rutinasL []Rutina
var redPrintf = color.New(color.FgRed).PrintfFunc()
var greenPrintf = color.New(color.FgGreen).PrintfFunc()
var bluePrintf = color.New(color.FgBlue).PrintfFunc()

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
func crearRutinaPersonalizada() *Rutina {
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
	rutina.Id = "5"

	fmt.Printf("La duración total de su rutina es de %d minutos.\n", rutina.DuracionTotal)
	rutinasL = append(rutinasL, rutina)
	fmt.Printf("\n\n\n\n RutinasL %+v", rutinasL)
	return &rutina
}

func consultaRutinaCreada(n int) Rutina {
	return rutinasL[n]
}

func main() {
	rutinasFile, err := os.OpenFile("rutinas.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer rutinasFile.Close()
	rutinas := []*Rutina{}
	if err := gocsv.UnmarshalFile(rutinasFile, &rutinas); err != nil { // Load rutinas from file
		panic(err)
	}
	// for _, rutina := range rutinas {
	// 	greenPrintf("Id %v\n", rutina.Id)
	// 	greenPrintf("Nombre de rutina: %v\n", rutina.NombreDeRutina)
	// 	greenPrintf("Ejercicios:  %v\n", rutina.Ejercicios)
	// 	greenPrintf("Duracion Total:  %v\n", rutina.DuracionTotal)
	//}

	if _, err := rutinasFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	// rutinas = append(rutinas, &Rutina{Id: "1", NombreDeRutina: "Rutina de Walter", Ejercicios: []Ejercicio{{Nombre: "flexiones", Duracion: 5, Tipo: "fuerza", Intensidad: "Media", Calorias: 10, Descripcion: "brazos a 90"}}, DuracionTotal: 10}) // Add rutinas
	// rutinas = append(rutinas, &Rutina{Id: "2", NombreDeRutina: "Rutina de Diame", Ejercicios: []Ejercicio{{Nombre: "flexiones", Duracion: 5, Tipo: "fuerza", Intensidad: "Media", Calorias: 10, Descripcion: "brazos a 90"}}, DuracionTotal: 10})
	// rutinas = append(rutinas, &Rutina{Id: "3", NombreDeRutina: "Rutina de Fran", Ejercicios: []Ejercicio{{Nombre: "flexiones", Duracion: 5, Tipo: "fuerza", Intensidad: "Media", Calorias: 10, Descripcion: "brazos a 90"}}, DuracionTotal: 10})
	// rutinas = append(rutinas, &Rutina{Id: "4", NombreDeRutina: "Rutina de Alexis", Ejercicios: []Ejercicio{{Nombre: "flexiones", Duracion: 5, Tipo: "fuerza", Intensidad: "Media", Calorias: 10, Descripcion: "brazos a 90"}}, DuracionTotal: 10})
	// csvContent, err2 := gocsv.MarshalString(&rutinas) // Get all rutinas as CSV string
	//err = gocsv.MarshalFile(&rutinas, rutinasFile)    // Use this to save the CSV back to the file
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Println(csvContent) // Display all rutinas as CSV string
	//reader := csv.NewReader(strings.NewReader(csvContent))

	// Leer todas las filas
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }

	// Mostrar la fila de consulta
	// consulta := records[2]
	// redPrintf("Fila de consulta:", consulta)
	// Acceder a un valor específico en la fila de consulta
	// consultita := consulta[2] 	// La columna 6 tiene el valor de las calorías
	// bluePrintf("\nColumna:", consultita)
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
		fmt.Print("Salir\n\n\n")
		scanner.Scan()
		comando := scanner.Text()
		comando = strings.ReplaceAll(comando, " ", "")

		switch strings.ToLower(comando) {
		case "nuevarutina":
			rutina := crearRutinaPersonalizada()
			rutinas = append(rutinas, rutina)
			err = gocsv.MarshalFile(&rutinas, rutinasFile)
			if err != nil {
				panic(err)
			}
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
			fmt.Printf("Selecciona la rutina que queres ver:\n")
			for i, v := range rutinasL {
				rutinaActual := v
				fmt.Printf("%v. %v\n", i+1, rutinaActual.NombreDeRutina)
			}
			scanner.Scan()
			seleccionIn := scanner.Text()
			seleccion, _ := strconv.Atoi(seleccionIn)
			if len(rutinasL) < seleccion {
				fmt.Println("La rutina seleccionada no existe")
				continue
			}
			rutina := consultaRutinaCreada(seleccion - 1)
			fmt.Printf("datos: %v", rutina)
		case "salir":
			fmt.Println("Gracias por usar el gestor de rutinas. ¡Hasta pronto!")
			return
		case "1":
			rutina := crearRutinaPersonalizada()
			fmt.Printf("Rutina creada: %+v\n", rutina)
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
			rutina := consultaRutinaCreada(seleccion - 1)
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
