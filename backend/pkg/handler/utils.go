package handler

import (
	"fmt"
	"net/http"

	echosessions "github.com/flatcar/nebraska/backend/pkg/sessions/echo"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func getTeamID(c echo.Context) string {
	if val, ok := c.Get("team_id").(string); ok {
		return val
	}
	return ""
}

func loggerWithUsername(_ zerolog.Logger, ctx echo.Context) zerolog.Logger {
	session := echosessions.GetSession(ctx)
	if session == nil {
		return l
	}

	username := session.Get("username")

	return l.With().Str("username", username.(string)).Logger()
}

func appNotFoundResponse(ctx echo.Context, appIDProductID string) error {
	return ctx.JSON(http.StatusBadRequest, map[string]string{
		"message": fmt.Sprintf("App not found for :%s", appIDProductID),
	})
}
