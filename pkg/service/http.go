package service

import (
	"fmt"

	"github.com/fagongzi/util/format"
	"github.com/labstack/echo"
)

const (
	apiVersion = "/v1"
)

// InitHTTPRouter init http router
func InitHTTPRouter(server *echo.Echo, ui, uiPrefix string) {
	serverV1 := server.Group(apiVersion)
	initClusterRouter(serverV1)
	initServerRouter(serverV1)
	initBindRouter(serverV1)
	initRoutingRouter(serverV1)
	initAPIRouter(serverV1)
	initSystemRouter(serverV1)
	initStatic(serverV1, ui, uiPrefix)
}

type limitQuery struct {
	limit   int64
	afterID uint64
}

func idParamFactory(ctx echo.Context) (interface{}, error) {
	value := ctx.Param("id")
	if value == "" {
		return nil, fmt.Errorf("missing id path value")
	}

	id, err := format.ParseStrUInt64(value)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func limitQueryFactory(ctx echo.Context) (interface{}, error) {
	query := &limitQuery{
		limit: limit,
	}

	value := ctx.QueryParam("limit")
	if value != "" {
		l, err := format.ParseStrInt64(value)
		if err != nil {
			return nil, err
		}
		query.limit = l
	}

	value = ctx.QueryParam("after")
	if value != "" {
		l, err := format.ParseStrUInt64(value)
		if err != nil {
			return nil, err
		}
		query.afterID = l
	}

	return query, nil
}

func emptyParamFactory(ctx echo.Context) (interface{}, error) {
	return nil, nil
}
