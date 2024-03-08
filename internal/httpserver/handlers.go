package httpserver

import (
	"net/http"

	"pxls-go/internal/model"

	"github.com/labstack/echo/v4"
)

// RegisterHandlers - Register all handlers
func RegisterHandlers(e *echo.Echo) {
	e.GET("/info", getInfo)
	e.GET("/access", getAccess)
	e.GET("/boards", getBoards)
	e.GET("/board/:name", getBoard)
}

// getInfo - GET /info
func getInfo(c echo.Context) error {
	res := model.InfoResponse{
		Name:    "pxls-go",
		Version: "0.1.0",
		Source:  "",
		Extensions: []string{
			"foo",
			"bar",
			"baz",
		},
	}

	return c.JSON(http.StatusOK, res)
}

// getAccess - GET /access
func getAccess(c echo.Context) error {
	res := model.AccessResponse{
		Permissions: []string{
			"foo",
			"bar",
			"baz",
		},
	}

	return c.JSON(http.StatusOK, res)
}

var fakeBoard = model.Board{
	Name:      "test",
	CreatedAt: 1234567890,
	Shape: [][]int{
		{10, 10},
	},
	Palette: []model.PaletteEntry{
		{
			Name:  "transparent",
			Color: 0x00000000,
		},
		{
			Name:  "black",
			Color: 0x000000FF,
		},
		{
			Name:  "white",
			Color: 0xFFFFFFFF,
		},
	},
	MaxAvailablePixels: 6,
}

func getBoards(c echo.Context) error {
	// TODO: Implement board fetching
	res := []model.Board{
		fakeBoard,
		fakeBoard,
		fakeBoard,
	}

	return c.JSON(http.StatusOK, res)
}

// getBoard - GET /board/:name
func getBoard(c echo.Context) error {
	name := c.Param("name")

	// TODO: Implement board fetching
	res := fakeBoard
	res.Name = name

	return c.JSON(http.StatusOK, res)
}
