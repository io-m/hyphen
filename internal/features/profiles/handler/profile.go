package profile_handler

import (
	"fmt"
	"net/http"

	profile_logic "github.com/io-m/hyphen/internal/features/profiles/logic"
	"github.com/io-m/hyphen/internal/shared/models"
	"github.com/io-m/hyphen/pkg/constants"
	"github.com/io-m/hyphen/pkg/helpers"
)

type IProfileAuth interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	OAuth(w http.ResponseWriter, r *http.Request)
}

type profileAuthHandler struct {
	profileLogic profile_logic.IProfileLogic
}

func NewProfileAuthHandler(profileLogic profile_logic.IProfileLogic) *profileAuthHandler {
	return &profileAuthHandler{
		profileLogic: profileLogic,
	}
}

func (uah *profileAuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	profileRequest, err := helpers.DecodePayload[*models.Profile](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}
	profileResponse, err := uah.profileLogic.RegisterProfile(r.Context(), profileRequest)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not register: %w", err), http.StatusBadRequest)
		return
	}
	helpers.SuccessResponse(w, profileResponse, "profile successfully registered", http.StatusCreated)
}

func (uah *profileAuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	profile, err := helpers.DecodePayload[*models.Profile](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}

	loggedInProfile, accessToken, refreshToken, err := uah.profileLogic.Login(r.Context(), profile)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while logging in: %w", err), http.StatusNotFound)
		return
	}

	w.Header().Add(constants.ACCESS_TOKEN_HEADER, *accessToken)
	w.Header().Add(constants.REFRESH_TOKEN_HEADER, *refreshToken)

	helpers.SuccessResponse(w, loggedInProfile, "profile successfully logged in")
}

func (uah *profileAuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

}
func (uah *profileAuthHandler) OAuth(w http.ResponseWriter, r *http.Request) {

}
