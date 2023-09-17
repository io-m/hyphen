package user_handler

import (
	"fmt"
	"net/http"

	user_logic "github.com/io-m/hyphen/internal/features/user/logic"
	"github.com/io-m/hyphen/internal/shared/models"
	"github.com/io-m/hyphen/pkg/constants"
	"github.com/io-m/hyphen/pkg/helpers"
)

type IUserAuth interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	OAuth(w http.ResponseWriter, r *http.Request)
}

type userAuthHandler struct {
	userLogic user_logic.IUserLogic
}

func NewUserAuthHandler(userLogic user_logic.IUserLogic) *userAuthHandler {
	return &userAuthHandler{
		userLogic: userLogic,
	}
}

func (uah *userAuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	userRequest, err := helpers.DecodePayload[*models.User](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}
	personResponse, err := uah.userLogic.RegisterUser(r.Context(), userRequest)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not register: %w", err), http.StatusBadRequest)
		return
	}
	helpers.SuccessResponse(w, personResponse, "Person successfully registered", http.StatusCreated)
}

func (uah *userAuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	c, err := helpers.DecodePayload[*models.User](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}

	loggedInUser, accessToken, refreshToken, err := uah.userLogic.Login(r.Context(), c)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while logging in: %w", err), http.StatusNotFound)
		return
	}

	w.Header().Add(constants.ACCESS_TOKEN_HEADER, *accessToken)
	w.Header().Add(constants.REFRESH_TOKEN_HEADER, *refreshToken)

	helpers.SuccessResponse(w, loggedInUser, "Person successfully logged in")
}

func (uah *userAuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

}
func (uah *userAuthHandler) OAuth(w http.ResponseWriter, r *http.Request) {

}
