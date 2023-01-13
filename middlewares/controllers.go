package middlewares

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/eliasacevedo/golang-microservice-template/core"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Parameter[T any] struct {
	Required bool
}

type BodyParameter[T any] struct {
	Parameter[T]
	Binding binding.BindingBody
}

type ControllerExecuteFunc[Query any, Body any, Uri any] func(*gin.Context, *Query, *Body, *Uri) uint

type ControllerConfig[Query any, Body any, Uri any] struct {
	Query   *Parameter[Query]
	Body    *BodyParameter[Body]
	Uri     *Parameter[Uri]
	Execute ControllerExecuteFunc[Query, Body, Uri]
}

func AddController[Query any, Body any, Uri any](
	engine *gin.Engine,
	method string,
	path string,
	config *ControllerConfig[Query, Body, Uri],
) {
	engine.Handle(method, path, func(ctx *gin.Context) {
		if engine == nil {
			panic("must specify engine parameter")
		}

		if config == nil {
			panic("must specify config parameter")
		}

		if config.Execute == nil {
			panic("must specify execute function in config parameter")
		}

		var queryValue Query
		if reflect.TypeOf(queryValue) != nil {
			code, err := GetQuery(ctx, &queryValue)
			if config.Query.Required && err != nil {
				ctx.Set(ERROR_CODE_KEY_CONTEXT, code)
				ctx.JSON(http.StatusBadRequest, nil)
				return
			}
		}

		var bValue Body
		if reflect.TypeOf(bValue) != nil {
			code, err := GetBody(ctx, config.Body.Binding, &bValue)
			if err != nil && config.Body.Required {
				ctx.Set(ERROR_CODE_KEY_CONTEXT, code)
				ctx.JSON(http.StatusBadRequest, nil)
				return
			}
		}

		var uriValue Uri
		if reflect.TypeOf(uriValue) != nil {
			code, err := GetUri(ctx, &uriValue)
			if err != nil && config.Uri.Required {
				ctx.Set(ERROR_CODE_KEY_CONTEXT, code)
				ctx.JSON(http.StatusBadRequest, nil)
				return
			}
		}

		errorCode := config.Execute(ctx, &queryValue, &bValue, &uriValue)

		if errorCode == 0 {
			return
		}

		ctx.Set(ERROR_CODE_KEY_CONTEXT, errorCode)
	})
}

func GetBody[T any](ctx *gin.Context, b binding.BindingBody, p *T) (uint, error) {
	if p == nil || ctx == nil {
		return uint(0), nil
	}

	if b == nil {
		return core.ErrBindingNotSpecified, errors.New("must specify which binding will use")
	}

	err := ctx.ShouldBindBodyWith(p, b)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return core.ErrInvalidBody, err
	}
	return uint(0), nil
}

func GetQuery[T any](ctx *gin.Context, p *T) (uint, error) {
	if p == nil || ctx == nil {
		return uint(0), nil
	}

	err := ctx.ShouldBindQuery(p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return core.ErrInvalidQuery, err
	}
	return uint(0), nil
}

func GetUri[T any](ctx *gin.Context, p *T) (uint, error) {
	if p == nil || ctx == nil {
		return uint(0), nil
	}

	err := ctx.ShouldBindUri(p)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return core.ErrInvalidUri, err
	}
	return uint(0), nil
}
