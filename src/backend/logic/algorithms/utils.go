package algorithms

import "github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic/model"

func GenerateUniformGrid(size int) []float64 {
	return GenerateGrid(size, 1.0)
}

func GenerateGrid(size int, max float64) []float64 {
	h := max / float64(size)
	var x []float64
	for i := 0; i <= size; i++ {
		x = append(x, h*float64(i))
	}
	return x
}

func GetTimeGridStep(params model.SolutionParameters) float64 {
	h := 1.0 / float64(params.NumericalGridSize)
	return params.Kurant * h * h / params.Eps
}
