package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phucnh/go-app-sample/application/dto"
	"github.com/phucnh/go-app-sample/errs"
	"github.com/phucnh/go-app-sample/pkg/idgenerator"
	"github.com/phucnh/go-app-sample/service"
)

type UserHandler struct {
	userService service.UserService
	idGenerator idgenerator.IDGenerator
}

// NewUserHandler setup rest api handlers for user
func NewUserHandler(echo *echo.Echo,
	userService service.UserService,
	idGenerator idgenerator.IDGenerator,
) {
	handler := &UserHandler{
		userService: userService,
		idGenerator: idGenerator,
	}
	echo.POST("/users", handler.CreateUser)
	echo.POST("/users/login", handler.Login)
}

// CreateUser is method for create user api endpoint
func (h *UserHandler) CreateUser(c echo.Context) error {
	user := dto.CreateUserRequest{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(errs.ErrInvalidRequestBody.Error()))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	user.ID = h.idGenerator.GenerateNewID()
	userEntity, err := h.userService.CreateUser(user.ToEntity())
	if err == errs.ErrEmailExisted {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
	}

	userDto := dto.NewUserResponseFromEntity(userEntity)
	return c.JSON(http.StatusCreated, dto.NewSuccessResponse(userDto))
}

// Login is method to process user login
func (h *UserHandler) Login(c echo.Context) error {
	loginParam := dto.UserLoginRequest{}
	if err := c.Bind(&loginParam); err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(errs.ErrInvalidRequestBody.Error()))
	}

	if err := c.Validate(loginParam); err != nil {
		return err
	}

	token, err := h.userService.Login(loginParam.Email, loginParam.Password)
	if err == errs.ErrInternalServer {
		return c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
	}

	loginResponse := map[string]string{
		"token": token,
	}

	return c.JSON(http.StatusCreated, dto.NewSuccessResponse(loginResponse))
}
