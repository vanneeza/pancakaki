package adminrepository

import (
	"database/sql"
	entity "pancakaki/internal/domain/entity/admin"
)

type AdminRepositoryImpl struct {
	Db *sql.DB
}

func NewAdminRepository(Db *sql.DB) AdminRepository {
	return &AdminRepositoryImpl{
		Db: Db,
	}
}

func (r *AdminRepositoryImpl) Create(admin *entity.Admin) (*entity.Admin, error) {
	stmt, err := r.Db.Prepare("INSERT INTO tbl_admin (name, password) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(admin.Name, admin.Password).Scan(&admin.Id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *AdminRepositoryImpl) FindAll() ([]entity.Admin, error) {
	var tbl_admin []entity.Admin
	rows, err := r.Db.Query("SELECT id, name, password FROM tbl_admin")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var admin entity.Admin
		err := rows.Scan(&admin.Id, &admin.Name, &admin.Password)
		if err != nil {
			return nil, err
		}
		tbl_admin = append(tbl_admin, admin)
	}

	return tbl_admin, nil
}

func (r *AdminRepositoryImpl) FindById(id int) (*entity.Admin, error) {
	var admin entity.Admin
	stmt, err := r.Db.Prepare("SELECT id, name, password FROM tbl_admin WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&admin.Id, &admin.Name, &admin.Password)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (r *AdminRepositoryImpl) Update(admin *entity.Admin) (*entity.Admin, error) {
	stmt, err := r.Db.Prepare("UPDATE tbl_admin SET name = $1, password = $2 WHERE id = $3")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Name, admin.Password, admin.Id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *AdminRepositoryImpl) Delete(adminId int) error {
	stmt, err := r.Db.Prepare("DELETE FROM tbl_admin WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(adminId)
	if err != nil {
		return err
	}

	return nil
}
