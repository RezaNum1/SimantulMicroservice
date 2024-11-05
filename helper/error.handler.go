package helper

import (
	"Expire/data/response"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	Code     int
	Message  string
	FileName string
	AtLine   int
}

func ResponseError(ctx *gin.Context, err CustomError) {
	fmt.Printf("🔴 Error: %d - %s | %s at Line: %d\n", err.Code, err.Message, err.FileName, err.AtLine)
	errResponse := response.Response{
		Success: false,
		Code:    int(err.Code),
		Message: fmt.Sprintf("%s - %s", http.StatusText(err.Code), err.Message),
		Data:    nil,
	}
	ctx.JSON(
		GetErrorCode(err.Code), errResponse)
}

func GetErrorCode(code int) int {
	switch code {
	case 400:
		return http.StatusBadRequest
	case 401:
		return http.StatusUnauthorized
	case 402:
		return http.StatusPaymentRequired
	case 403:
		return http.StatusForbidden
	case 404:
		return http.StatusNotFound
	case 405:
		return http.StatusMethodNotAllowed
	case 500:
		return http.StatusInternalServerError
	case 501:
		return http.StatusNotImplemented
	case 502:
		return http.StatusBadGateway
	case 503:
		return http.StatusServiceUnavailable
	case 504:
		return http.StatusGatewayTimeout
	case 505:
		return http.StatusHTTPVersionNotSupported
	default:
		return http.StatusInternalServerError
	}
}

func GetFileAndLine(err error) (fileName string, atLine int) {
	if err == nil {
		return "", 0
	}

	// Get the program counter (PC) for the error
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "", 0
	}

	// Get the file and line information for the PC
	frame, _ := runtime.CallersFrames([]uintptr{pc}).Next()

	return frame.File, frame.Line
}
