package merkrepository

import (
	"database/sql"
	entity "pancakaki/internal/domain/entity/merk"
)

type MerkRepositoryImpl struct {
	Db *sql.DB
}

func NewMerkRepository(Db *sql.DB) MerkRepository {
	return &MerkRepositoryImpl{
		Db: Db,
	}
}

func (r *MerkRepositoryImpl) Create(merk *entity.Merk) (*entity.Merk, error) {
	stmt, err := r.Db.Prepare("INSERT INTO tbl_merk (name) VALUES ($1) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(merk.Name).Scan(&merk.Id)
	if err != nil {
		return nil, err
	}

	return merk, nil
}

func (r *MerkRepositoryImpl) FindAll() ([]entity.Merk, error) {
	var tbl_merk []entity.Merk
	rows, err := r.Db.Query("SELECT id, name FROM tbl_merk WHERE is_deleted FALSE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var merk entity.Merk
		err := rows.Scan(&merk.Id, &merk.Name)
		if err != nil {
			return nil, err
		}
		tbl_merk = append(tbl_merk, merk)
	}

	return tbl_merk, nil
}

func (r *MerkRepositoryImpl) FindById(id int) (*entity.Merk, error) {
	var merk entity.Merk
	stmt, err := r.Db.Prepare("SELECT id, name FROM tbl_merk WHERE id = $1 AND is_deleted = FALSE")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&merk.Id, &merk.Name)
	if err != nil {
		return nil, err
	}

	return &merk, nil
}

func (r *MerkRepositoryImpl) Update(merk *entity.Merk) (*entity.Merk, error) {
	stmt, err := r.Db.Prepare("UPDATE tbl_merk SET name = $1 WHERE id = $2 AND is_deleted = FALSE")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(merk.Name, merk.Id)
	if err != nil {
		return nil, err
	}

	return merk, nil
}

func (r *MerkRepositoryImpl) Delete(merkId int) error {
	stmt, err := r.Db.Prepare("UPDATE tbl_merk SET is_deleted = TRUE WHERE id= $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(merkId)
	if err != nil {
		return err
	}

	return nil
}
