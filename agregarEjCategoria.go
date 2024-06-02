package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
