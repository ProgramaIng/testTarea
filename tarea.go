package main

import "fmt"

type tarea struct {
	nombre string
	estado bool
}

func main() {
	sliceListaDeTarea := []tarea{
		{
			nombre: "Desayunar",
			estado: false,
		}, {
			nombre: "Almorzar",
			estado: false,
		}, {
			nombre: "Cenar",
			estado: false,
		}}

	perrito := seleccionDeTarea(sliceListaDeTarea, "Agregar")
	imprimirTarea(perrito)
}

// seleccionDeTareas ésta funcion permite al usuario seleccionar la tarea que desea ejecutar.
func seleccionDeTarea(perrito []tarea, accionTarea string) []tarea {

	switch accionTarea {
	case "Agregar":
		perrito1 := agregarTarea(perrito, "lavar")
		imprimirTarea(perrito1)
		return perrito1
	case "Completar":
		completarTarea(perrito, 0)
		imprimirTarea(perrito)
	case "Imprimir":
		imprimirTarea(perrito)
	}

	return perrito

}

// imprimirTarea ésta función muestra en consola cada tarea, junto a su estado (completada o no completada)
func imprimirTarea(carro []tarea) {

	for _, tarea := range carro {
		fmt.Println(fmt.Sprintf(" Tarea: %s Estado: %t", tarea.nombre, tarea.estado))

	}
}

func completarTarea(listaDeTarea []tarea, tareaCompletar int) {

	for indice := range listaDeTarea {
		if tareaCompletar == indice {
			listaDeTarea[indice].estado = true
			fmt.Println(" indice: Completada", indice, listaDeTarea[indice].estado)
			return
		}
	}

}

// agregarTarea ésta función solicita al usuario la nueva tarea a añadir a la lista de tareas y la agrega al map tareas.
func agregarTarea(listaDeTarea []tarea, nombreTareaAgregar string) []tarea {

	var tareaNueva tarea
	tareaNueva.nombre = nombreTareaAgregar

	listaDeTarea = append(listaDeTarea, tareaNueva)

	return listaDeTarea
}
