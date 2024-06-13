package main

import (
	"os"

	"github.com/fatih/color"
	gocsv "github.com/gocarina/gocsv"
)

// Ejercicio representa un ejercicio individual en la rutina.
type Ejercicio struct {
	Nombre             string
	Duracion           int // Duración del ejercicio en minutos
	Tipo               string
	Intensidad         string
	Calorias           int
	Descripcion        string
	Dificultad         string
	PuntosCardio       int
	PuntosFlexibilidad int
	PuntosFuerza       int
}

// Rutina representa una serie de ejercicios.
type Rutina struct {
	NombreDeRutina            string
	Ejercicios                []Ejercicio
	DuracionTotal             int
	PuntosCardioTotales       int
	PuntosFlexibilidadTotales int
	PuntosFuerzaTotales       int
}

type RutinaCsv struct {
	Id                        string      `csv:"client_id"`
	NombreDeRutina            string      `csv:"client_nombre_de_rutina"`
	Ejercicios                []Ejercicio `csv:"client_ejercicios"`
	DuracionTotal             int         `csv:"client_duracion_total"`
	PuntosCardioTotales       int         `csv:"client_puntos_cardio_totales"`
	PuntosFlexibilidadTotales int         `csv:"client_puntos_flexibilidad_totales"`
	PuntosFuerzaTotales       int         `csv:"client_puntos_fuerza_totales"`
}

// Categorías de ejercicios por tipo usando un mapa.
var categorias = make(map[string][]Ejercicio)
var rutinasL []Rutina
var redPrintf = color.New(color.FgRed).PrintfFunc()
var greenPrintf = color.New(color.FgGreen).PrintfFunc()
var bluePrintf = color.New(color.FgBlue).PrintfFunc()

func guardarRutinasEnCSV(rutinas []*Rutina) error {
	rutinasFile, err := os.OpenFile("rutinas.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer rutinasFile.Close()

	if err := gocsv.MarshalFile(&rutinas, rutinasFile); err != nil {
		return err
	}

	return nil
}

// FUNCIONES:
// func agregarEjercicioARutina(rutina *Rutina, ejercicio Ejercicio)
// func agregarEjercicioARutinaExistente(nombreDeRut int)
// func agregarEjercicioACategoria(nombre, tipo string, duracion int, intensidad string, calorias int, descripcion string)
// func solicitarYAgregarEjercicio()
// func crearRutinaPersonalizada() Rutina
// func consultaRutinaCreada(n int) Rutina
// func modificarRutina(nombreDeRut int)
// func seleccionarEjerciciosParaTiempoMaximo(categoria string, tiempoDisponible int) []Ejercicio
// func agregarEjerciciosMaximosARutina(nombreDeRut int, categoria string, tiempoDisponible int)
// func seleccionarEjerciciosMinDuracion(rutina Rutina, categoria string, tiempoDisponible int, dificultad string, tipo string) []Ejercicio
// func ejercicioEnRutina(rutina Rutina, ejercicio Ejercicio) bool
// func (r *RutinaCsv) String() string
