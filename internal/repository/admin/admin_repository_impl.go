package adminrepository

import (
	"database/sql"
	"fmt"
	"pancakaki/internal/domain/entity"
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
	stmt, err := r.Db.Prepare("INSERT INTO tbl_admin (username, password) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(admin.Username, admin.Password).Scan(&admin.Id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *AdminRepositoryImpl) FindAll() ([]entity.Admin, error) {
	var admins []entity.Admin
	rows, err := r.Db.Query("SELECT id, username, password, role FROM tbl_admin WHERE is_deleted = FALSE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var admin entity.Admin
		err := rows.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Role)
		if err == sql.ErrNoRows {
			return admins, fmt.Errorf("admin data not found")
		} else if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins, nil
}

func (r *AdminRepositoryImpl) FindById(id int, username string) (*entity.Admin, error) {
	var admin entity.Admin
	stmt, err := r.Db.Prepare("SELECT id, username, password, role FROM tbl_admin WHERE is_deleted = 'FALSE' AND id = $1 OR username = $2")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id, username)
	err = row.Scan(&admin.Id, &admin.Username, &admin.Password, &admin.Role)
	if err == sql.ErrNoRows {
		return &admin, fmt.Errorf("admin with id %d not found", id)
	} else if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (r *AdminRepositoryImpl) Update(admin *entity.Admin) (*entity.Admin, error) {
	stmt, err := r.Db.Prepare("UPDATE tbl_admin SET username = $1, password = $2 WHERE id = $3")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(admin.Username, admin.Password, admin.Id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *AdminRepositoryImpl) Delete(adminId int) error {
	stmt, err := r.Db.Prepare("Update tbl_admin SET is_deleted = TRUE WHERE id = $1")
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
