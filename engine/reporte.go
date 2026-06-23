package engine

import (
	"fmt"
	"strings"
	"poo_taller15/strategy"
)

func GenerarReporteJustificado(
	ctx *strategy.Contexto) string {

	var sb strings.Builder

	sb.WriteString("===== REPORTE =====\n")

	sb.WriteString(
		fmt.Sprintf("Docente: %s\n",
			ctx.Docente.Nombre))

	sb.WriteString(
		fmt.Sprintf("Puntaje Base: %.2f\n",
			ctx.Evaluacion.PuntajeBase))

	sb.WriteString(
		fmt.Sprintf("Puntaje Final: %.2f\n",
			ctx.Evaluacion.PuntajeFinal))

	sb.WriteString("\nReglas Aplicadas:\n")

	for _, r := range ctx.ReglasAplicadas {
		sb.WriteString("- " + r + "\n")
	}

	return sb.String()
}