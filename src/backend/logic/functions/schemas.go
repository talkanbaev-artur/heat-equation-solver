package functions

import (
	"math"

	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic/model"
)

var Schemas = map[int]model.Schema{
	1: {ID: 1, Name: "Явная схема", Value: func(k float64) float64 { return 0 }},
	2: {ID: 2, Name: "Схема Кранка-Николсона", Value: crancNikolsonScheme},
	3: {ID: 3, Name: "Схема с опережением", Value: func(k float64) float64 { return 1 }},
	4: {ID: 4, Name: "Схема с минимальной вязкостью", Value: minimalViscosityScheme},
	5: {ID: 5, Name: "Схема, сохраняющая монотонность", Value: preservingScheme},
	6: {ID: 6, Name: "Схема наивысшего порядка сходимости", Value: highestOrderScheme},
}

func GetSchema(id int) model.Schema {
	return Schemas[id]
}

func crancNikolsonScheme(k float64) float64 {
	return 0.5
}

func minimalViscosityScheme(k float64) float64 {
	return math.Max(0.5, 1-(0.5*k))
}

func preservingScheme(k float64) float64 {
	return math.Max(0.5, 1-(0.75*k))
}

func highestOrderScheme(k float64) float64 {
	return 0.5 - (1.0 / (12.0 * k))
}
