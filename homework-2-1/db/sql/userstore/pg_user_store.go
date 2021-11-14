package userstore

import (
	"context"
	"database/sql"

	"user-catalog/db/sql/uow"
	"user-catalog/domain/user"
)

// PgUserStore implements user.UserRepo
type PgUserStore struct {
	wu *uow.UnitOfWork
}

func NewPgUserStore(db *sql.DB) *PgUserStore {
	return &PgUserStore{
		wu: uow.NewUnitOfWork(db),
	}
}

func (s *PgUserStore) columns() string {
	return "id, name, phone, created_at"
}

func (s *PgUserStore) fields(u *UserDTO) []interface{} {
	return []interface{}{&u.ID, &u.Name, &u.Phone, &u.CreatedAt}
}

func (s *PgUserStore) AddUser(ctx context.Context, u *user.User) (*user.User, error) {
	return nil, nil
}

func (s *PgUserStore) FindUserByName(ctx context.Context, userName string) (*user.User, error) {
	return nil, nil
}

func (s *PgUserStore) FindUserByGroup(ctx context.Context, groupName string) ([]*user.User, error) {
	return []*user.User{}, nil
}
