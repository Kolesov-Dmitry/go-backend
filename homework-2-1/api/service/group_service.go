package service

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"

	"user-catalog/db/sql/groupstore"
	"user-catalog/domain/group"
)

// UserService implements srv.Service interface
type GroupService struct {
	ur group.GroupRepo
}

// NewUserService makes new UserService instance
// Inputs:
//   us - UrlStore implementation
//   ls - LinkingStore implementation
// Output:
//   Returns new ShortenService instance
func NewGroupService(db *sql.DB) *GroupService {
	return &GroupService{
		ur: groupstore.NewPgGroupStore(db),
	}
}

// shortenPostHandler '/api/v1/group' POST request handler
func (s *GroupService) groupPostHandler(rw http.ResponseWriter, r *http.Request) {
}

// userByNameGetHandler '/api/v1/groupByName' GET request handler
func (s *GroupService) groupByNameGetHandler(rw http.ResponseWriter, r *http.Request) {
}

// userByGroupGetHandler '/api/v1/groupByUsers' GET request handler
func (s *GroupService) groupByUsersGetHandler(rw http.ResponseWriter, r *http.Request) {
}

// memberPostHandler '/api/v1/member' POST request handler
func (s *GroupService) memberPostHandler(rw http.ResponseWriter, r *http.Request) {
}

// memberPostHandler '/api/v1/member' DELETE request handler
func (s *GroupService) memberDeleteHandler(rw http.ResponseWriter, r *http.Request) {
}

// Register registers service handlers
// Inputs:
//   router - HTTP mux router
func (s *GroupService) Register(router *mux.Router) {
	router.HandleFunc("/api/v1/group", s.groupPostHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/v1/groupByName", s.groupByNameGetHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/groupByUsers", s.groupByUsersGetHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/member", s.memberPostHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/v1/member", s.memberDeleteHandler).Methods(http.MethodDelete)
}
