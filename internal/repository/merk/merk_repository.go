package merkrepository

import (
	"database/sql"
	"fmt"
	entity "pancakaki/internal/domain/entity/merk"
)

type MerkRepository interface {
	InsertMerk(newMerk *entity.Merk) (*entity.Merk, error)
	UpdateMerk(updateMerk *entity.Merk) (*entity.Merk, error)
	DeleteMerk(deleteMerk *entity.Merk) error
	FindMerkById(id int) (*entity.Merk, error)
	FindMerkByName(name string) (*entity.Merk, error)
	FindAllMerk() ([]entity.Merk, error)
}

type merkRepository struct {
	db *sql.DB
}

// DeleteMerk implements MerkRepository
func (repo *merkRepository) DeleteMerk(deleteMerk *entity.Merk) error {
	stmt, err := repo.db.Prepare("UPDATE tbl_merk SET is_deleted = true WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to delete merk : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(deleteMerk.Id)
	if err != nil {
		return fmt.Errorf("failed to delete merk : %w", err)
	}

	return nil
}

// FindAllMerk implements MerkRepository
func (repo *merkRepository) FindAllMerk() ([]entity.Merk, error) {
	var merks []entity.Merk
	rows, err := repo.db.Query("SELECT id, name FROM tbl_merk")
	if err != nil {
		return nil, fmt.Errorf("failed to get merk : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var merk entity.Merk
		err := rows.Scan(&merk.Id, &merk.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to get merk : %w", err)
		}
		merks = append(merks, merk)
	}

	return merks, nil
}

// FindMerkById implements MerkRepository
func (repo *merkRepository) FindMerkById(id int) (*entity.Merk, error) {
	var merk entity.Merk
	stmt, err := repo.db.Prepare("SELECT id, name FROM tbl_merk WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&merk.Id, &merk.Name)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("merk with id %d not found", id)
	} else if err != nil {
		return nil, err
	}

	return &merk, nil
}

// FindMerkByName implements MerkRepository
func (repo *merkRepository) FindMerkByName(name string) (*entity.Merk, error) {
	var merk entity.Merk
	stmt, err := repo.db.Prepare("SELECT id, name FROM tbl_merk WHERE name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&merk.Id, &merk.Name)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("merk with name %s not found", name)
	} else if err != nil {
		return nil, err
	}

	return &merk, nil
}

// InsertMerk implements MerkRepository
func (repo *merkRepository) InsertMerk(newMerk *entity.Merk) (*entity.Merk, error) {
	stmt, err := repo.db.Prepare("INSERT INTO tbl_merk (name) VALUES ($1) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to insert merk : %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newMerk.Name).Scan(&newMerk.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert merk : %w", err)
	}

	return newMerk, nil
}

// UpdateMerk implements MerkRepository
func (repo *merkRepository) UpdateMerk(updateMerk *entity.Merk) (*entity.Merk, error) {
	stmt, err := repo.db.Prepare("UPDATE tbl_merk SET name = $1 WHERE id = $2")
	if err != nil {
		return nil, fmt.Errorf("failed to update merk : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateMerk.Name, updateMerk.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update merk : %w", err)
	}

	return updateMerk, nil
}

func NewMerkRepository(db *sql.DB) MerkRepository {
	return &merkRepository{db: db}
}
