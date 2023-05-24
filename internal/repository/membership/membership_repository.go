package membershiprepository

import (
	"database/sql"
	"fmt"
	entity "pancakaki/internal/domain/entity/membership"
)

type MembershipRepository interface {
	GetMembershipById(id int) (*entity.Membership, error)
}

type membershipRepository struct {
	db *sql.DB
}

func (repo *membershipRepository) GetMembershipById(id int) (*entity.Membership, error) {
	var membership entity.Membership
	stmt, err := repo.db.Prepare("SELECT id,name,tax,price FROM tbl_membership WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&membership.Id, &membership.Name, &membership.Tax, &membership.Price)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("membership with id %d not found", id)
	} else if err != nil {
		return nil, err
	}

	return &membership, nil
}

func NewMembershipRepository(db *sql.DB) MembershipRepository {
	return &membershipRepository{db: db}
}
