package model

import "github.com/google/uuid"

type Numericals struct {
	MethodID          int     `json:"id"`
	MethodTitle       string  `json:"title"`
	MethodDescription string  `json:"description"`
	MethodIcon        *string `json:"icon"`
}

type CacheData struct {
	ID         uuid.UUID `json:"id"`
	OriginalX  []float64 `json:"xOriginal"`
	NumericalX []float64 `json:"xNumer"`
	TimePoints []float64 `json:"timePoints"`
}

type SolutionFrame struct {
	Original  []float64 `json:"original"`
	Numerical []float64 `json:"numerical"`
}

type SolutionParameters struct {
	NumericalGridSize int     `json:"n"`
	TimeMax           float64 `json:"t"`
	SchemaID          int     `json:"schema"`
	TaskID            int     `json:"task"`
	Eps               float64 `json:"eps"`
	Kurant            float64 `json:"kurant"`
}

type Schema struct {
	ID    int                     `json:"id"`
	Name  string                  `json:"name"`
	Value func(k float64) float64 `json:"-"`
}

type RF func(x float64) float64

type D2RF func(x, t float64) float64

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phi0     RF
	Phi1     RF
	Theta    RF
	Solution D2RF
}
