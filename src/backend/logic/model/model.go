package model

type Numericals struct {
	MethodID          int     `json:"id"`
	MethodTitle       string  `json:"title"`
	MethodDescription string  `json:"description"`
	MethodIcon        *string `json:"icon"`
}

type CacheData struct {
	OriginalX  []float64 `json:"xOriginal"`
	NumericalX []float64 `json:"xNumer"`
}

type SolutionFrame struct {
	Original  []float64 `json:"original"`
	Numerical []float64 `json:"numerical"`
}

type SolutionParameters struct {
	NumericalGridSize int     `json:"n"`
	TimePoints        int     `json:"m"`
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
