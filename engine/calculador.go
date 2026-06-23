package engine

import (
	"poo_taller15/strategy"
	"sort"
)

func CalcularPuntajeConReglas(ctx *strategy.Contexto,
	estrategias []strategy.EstrategiaPonderacion) (float64, []string) {

	ctx.Evaluacion.PuntajeFinal = ctx.Evaluacion.PuntajeBase

	sort.Slice(estrategias, func(i, j int) bool {
		return estrategias[i].Prioridad() <
			estrategias[j].Prioridad()
	})

	for _, e := range estrategias {

		if e.Aplica(ctx) {

			ctx.Evaluacion.PuntajeFinal =
				e.Calcular(ctx)

			ctx.ReglasAplicadas =
				append(ctx.ReglasAplicadas,
					e.Nombre())
		}
	}

	return ctx.Evaluacion.PuntajeFinal,
		ctx.ReglasAplicadas
}