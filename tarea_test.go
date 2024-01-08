package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTarea(t *testing.T) {
	hecho := []struct {
		nombreCasoPrueba  string
		accionTarea       string
		tareasActuales    []tarea
		resultadoEsperado []tarea
	}{
		{
			nombreCasoPrueba: "Agregar una tarea",
			accionTarea:      "Agregar",
			tareasActuales:   []tarea{{nombre: "Desayunar", estado: false}},
			resultadoEsperado: []tarea{
				{
					nombre: "Desayunar",
					estado: false,
				},
				{
					nombre: "lavar",
					estado: false,
				}},
		},
		{
			nombreCasoPrueba:  "Completar una tarea",
			accionTarea:       "Completar",
			tareasActuales:    []tarea{{nombre: "Almorzar", estado: false}},
			resultadoEsperado: []tarea{{nombre: "Almorzar", estado: true}},
		},
		{
			nombreCasoPrueba:  "Imprimir una tarea",
			accionTarea:       "Imprimir",
			tareasActuales:    []tarea{{nombre: "Cenar", estado: false}},
			resultadoEsperado: []tarea{{nombre: "Cenar", estado: false}},
		},
	}

	for _, accion := range hecho {
		actividad := seleccionDeTarea(accion.tareasActuales, accion.accionTarea)
		assert.Equal(t, accion.resultadoEsperado, actividad)
	}
}

func TestTarea_Adicionar(t *testing.T) {

	tareasActuales := []tarea{{nombre: "Desayunar", estado: false}}
	resultadoEsperado := []tarea{
		{
			nombre: "Desayunar",
			estado: false,
		},
		{
			nombre: "lavar",
			estado: false,
		}}

	actividad := seleccionDeTarea(tareasActuales, "Agregar")
	assert.Equal(t, resultadoEsperado, actividad)
}
