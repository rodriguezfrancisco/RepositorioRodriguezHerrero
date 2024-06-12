package main

import (
	"fmt"
	"os"
	"strings"

	gocsv "github.com/gocarina/gocsv"
)

func (r *RutinaCsv) String() string {
	ejerciciosStr := make([]string, len(r.Ejercicios))
	for i, e := range r.Ejercicios {
		ejerciciosStr[i] = fmt.Sprintf("{Nombre: %s, Duracion: %d, Tipo: %s, Intensidad: %s, Calorias: %d, Descripcion: %s}",
			e.Nombre, e.Duracion, e.Tipo, e.Intensidad, e.Calorias, e.Descripcion)
	}
	return fmt.Sprintf("Id: %s, NombreDeRutina: %s, Ejercicios: [%s], DuracionTotal: %d",
		r.Id, r.NombreDeRutina, strings.Join(ejerciciosStr, ", "), r.DuracionTotal)
}

func consultaRutinaCreada(n int) string {
	rutinasFile, err := os.OpenFile("rutinas.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer rutinasFile.Close()
	rutinasCsv := []*RutinaCsv{}
	if err := gocsv.UnmarshalFile(rutinasFile, &rutinasCsv); err != nil { // Load rutinas from file
		panic(err)
	}
	rutinaString := rutinasCsv[n].String()
	return rutinaString
}
