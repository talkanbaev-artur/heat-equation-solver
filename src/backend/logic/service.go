package logic

import (
	"context"
	"math"

	"github.com/google/uuid"
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic/algorithms"
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic/functions"
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic/model"
)

type APIService interface {
	GetNumericals(ctx context.Context) []model.Numericals
	GetSchemas(ctx context.Context) []model.Schema
	GetCacheableData(ctx context.Context, p model.SolutionParameters) model.CacheData
	GetSolution(ctx context.Context, p model.SolutionParameters, t int, id uuid.UUID) model.SolutionFrame
}

func NewAPIService() APIService {
	return &service{dataCache: make(map[uuid.UUID]model.CacheData), processingData: make(map[procID]model.SolutionFrame)}
}

type service struct {
	dataCache      map[uuid.UUID]model.CacheData
	processingData map[procID]model.SolutionFrame
}

type procID struct {
	paramsId  uuid.UUID
	timePoint int
}

func (s *service) GetNumericals(ctx context.Context) []model.Numericals {
	nums := []model.Numericals{
		{MethodID: 1, MethodTitle: "Numerical methods for heat equation solving"},
	}

	return nums
}

func (s *service) GetSchemas(ctx context.Context) []model.Schema {
	schemas := make([]model.Schema, len(functions.Schemas))
	for i, s := range functions.Schemas {
		schemas[i-1] = s
	}
	return schemas
}

func (s *service) GetCacheableData(ctx context.Context, p model.SolutionParameters) model.CacheData {
	data := model.CacheData{ID: uuid.New()}
	data.NumericalX = algorithms.GenerateUniformGrid(p.NumericalGridSize)
	data.OriginalX = algorithms.GenerateUniformGrid(50000)
	tau := algorithms.GetTimeGridStep(p)
	timePoints := int(math.Round(p.TimeMax/tau)) + 1
	data.TimePoints = algorithms.GenerateGrid(timePoints, p.TimeMax)
	s.dataCache[data.ID] = data
	return data
}

func (s *service) GetSolution(ctx context.Context, p model.SolutionParameters, t int, id uuid.UUID) model.SolutionFrame {
	c := s.dataCache[id]
	timePoint := c.TimePoints[t]
	task := functions.GetTask(p.TaskID)(p.Eps, timePoint)
	var res model.SolutionFrame
	res.Original = functions.EvaluateOriginal(c.OriginalX, task.Theta)
	if t == 0 {
		res.Numerical = functions.EvaluateOriginal(c.NumericalX, task.Theta)
		s.processingData[procID{paramsId: id, timePoint: t}] = res
		return res
	}
	prev := s.processingData[procID{paramsId: id, timePoint: t - 1}]
	res.Numerical = algorithms.Solve(p, timePoint, prev.Numerical)
	s.processingData[procID{paramsId: id, timePoint: t}] = res
	return res
}
