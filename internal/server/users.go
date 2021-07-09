package server

import (
	"github.com/labstack/echo"
	"net/http"
	"sighupio/permission-manager/internal/resources"
)

func listUsers(c echo.Context) error {
	ac := c.(*AppContext)

	users, err := ac.ResourceService.UserList()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
	ac := c.(*AppContext)

	type request struct {
		Name string `json:"name" validate:"required"`
	}

	type response = resources.User

	r := new(request)

	err := ac.validateAndBindRequest(r)

	if err != nil {
		return err
	}

	if !isValidUsername(r.Name) {
		return ac.errorResponse(invalidUsernameError)
	}

	u, err := ac.ResourceService.UserCreate(r.Name)

	if err != nil {
		return err
	}

	return ac.okResponseWithData(response{Name: u.Name})
}

func deleteUser(c echo.Context) error {
	ac := c.(*AppContext)

	type Request struct {
		Username string `json:"username" validate:"required"`
	}

	r := new(Request)

	err := ac.validateAndBindRequest(r)

	if err != nil {
		return err
	}

	err = ac.ResourceService.UserDelete(r.Username)

	if err != nil {
		return err
	}

	return ac.okResponse()
}
