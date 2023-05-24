package membershiprepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
)

type MembershipRepositoryImpl struct {
	Db *sql.DB
}

func NewMembershipRepository(Db *sql.DB) MembershipRepository {
	return &MembershipRepositoryImpl{
		Db: Db,
	}
}

func (r *MembershipRepositoryImpl) Create(membership *entity.Membership) (*entity.Membership, error) {
	stmt, err := r.Db.Prepare("INSERT INTO tbl_membership (name, tax, price) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(membership.Name, membership.Tax, membership.Price).Scan(&membership.Id)
	if err != nil {
		return nil, err
	}

	return membership, nil
}

func (r *MembershipRepositoryImpl) FindAll() ([]entity.Membership, error) {
	var tbl_membership []entity.Membership
	rows, err := r.Db.Query(`SELECT id, name, tax, price FROM tbl_membership WHERE is_deleted = 'FALSE'`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var membership entity.Membership
		err := rows.Scan(&membership.Id, &membership.Name, &membership.Tax, &membership.Price)
		if err != nil {
			return nil, err
		}
		tbl_membership = append(tbl_membership, membership)
	}

	return tbl_membership, nil
}

func (r *MembershipRepositoryImpl) FindById(id int) (*entity.Membership, error) {
	var membership entity.Membership
	stmt, err := r.Db.Prepare("SELECT id, name, tax, price FROM tbl_membership WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&membership.Id, &membership.Name, &membership.Tax, &membership.Price)
	if err != nil {
		return nil, err
	}

	return &membership, nil
}

func (r *MembershipRepositoryImpl) Update(membership *entity.Membership) (*entity.Membership, error) {
	stmt, err := r.Db.Prepare(`UPDATE tbl_membership 
	SET name = $1, tax = $2, price = $3	WHERE id = $4`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(membership.Name, membership.Tax, membership.Price, membership.Id)
	if err != nil {
		return nil, err
	}

	return membership, nil
}

func (r *MembershipRepositoryImpl) Delete(membershipId int) error {
	stmt, err := r.Db.Prepare("UPDATE tbl_membership SET is_deleted = TRUE WHERE id= $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(membershipId)
	if err != nil {
		return err
	}

	return nil
}
