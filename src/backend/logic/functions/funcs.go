package functions

import (
	"math"

	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic/model"
)

type TaskFactory func(eps, t float64) model.Task

const pi = math.Pi

func GetTask(id int) TaskFactory {
	return funcsFactories[id]
}

var funcsFactories map[int]TaskFactory = map[int]TaskFactory{
	1: createProblem1,
}

func constF(x float64) model.RF {
	return func(_ float64) float64 {
		return x
	}
}

func theta1(phi0, phi1 float64) model.RF {
	return func(x float64) float64 {
		return math.Sin(pi*x) + x*phi1 + (1-x)*phi0
	}
}

func sol1(phi0, phi1, eps float64) model.D2RF {
	return func(x, t float64) float64 {
		return math.Sin(pi*x)*math.Exp(-1*math.Pow(pi, 2)*eps*t) + x*phi1 + (1-x)*phi0
	}
}

func createProblem1(eps, t float64) model.Task {
	phi1 := constF(1)(t)
	phi0 := constF(1)(t)
	return model.Task{ID: 1, Name: "Prolbem 1", Phi0: constF(1), Phi1: constF(1), Theta: theta1(phi0, phi1)}
}

func EvaluateOriginal(grid []float64, f model.RF) []float64 {
	var res []float64
	for _, v := range grid {
		res = append(res, f(v))
	}
	return res
}
