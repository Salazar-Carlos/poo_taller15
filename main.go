package main

import (
	"fmt"

	"poo_taller15/engine"
	"poo_taller15/model"
	"poo_taller15/strategy"
)

func ejecutarEscenario(
	docente model.Docente,
	competencia model.Competencia,
	antiguedad int) {

	ctx := strategy.Contexto{
		Docente: docente,
		Competencia: competencia,
		Evaluacion: model.Evaluacion{
			PuntajeBase: 100,
		},
		Antiguedad: antiguedad,
	}

	reglas := []strategy.EstrategiaPonderacion{
		strategy.ReglaR1{},
		strategy.ReglaR2{},
		strategy.ReglaR3{},
		strategy.ReglaR4{},
		strategy.ReglaR5{},
	}

	engine.CalcularPuntajeConReglas(
		&ctx,
		reglas)

	fmt.Println(
		engine.GenerarReporteJustificado(
			&ctx))
}

func main() {

	fmt.Println("ESCENARIO 1")
	ejecutarEscenario(
		model.Docente{
			Nombre: "Juan",
			Departamento: "Ciencias",
			Cargo: "Profesor",
		},
		model.Competencia{
			Tipo: "investigacion",
		},
		5)

	fmt.Println("ESCENARIO 2")
	ejecutarEscenario(
		model.Docente{
			Nombre: "Ana",
			Cargo: "Director",
		},
		model.Competencia{
			Tipo: "gestion",
		},
		5)

	fmt.Println("ESCENARIO 3")
	ejecutarEscenario(
		model.Docente{
			Nombre: "Pedro",
			Departamento: "Ciencias",
			Cargo: "Director",
		},
		model.Competencia{
			Tipo: "gestion",
		},
		5)

	fmt.Println("ESCENARIO 4")
	ejecutarEscenario(
		model.Docente{
			Nombre: "Maria",
			Departamento: "Ciencias",
			Cargo: "Profesor",
		},
		model.Competencia{
			Tipo: "investigacion",
		},
		15)
}