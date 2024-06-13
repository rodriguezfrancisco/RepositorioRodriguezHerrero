package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreacionEjercicios(t *testing.T) {

	agregarEjercicioACategoria("Flexiones", "Fuerza", 30, "Media", 10, "brazos a 90", 30, 5, 80)
	_, esta := categorias["flexibilidad"]
	assert.False(t, esta)
	agregarEjercicioACategoria("Salto estrella", "Cardio", 35, "Baja", 15, "saltar haciendo estrellas", 70, 20, 5)
	agregarEjercicioACategoria("Espagat", "Flexibilidad", 20, "Alta", 16, "Tijeras en el piso", 10, 100, 0)
	_, esta = categorias["flexibilidad"]
	assert.True(t, esta)
	agregarEjercicioACategoria("Estiramiento de triceps", "Flexibilidad", 20, "Baja", 15, "poner un brazo atras de la cabeza con el codo apuntando al techo y tirarlo para abajo", 0, 100, 0)
	agregarEjercicioACategoria("Biceps", "Fuerza", 20, "Media", 15, "levantar las pesas", 15, 0, 80)
	agregarEjercicioACategoria("Sentadilla", "Fuerza", 15, "Media", 15, "bajar", 40, 10, 70)
	agregarEjercicioACategoria("Press banca", "Fuerza", 25, "Alta", 28, "acostarse boca arriba y levantar las pesas", 20, 0, 90)
	agregarEjercicioACategoria("Estocadas", "Cardio", 25, "Baja", 15, "poner una pierna adelante y bajar", 65, 60, 40)

	j := categorias["fuerza"][0]
	assert.Equal(t, j.Nombre, "Flexiones")
	assert.Equal(t, 5, j.PuntosFlexibilidad)

	j = categorias["fuerza"][3]
	assert.Equal(t, "acostarse boca arriba y levantar las pesas", j.Descripcion)

}

func TestRutinas(t *testing.T) {
	var rutina Rutina
	rutina.NombreDeRutina = "Juan"
	e1 := categorias["fuerza"][0]
	e2 := categorias["fuerza"][3]
	agregarEjercicioARutina(&rutina, e1)
	agregarEjercicioARutina(&rutina, e2)
	ej := rutina.Ejercicios[1]
	assert.Equal(t, "Press banca", ej.Nombre)
	assert.Equal(t, 170, rutina.PuntosFuerzaTotales)
	assert.Equal(t, 55, rutina.DuracionTotal)

	rutinasL = append(rutinasL, rutina)

	var rutinaDos Rutina
	rutinaDos.NombreDeRutina = "Ale"
	e3 := categorias["cardio"][0]
	e4 := categorias["flexibilidad"][1]
	agregarEjercicioARutina(&rutinaDos, e3)
	agregarEjercicioARutina(&rutinaDos, e4)

	rutinasL = append(rutinasL, rutinaDos)

}

func TestEjerciciosMaximos(t *testing.T) {

	// agregarEjercicioACategoria("Flexiones", "Fuerza", 30, "Media", 10, "brazos a 90", 30, 5, 80)
	// agregarEjercicioACategoria("Salto estrella", "Cardio", 35, "Baja", 15, "saltar haciendo estrellas", 70, 20, 5)
	// agregarEjercicioACategoria("Espagat", "Flexibilidad", 20, "Alta", 15, "Tijeras en el piso", 10, 100, 0)
	// agregarEjercicioACategoria("Estiramiento de triceps", "Flexibilidad", 20, "Baja", 15, "poner un brazo atras de la cabeza con el codo apuntando al techo y tirarlo para abajo", 0, 100, 0)
	// agregarEjercicioACategoria("Biceps", "Fuerza", 20, "Media", 15, "levantar las pesas", 15, 0, 80)
	// agregarEjercicioACategoria("Sentadilla", "Fuerza", 15, "Media", 15, "bajar", 40, 10, 70)
	// agregarEjercicioACategoria("Press banca", "Fuerza", 25, "Alta", 28, "acostarse boca arriba y levantar las pesas", 20, 0, 90)
	// agregarEjercicioACategoria("Estocadas", "Cardio", 25, "Baja", 15, "poner una pierna adelante y bajar", 65, 60, 40)

	ejercicios := seleccionarEjerciciosParaTiempoMaximo("fuerza", 50)

	assert.Equal(t, "Sentadilla", ejercicios[0].Nombre)
	assert.Equal(t, "Biceps", ejercicios[1].Nombre)
	assert.Equal(t, 2, len(ejercicios))

	ejercicios = seleccionarEjerciciosParaTiempoMaximo("cardio", 40)
	assert.Equal(t, "Estocadas", ejercicios[0].Nombre)
	assert.Equal(t, 1, len(ejercicios))
}

func TestEficiencia(t *testing.T) {
	ejercicios := seleccionarEjerciciosEficientes("fuerza", 40)
	assert.Equal(t, "Press banca", ejercicios[0].Nombre)
	assert.Equal(t, "Sentadilla", ejercicios[1].Nombre)
	assert.Equal(t, 2, len(ejercicios))

	ejercicios = seleccionarEjerciciosEficientes("flexibilidad", 40)
	assert.Equal(t, "Espagat", ejercicios[0].Nombre)
	assert.Equal(t, "Estiramiento de triceps", ejercicios[1].Nombre)

	assert.Equal(t, 2, len(ejercicios))

}
