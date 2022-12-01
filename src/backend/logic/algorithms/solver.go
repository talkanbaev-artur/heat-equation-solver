package algorithms

import (
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic/functions"
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic/model"
)

func Solve(p model.SolutionParameters, timePoint float64, prevVals []float64) []float64 {
	n := len(prevVals)
	var a, b, c, f = make([]float64, n), make([]float64, n), make([]float64, n), make([]float64, n)
	schema, task := functions.GetSchema(p.SchemaID), functions.GetTask(p.TaskID)(p.Eps, timePoint)
	b[0], a[0], c[0], f[0] = 1, 0, 0, task.Phi0(timePoint)
	b[n-1], a[n-1], c[n-1], f[n-1] = 1, 0, 0, task.Phi1(timePoint)

	A := schema.Value(p.Kurant) * p.Kurant
	B := 1 + A + A
	A0 := (1 - schema.Value(p.Kurant)) * p.Kurant
	B0 := 1 - A0 - A0

	for i := 1; i < n-1; i++ {
		a[i], b[i], c[i] = A, B, A
		f[i] = A0*prevVals[i-1] + B0*prevVals[i] + A0*prevVals[i+1]
	}
	solution, _ := solveTridiagonal(a, b, c, f)
	return solution
}
