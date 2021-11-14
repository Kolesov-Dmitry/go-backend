package groupstore

import (
	"context"
	"database/sql"
	"fmt"
	"user-catalog/db/sql/uow"
	"user-catalog/domain/group"

	"github.com/google/uuid"
)

// PgGroupStore implements group.GroupRepo
type PgGroupStore struct {
	wu *uow.UnitOfWork
}

func NewPgGroupStore(db *sql.DB) *PgGroupStore {
	return &PgGroupStore{
		wu: uow.NewUnitOfWork(db),
	}
}

func (s *PgGroupStore) columns() string {
	return "id, name, created_at"
}

func (s *PgGroupStore) fields(g *GroupDTO) []interface{} {
	return []interface{}{&g.ID, &g.Name, &g.CreatedAt}
}

func (s *PgGroupStore) AddGroup(ctx context.Context, g *group.Group) (*group.Group, error) {
	gr := &group.Group{
		ID:   uuid.New(),
		Name: g.Name,
	}

	err := s.wu.WithTx(ctx, func(wu uow.UnitOfWork) error {
		_, err := wu.Tx().ExecContext(ctx,
			fmt.Sprintf("INSERT INTO groups(id, name) VALUES($1, $2) RETURNING %s", s.columns()), gr.ID, gr.Name,
		)
		return err
	})

	if err != nil {
		return nil, err
	}

	return gr, nil
}

func (s *PgGroupStore) AppendUserToGroups(ctx context.Context, userName string, groups []string) error {
	return nil
}

func (s *PgGroupStore) RemoveUserFromGroups(ctx context.Context, userName string, groups []string) error {
	return nil
}

func (s *PgGroupStore) FindGroupByName(ctx context.Context, groupName string) (*group.Group, error) {
	return nil, nil
}

func (s *PgGroupStore) FindGroupByUsers(ctx context.Context, users []string) ([]*group.Group, error) {
	return []*group.Group{}, nil
}
