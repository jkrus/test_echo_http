package handler

import (
	"github.com/jkrus/test_echo_http/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(c echo.Context) error {
	var input model.User

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return err
	}

	id, err := h.services.User.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return err
}

func (h *Handler) getAllUsers(c echo.Context) error {
	items, err := h.services.User.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, items)

	return err
}

func (h *Handler) getById(c echo.Context) error {
	var (
		input model.User
		err   error
	)
	input.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return err
	}

	user, err := h.services.User.GetById(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	c.JSON(http.StatusOK, user)
	return err
}

func (h *Handler) updateUser(c echo.Context) error {
	var (
		input model.User
		err   error
	)

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return err
	}

	input.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return err
	}

	if err := h.services.User.Update(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
	return err
}

func (h *Handler) deleteUser(c echo.Context) error {
	var (
		input model.User
		err   error
	)
	input.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return err
	}

	err = h.services.User.Delete(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
	return err
}
