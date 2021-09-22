package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/GOLANG-NINJA/crud-app/internal/domain"
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

	token, err := h.usersService.SignIn(r.Context(), inp)
	if err != nil {
		logError("signIn", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]string{
		"token": token,
	})
	if err != nil {
		logError("signIn", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
