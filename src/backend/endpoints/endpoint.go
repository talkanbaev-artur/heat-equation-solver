package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic"
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic/model"
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/util"
)

type Endpoints struct {
	GetNumericals    endpoint.Endpoint
	GetSchemas       endpoint.Endpoint
	GetCacheData     endpoint.Endpoint
	GetFrameSolution endpoint.Endpoint
}

func CreateEndpoints(s logic.APIService, lg log.Logger) Endpoints {
	es := Endpoints{}
	es.GetNumericals = MakeGetNumericalsEndpoint(s, lg)
	es.GetSchemas = makeGetSchemas(s, lg)
	es.GetCacheData = makeGetCacheData(s, lg)
	es.GetFrameSolution = makeGetFrameSolution(s, lg)
	return es
}

func MakeGetNumericalsEndpoint(s logic.APIService, lg log.Logger) endpoint.Endpoint {
	e := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data := s.GetNumericals(ctx)
		return data, nil
	}
	e = util.LoggingMiddleware(lg, "list")(e)
	return e
}

func makeGetCacheData(s logic.APIService, lg log.Logger) endpoint.Endpoint {
	e := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		p := request.(model.SolutionParameters)
		data := s.GetCacheableData(ctx, p)
		return data, nil
	}
	e = util.LoggingMiddleware(lg, "get cacheable")(e)
	return e
}

type GetFrameSolutionInput struct {
	ParamsID  uuid.UUID                `json:"id"`
	Params    model.SolutionParameters `json:"params"`
	TimePoint float64                  `json:"t"`
}

func makeGetFrameSolution(s logic.APIService, lg log.Logger) endpoint.Endpoint {
	e := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		p := request.(GetFrameSolutionInput)
		data := s.GetSolution(ctx, p.Params, p.TimePoint, p.ParamsID)
		return data, nil
	}
	e = util.LoggingMiddleware(lg, "get solution")(e)
	return e
}

func makeGetSchemas(s logic.APIService, lg log.Logger) endpoint.Endpoint {
	e := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data := s.GetSchemas(ctx)
		return data, nil
	}
	e = util.LoggingMiddleware(lg, "get schemas")(e)
	return e
}
