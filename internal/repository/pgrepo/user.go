package pgrepo

import (
	"context"
	"fmt"
	"github.com/ZhanserikKalmukhambet/blog-api/internal/model"
)

func (p *Postgres) CreateUser(ctx context.Context, u *model.User) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                username, -- 1 
			                first_name, -- 2
			                last_name, -- 3
			                hashed_password -- 4
			                )
			VALUES ($1, $2, $3, $4)
			`, usersTable)

	_, err := p.Pool.Exec(ctx, query, u.Username, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return err
	}

	return nil
}
