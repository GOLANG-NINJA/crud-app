package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/GOLANG-NINJA/crud-app/internal/domain"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logError("signUp", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp domain.SignUpInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		logError("signUp", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := inp.Validate(); err != nil {
		logError("signUp", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.usersService.SignUp(r.Context(), inp)
	if err != nil {
		logError("signUp", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logError("signIn", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp domain.SignInInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		logError("signIn", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := inp.Validate(); err != nil {
		logError("signIn", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := h.usersService.SignIn(r.Context(), inp)
	if err != nil {
		logError("signIn", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]string{
		"token": accessToken,
	})
	if err != nil {
		logError("signIn", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh-token")
	if err != nil {
		logError("refresh", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logrus.Infof("%s", cookie.Value)

	accessToken, refreshToken, err := h.usersService.RefreshTokens(r.Context(), cookie.Value)
	if err != nil {
		logError("signIn", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]string{
		"token": accessToken,
	})
	if err != nil {
		logError("signIn", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Set-Cookie", fmt.Sprintf("refresh-token='%s'; HttpOnly", refreshToken))
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
