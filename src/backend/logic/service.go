package logic

import (
	"context"

	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic/functions"
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic/model"
)

type APIService interface {
	GetNumericals(ctx context.Context) []model.Numericals
	GetSchemas(ctx context.Context) []model.Schema
	GetCacheableData(ctx context.Context, p model.SolutionParameters) model.CacheData
	GetSolution(ctx context.Context, p model.SolutionParameters, timePoint float64) model.SolutionFrame
}

func NewAPIService() APIService {
	return service{}
}

type service struct {
}

func (s service) GetNumericals(ctx context.Context) []model.Numericals {
	nums := []model.Numericals{
		{MethodID: 1, MethodTitle: "Numerical methods for 2nd-order ODE Boundary problems"},
	}

	return nums
}

func (s service) GetSchemas(ctx context.Context) []model.Schema {
	schemas := []model.Schema{}
	for _, s := range functions.Schemas {
		schemas = append(schemas, s)
	}
	return schemas
}

func (s service) GetCacheableData(ctx context.Context, p model.SolutionParameters) model.CacheData {
	return model.CacheData{}
}

func (s service) GetSolution(ctx context.Context, p model.SolutionParameters, timePoint float64) model.SolutionFrame {
	return model.SolutionFrame{}
}
