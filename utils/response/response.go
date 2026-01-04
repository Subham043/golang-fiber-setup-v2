package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Code    int    `json:"code"`
	Path    string `json:"path,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

// A fuction to return beautiful responses.
func Json(c *fiber.Ctx, resp Response) error {
	// Set status
	if resp.Code == 0 {
		resp.Code = fiber.StatusOK
	}
	c.Status(resp.Code)

	// Return JSON
	return c.JSON(resp)
}
