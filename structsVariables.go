package main

import (
	"github.com/fatih/color"
	"os"
	
	

	gocsv "github.com/gocarina/gocsv"
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

type RutinaCsv struct {
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
