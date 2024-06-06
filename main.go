package main

import (
	"bufio"
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

	fmt.Print("Ingrese la cantidad de calorías quemadas: ")
	scanner.Scan()
	caloriasStr := scanner.Text()
	calorias, err := strconv.Atoi(caloriasStr)
	if err != nil {
		fmt.Println("Calorías inválidas. Por favor, ingrese un número válido.")
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
		case "1", "nombre":
			greenPrintf("\n\n\n\n\nEjercicios disponibles en la categoría %s:\n\n", categoria)
			for i, ejercicio := range ejercicios {
				bluePrintf("%d.", i+1)
				fmt.Printf(" %s", ejercicio.Nombre)
				redPrintf(" (Duración: %d minutos, Intensidad: %s)\n", ejercicio.Duracion, ejercicio.Intensidad)
			}
		case "2", "duracion":
			greenPrintf("\n\n\n\n\nEjercicios disponibles en la categoría %s:\n\n", categoria)
			for i, ejercicio := range ejercicios {
				bluePrintf("%d.", i+1)
				fmt.Printf(" %s", ejercicio.Nombre)
				redPrintf(" (Duración: %d minutos, Intensidad: %s)\n", ejercicio.Duracion, ejercicio.Intensidad)
			}
		case "3", "calorias":
			greenPrintf("\n\n\n\n\nEjercicios disponibles en la categoría %s:\n\n", categoria)
			for i, ejercicio := range ejercicios {
				bluePrintf("%d.", i+1)
				fmt.Printf(" %s", ejercicio.Nombre)
				redPrintf(" (Duración: %d minutos, Intensidad: %s)\n", ejercicio.Duracion, ejercicio.Intensidad)
			}
		default:
			fmt.Println("Opción inválida. Por favor, seleccione una opción válida.")
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

		fmt.Printf("Ejercicio '%s' agregado a la rutina '%s'. Duración total de la rutina: %d minutos.\n",
			ejercicioSeleccionado.Nombre, rutina.NombreDeRutina, rutina.DuracionTotal)
	}
	rutina.Id = strconv.Itoa(len(rutinasL) + 1)
	return &rutina
}

// Función para guardar una rutina en el archivo CSV.
func guardarRutina(rutina *Rutina) {
	// Agregar la rutina a la lista global de rutinas
	rutinasL = append(rutinasL, *rutina)

	// Guardar la lista actualizada en el archivo CSV
	file, err := os.OpenFile("rutinas.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	err = gocsv.MarshalFile(&rutinasL, file)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
	}
}

// Función para buscar una rutina por su nombre en el archivo CSV.
func buscarRutinaPorNombre(nombre string) (*Rutina, error) {
	nombre = strings.ToLower(nombre)
	nombre = strings.ReplaceAll(nombre, " ", "")

	file, err := os.OpenFile("rutinas.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()

	var rutinas []Rutina
	if err := gocsv.UnmarshalFile(file, &rutinas); err != nil {
		return nil, fmt.Errorf("error al leer el archivo CSV: %v", err)
	}

	for _, rutina := range rutinas {
		if rutina.NombreDeRutina == nombre {
			return &rutina, nil
		}
	}
	return nil, fmt.Errorf("no se encontró ninguna rutina con el nombre: %s", nombre)
}

func main() {
	// Ejemplo de uso
	// agregar ejercicios y categorías aquí

	scanner := bufio.NewScanner(os.Stdin)

	for {
		greenPrintf("\n\n\n\n\n\n\nSeleccione una opción:\n")
		bluePrintf("1. Agregar ejercicio\n")
		bluePrintf("2. Crear rutina personalizada\n")
		bluePrintf("3. Guardar rutina\n")
		bluePrintf("4. Buscar rutina por nombre\n")
		bluePrintf("5. Salir\n")
		greenPrintf("\n\n\n___________________________________\n")
		scanner.Scan()
		opcion := scanner.Text()

		switch opcion {
		case "1":
			solicitarYAgregarEjercicio()
		case "2":
			rutina := crearRutinaPersonalizada()
			greenPrintf("\n\nRutina '%s' creada con éxito.\n", rutina.NombreDeRutina)
		case "3":
			if len(rutinasL) == 0 {
				redPrintf("\n\nNo hay rutinas para guardar.\n")
			} else {
				guardarRutina(&rutinasL[len(rutinasL)-1])
				greenPrintf("\n\nRutina guardada con éxito.\n")
			}
		case "4":
			greenPrintf("\n\nIngrese el nombre de la rutina que desea buscar:\n")
			scanner.Scan()
			nombre := scanner.Text()
			rutina, err := buscarRutinaPorNombre(nombre)
			if err != nil {
				redPrintf("\n\nError: %v\n", err)
			} else {
				greenPrintf("\n\nRutina encontrada: %s\n", rutina.NombreDeRutina)
				for _, ejercicio := range rutina.Ejercicios {
					bluePrintf("- %s (Duración: %d minutos, Intensidad: %s)\n", ejercicio.Nombre, ejercicio.Duracion, ejercicio.Intensidad)
				}
				greenPrintf("Duración total: %d minutos\n", rutina.DuracionTotal)
			}
		case "5":
			return
		default:
			redPrintf("\n\nOpción inválida. Por favor, seleccione una opción válida.\n")
		}
	}
}
