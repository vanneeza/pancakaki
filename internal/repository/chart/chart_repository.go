package chartrepository

import "pancakaki/internal/domain/entity"

type ChartRepository interface {
	Create(chart *entity.Chart) (*entity.Chart, error)
	FindAll(customerId int) ([]entity.Chart, error)
	FindById(id int) (*entity.Chart, error)
	Update(chart *entity.Chart) (*entity.Chart, error)
	Delete(chartId int) error
}
