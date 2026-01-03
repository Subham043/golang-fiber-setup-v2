package request_id

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/fx"
)

type RequestIDMiddleware struct {
}

type IRequestIDMiddleware interface {
	RequestIDMiddleware() fiber.Handler
}

func NewRequestIDMiddleware() *RequestIDMiddleware {
	return &RequestIDMiddleware{}
}

func (r *RequestIDMiddleware) RequestIDMiddleware() fiber.Handler {
	return requestid.New()
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewRequestIDMiddleware),
	)
}
