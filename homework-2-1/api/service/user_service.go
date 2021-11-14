package service

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"

	"user-catalog/db/sql/userstore"
	"user-catalog/domain/user"
)

// UserService implements srv.Service interface
type UserService struct {
	ur user.UserRepo
}

// NewUserService makes new UserService instance
// Inputs:
//   us - UrlStore implementation
//   ls - LinkingStore implementation
// Output:
//   Returns new ShortenService instance
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		ur: userstore.NewPgUserStore(db),
	}
}

// shortenPostHandler '/api/v1/user' POST request handler
func (s *UserService) userPostHandler(rw http.ResponseWriter, r *http.Request) {
}

// userByNameGetHandler '/api/v1/userByName' GET request handler
func (s *UserService) userByNameGetHandler(rw http.ResponseWriter, r *http.Request) {
}

// userByGroupGetHandler '/api/v1/userByGroup' GET request handler
func (s *UserService) userByGroupGetHandler(rw http.ResponseWriter, r *http.Request) {
}

// Register registers service handlers
// Inputs:
//   router - HTTP mux router
func (s *UserService) Register(router *mux.Router) {
	router.HandleFunc("/api/v1/user", s.userPostHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/v1/userByName", s.userByNameGetHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/userByGroup", s.userByGroupGetHandler).Methods(http.MethodGet, http.MethodOptions)
}
