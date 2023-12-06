package handler

import (
	"GOLANG/entity"
	"GOLANG/internal/http/validator"
	"GOLANG/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserUsecase
}

// melakukan instace dari user handler
func NewUserHandler(userService service.UserUsecase) *UserHandler {
	return &UserHandler{userService}
}

// func untuk melakukan getAll User
func (h *UserHandler) GetAllUser(ctx echo.Context) error {
	users, err := h.userService.GetAll(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
}

// func untuk melakukan createUser
func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var input struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	//ini func untuk error checking
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.NewUser(input.Name, input.Email, input.Password)
	err := h.userService.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	//kalau retrun nya kaya gini akan tampil pesan "User created successfully"
	return ctx.JSON(http.StatusCreated, "User created successfully")

	//tapi kalau bikin retrun nya kaya gini bakal tampil data user yang baru dibuat
	//return ctx.JSON(http.StatusCreated, user)

	// return ctx.JSON(http.StatusOK, map[string]interface{}{
	// 	"message": "User updated successfully",
	// 	"data": map[string]interface{}{
	// 		"id":       user.ID,
	// 		"name":     user.Name,
	// 		"email":    user.Email,
	// 		"password": user.Password,
	// 		"created":  user.CreatedAt,
	// 	},
	// })
}

// func untuk melakukan updateUser by id
func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var input struct {
		ID       int64  `param:"id" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.UpdateUser(input.ID, input.Name, input.Email, input.Password)
	err := h.userService.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    user,
		// "updated": user.UpdatedAt, //buat munculin si updateAt nya
	})
}

// // func untuk melakukan getUser by id
// func (h *UserHandler) GetUserByID(ctx echo.Context) error {
// 	id := ctx.Param("id")
// 	user, err := h.userService.GetUserByID(ctx.Request().Context(), id)
// 	if err != nil {
// 		return ctx.JSON(http.StatusUnprocessableEntity, err)
// 	}

// 	return ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"data": user,
// 	})
// }

func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Jika tidak dapat mengonversi ID menjadi int64, kembalikan respons error
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	user, err := h.userService.GetUserByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":       user.ID,
			"name":     user.Name,
			"email":    user.Email,
			"password": user.Password,
			"created":  user.CreatedAt,
			"updated":  user.UpdatedAt,
		},
	})
}

// func untuk melakukan deleteUser by id
// func (h *UserHandler) DeleteUser(ctx echo.Context) error {
// 	idStr := ctx.Param("id")
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		// Jika tidak dapat mengonversi ID menjadi int64, kembalikan respons error
// 		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"error": "Invalid ID",
// 		})
// 	}

// 	err = h.userService.Delete(ctx.Request().Context(), id)
// 	if err != nil {
// 		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 	}

// 	return ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "User deleted successfully",
// 	})
// }

func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.userService.Delete(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

//NOTE :
// FOLDER INI UNTUK MEMANGGIL SERVICE DAN REPOSITORY
