package middlewares

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/eliasacevedo/golang-microservice-template/core"
	"github.com/eliasacevedo/golang-microservice-template/utilities"
	"github.com/gin-gonic/gin"
)

const ERROR_CODE_KEY_CONTEXT = "ErrorCode"

func DataReturnMiddleware(l *utilities.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		bw := &BodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bw

		c.Next()

		errorCode, err := getErrorCode(c)
		if err != nil {
			l.PanicApp("error code is not a number")
		}

		if len(c.Errors) > 0 || errorCode > 0 {
			writeErrors(c, l, bw, errorCode)
			return
		}

		if bw.body.Len() <= 0 {
			return
		}

		if strings.Contains(c.Writer.Header().Get("Content-Type"), "application/json") {
			writeJson(c, l, bw)
			return
		}

		bw.ResponseWriter.Write(bw.body.Bytes())
	}
}

func writeErrors(c *gin.Context, l *utilities.Logger, bw *BodyWriter, errorCode uint) {
	response := core.BaseResponse{
		Data:      nil,
		ErrorCode: errorCode,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		l.PanicApp("invalid kind/format json/data")
	}

	bw.ResponseWriter.Write(jsonResponse)
}

func writeJson(c *gin.Context, l *utilities.Logger, bw *BodyWriter) {
	var data interface{}
	err := json.Unmarshal(bw.body.Bytes(), &data)
	if err != nil {
		l.PanicApp(fmt.Errorf("couldn't convert result to json: %w", err).Error())
	}
	response := core.BaseResponse{
		Data:      data,
		ErrorCode: uint(0),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		l.PanicApp("invalid kind/format json/data")
	}

	bw.ResponseWriter.Write(jsonResponse)
	bw.body.Reset()
}

func getErrorCode(c *gin.Context) (uint, error) {
	v, ok := c.Get(ERROR_CODE_KEY_CONTEXT)
	if !ok {
		return uint(0), nil
	}

	ec, ok := v.(uint)
	if !ok {
		return uint(0), errors.New("error code context value is not a uint")
	}
	return ec, nil
}

type BodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}
