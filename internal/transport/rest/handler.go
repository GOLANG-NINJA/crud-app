package rest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/GOLANG-NINJA/crud-app/internal/domain"

	"github.com/gorilla/mux"
)

type Books interface {
	Create(ctx context.Context, book domain.Book) error
	GetByID(ctx context.Context, id int64) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int64) error
}

type Handler struct {
	booksService Books
}

func NewHandler(books Books) *Handler {
	return &Handler{
		booksService: books,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	books := r.PathPrefix("/books").Subrouter()
	{
		books.HandleFunc("/{id:[0-9]+}", h.getBookByID).Methods(http.MethodGet)
		books.HandleFunc("", h.createBook).Methods(http.MethodPost)
	}

	return r
}

func (h *Handler) getBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	book, err := h.booksService.GetByID(context.TODO(), id)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(book)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var book domain.Book
	if err = json.Unmarshal(reqBytes, &book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.booksService.Create(context.TODO(), book)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
