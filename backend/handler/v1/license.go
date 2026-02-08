package v1

import (
	"github.com/labstack/echo/v4"

	"github.com/chaitin/panda-wiki/handler"
)

type LicenseHandler struct {
	*handler.BaseHandler
}

func NewLicenseHandler(e *echo.Echo, baseHandler *handler.BaseHandler) *LicenseHandler {
	h := &LicenseHandler{
		BaseHandler: baseHandler,
	}
	group := e.Group("/api/v1")
	group.GET("/license", h.GetLicense)
	return h
}

// GetLicense returns license information for community edition
//
//	@Summary		GetLicense
//	@Description	Get license information
//	@Tags			license
//	@Produce		json
//	@Success		200	{object}	domain.Response{data=map[string]interface{}}
//	@Router			/api/v1/license [get]
func (h *LicenseHandler) GetLicense(c echo.Context) error {
	return h.NewResponseWithData(c, map[string]interface{}{
		"edition":     "community",
		"max_kb":      999,
		"max_user":    999,
		"expired":     false,
		"expire_time": nil,
	})
}
