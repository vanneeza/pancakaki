package chartrepository

import (
	"database/sql"
	"fmt"
	"log"
	"pancakaki/internal/domain/entity"
)

type ChartRepositoryImpl struct {
	Db *sql.DB
}

func NewChartRepository(Db *sql.DB) ChartRepository {
	return &ChartRepositoryImpl{
		Db: Db,
	}
}

func (r *ChartRepositoryImpl) Create(chart *entity.Chart) (*entity.Chart, error) {
	stmt, err := r.Db.Prepare("INSERT INTO tbl_chart (quantity, total, customer_id, product_id) VALUES ($1, $2, $3, $4) RETURNING id")
	s := fmt.Sprintf("INSERT INTO tbl_chart (quantity, total, customer_id, product_id) VALUES (%d, %f, %d, %d) RETURNING id", chart.Qty, chart.Total, chart.CustomerId, chart.ProductId)
	log.Println(s)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(chart.Qty, chart.Total, chart.CustomerId, chart.ProductId).Scan(&chart.Id)

	if err != nil {
		return nil, err
	}

	return chart, nil
}

func (r *ChartRepositoryImpl) FindAll(customerId int) ([]entity.Chart, error) {
	var tbl_chart []entity.Chart
	rows, err := r.Db.Query(`SELECT id, quantity, total, customer_id, product_id FROM tbl_chart WHERE is_deleted = 'FALSE' AND customer_id = $1`, customerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var chart entity.Chart
		err := rows.Scan(&chart.Id, &chart.Qty, &chart.Total, &chart.CustomerId, &chart.ProductId)
		if err != nil {
			return nil, err
		}
		tbl_chart = append(tbl_chart, chart)
	}

	return tbl_chart, nil
}

func (r *ChartRepositoryImpl) FindById(id int) (*entity.Chart, error) {
	var chart entity.Chart
	stmt, err := r.Db.Prepare("SELECT id, quantity, total, customer_id, product_id FROM tbl_chart WHERE id = $1 AND is_deleted = 'FALSE'")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&chart.Id, &chart.Qty, &chart.Total, &chart.CustomerId, &chart.ProductId)
	if err != nil {
		return nil, err
	}

	return &chart, nil
}

func (r *ChartRepositoryImpl) Update(chart *entity.Chart) (*entity.Chart, error) {
	stmt, err := r.Db.Prepare(`UPDATE tbl_chart	SET quantity = $1, total = $2 WHERE id = $3`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(chart.Qty, chart.Total, chart.Id)
	if err != nil {
		return nil, err
	}

	return chart, nil
}

func (r *ChartRepositoryImpl) Delete(chartId int) error {
	stmt, err := r.Db.Prepare("UPDATE tbl_chart SET is_deleted = TRUE WHERE id= $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(chartId)
	if err != nil {
		return err
	}

	return nil
}
