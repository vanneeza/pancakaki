package ownerrepository

import (
	"database/sql"
	"fmt"
	"pancakaki/internal/domain/entity"
)

type OwnerRepository interface {
	CreateOwner(newOwner *entity.Owner) (*entity.Owner, error)
	GetOwnerByNoHp(noHp string) (*entity.Owner, error)
	GetOwnerById(id int) (*entity.Owner, error)
	GetOwnerByEmail(email string) (*entity.Owner, error)
	GetTaxAndStoreOwner(productId int) (string, float64, error)
	UpdateOwner(updateOwner *entity.Owner) (*entity.Owner, error)
	DeleteOwner(id int) error
	// UpdateMembershipOwner(id int)
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

func (repo *ownerRepository) GetOwnerByNoHp(noHp string) (*entity.Owner, error) {
	var owner entity.Owner
	stmt, err := repo.db.Prepare("SELECT id,name,no_hp,email,password,membership_id, role FROM tbl_owner WHERE no_hp = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(noHp).Scan(&owner.Id, &owner.Name, &owner.NoHp, &owner.Email, &owner.Password, &owner.MembershipId, &owner.Role)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("owner with no hp %s not found", noHp)
	} else if err != nil {
		return nil, err
	}

	return &owner, nil
}

func (repo *ownerRepository) GetOwnerById(id int) (*entity.Owner, error) {
	var owner entity.Owner
	stmt, err := repo.db.Prepare("SELECT id,name,no_hp,email,password,membership_id,role FROM tbl_owner WHERE id = $1 AND is_deleted = false")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&owner.Id, &owner.Name, &owner.NoHp, &owner.Email, &owner.Password, &owner.MembershipId, &owner.Role)
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
func (repo *ownerRepository) GetTaxAndStoreOwner(productId int) (string, float64, error) {
	var storeName string
	var tax float64
	stmt, err := repo.db.Prepare(`SELECT tbl_store.Name, tbl_membership.Tax FROM tbl_membership
	INNER JOIN tbl_owner on tbl_membership.id = tbl_owner.membership_id
	INNER JOIN tbl_store ON tbl_owner.id = tbl_store.owner_id
	INNER JOIN tbl_product ON tbl_store.id = tbl_product.store_id WHERE tbl_product.id = $1`)
	if err != nil {
		return storeName, tax, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(productId).Scan(&storeName, &tax)
	if err == sql.ErrNoRows {
		return "", 0, fmt.Errorf("product with id %d not found", productId)
	} else if err != nil {
		return "", 0, err
	}

	return storeName, tax, nil
}

func NewOwnerRepository(db *sql.DB) OwnerRepository {
	return &ownerRepository{db: db}
}
