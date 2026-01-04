package error_handler

import (
	"fmt"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/subham043/golang-fiber-setup/utils/response"
)

// A struct to handle error with custom error handler.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

// Implement the error interface
func (e *Error) Error() string {
	return e.Message
}

func ServerErrorHandler(c *fiber.Ctx, err error) error {
	path := c.Path()
	errResponse := response.Response{
		Code:    fiber.StatusInternalServerError,
		Path:    path,
		Message: err.Error(),
	}

	switch e := err.(type) {
	//check fiber error
	case *fiber.Error:
		errResponse.Code = e.Code
		errResponse.Message = e.Message

	//check custom error
	case *Error:
		errResponse.Code = e.Code
		errResponse.Message = e.Message

	//check validation error
	case validation.Errors:
		errResponse.Code = fiber.StatusBadRequest
		errResponse.Message = fiber.ErrBadRequest.Error()
		errResponse.Errors = e.Filter()
	}

	// errFields := []zap.Field{
	// 	zap.Int("status", errResponse.Code),
	// 	zap.String("method", c.Method()),
	// 	zap.String("path", errResponse.Path),
	// 	zap.String("ip", c.IP()),
	// 	zap.String("user_agent", string(c.Request().Header.UserAgent())),
	// 	zap.String("message", errResponse.Message),
	// }
	// //if Errors for validation then append it to errFields
	// if errResponse.Errors != nil {
	// 	errFields = append(errFields, zap.Any("errors", errResponse.Errors))
	// }
	// log.Error("error", errFields...)

	return response.Json(c, errResponse)
}

func LimititerErrorHandler(c *fiber.Ctx) error {
	// Get the Retry-After value from response headers
	retryAfterStr := c.GetRespHeader(fiber.HeaderRetryAfter, "0")

	// Convert the string to an integer (seconds)
	retryAfterSeconds, _ := strconv.Atoi(retryAfterStr)

	return response.Json(c, response.Response{
		Code:    fiber.StatusTooManyRequests,
		Path:    c.Path(),
		Message: fmt.Sprintf("Too many requests, try again in %v seconds", retryAfterSeconds),
	})
}

func JWTErrorHandler(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	return response.Json(c, response.Response{
		Code:    fiber.StatusUnauthorized,
		Path:    c.Path(),
		Message: fiber.ErrUnauthorized.Error(),
	})
}
