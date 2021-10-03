package presenter

import (
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Code     int      `json:"code"`
		Status   string   `json:"message"`
		Messages string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = "Something not right"
	response.Meta.Code = status
	response.Meta.Messages = err.Error()

	return c.JSON(status, response)
}
