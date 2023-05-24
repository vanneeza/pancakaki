package ownerrepository

import (
	"database/sql"
	"fmt"
	entity "pancakaki/internal/domain/entity/owner"
)

type OwnerRepository interface {
	CreateOwner(newOwner *entity.Owner) (*entity.Owner, error)
	GetOwnerByName(name string) (*entity.Owner, error)
	GetOwnerById(id int) (*entity.Owner, error)
	GetOwnerByEmail(email string) (*entity.Owner, error)
	UpdateOwner(updateOwner *entity.Owner) (*entity.Owner, error)
	DeleteOwner(id int) error
}

type ownerRepository struct {
	db *sql.DB
}

func (repo *ownerRepository) CreateOwner(newOwner *entity.Owner) (*entity.Owner, error) {
	stmt, err := repo.db.Prepare("INSERT INTO tbl_owner (name, no_hp, email, password, membership_id) VALUES ($1,$2,$3,$4,$5) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to create owner : %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newOwner.Name, newOwner.NoHp, newOwner.Email, newOwner.Password, newOwner.MembershipId).Scan(&newOwner.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to create owner : %w", err)
	}

	return newOwner, nil
}

func (repo *ownerRepository) GetOwnerByName(name string) (*entity.Owner, error) {
	var owner entity.Owner
	stmt, err := repo.db.Prepare("SELECT id,name,no_hp,email,password,membership_id FROM tbl_owner WHERE name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&owner.Id, &owner.Name, &owner.NoHp, &owner.Email, &owner.Password, &owner.MembershipId)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("owner with name %s not found", name)
	} else if err != nil {
		return nil, err
	}

	return &owner, nil
}

func (repo *ownerRepository) GetOwnerById(id int) (*entity.Owner, error) {
	var owner entity.Owner
	stmt, err := repo.db.Prepare("SELECT id,name,no_hp,email,password,membership_id FROM tbl_owner WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&owner.Id, &owner.Name, &owner.NoHp, &owner.Email, &owner.Password, &owner.MembershipId)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("owner with id %d not found", id)
	} else if err != nil {
		return nil, err
	}

	return &owner, nil
}

func (repo *ownerRepository) GetOwnerByEmail(email string) (*entity.Owner, error) {
	var owner entity.Owner
	stmt, err := repo.db.Prepare("SELECT id,name,no_hp,email,password,membership_id FROM tbl_owner WHERE email = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(email).Scan(&owner.Id, &owner.Name, &owner.NoHp, &owner.Email, &owner.Password, &owner.MembershipId)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("owner with email %s not found", email)
	} else if err != nil {
		return nil, err
	}

	return &owner, nil
}

func (repo *ownerRepository) UpdateOwner(updateOwner *entity.Owner) (*entity.Owner, error) {
	stmt, err := repo.db.Prepare("UPDATE tbl_owner SET name = $1, no_hp=$2, email = $3, password = $4, membership_id = $5 WHERE id = $6")
	if err != nil {
		return nil, fmt.Errorf("failed to update owner : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateOwner.Name, updateOwner.NoHp, updateOwner.Email, updateOwner.Password, updateOwner.MembershipId, updateOwner.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update owner : %w", err)
	}

	return updateOwner, nil
}

func (repo *ownerRepository) DeleteOwner(id int) error {
	stmt, err := repo.db.Prepare("UPDATE tbl_owner SET is_deleted = true WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to delete owner : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to delete owner : %w", err)
	}

	return nil
}

func NewOwnerRepository(db *sql.DB) OwnerRepository {
	return &ownerRepository{db: db}
}
