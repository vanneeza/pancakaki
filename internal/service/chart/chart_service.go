package chartservice

import webchart "pancakaki/internal/domain/web/chart"

type ChartService interface {
	Register(req webchart.ChartCreateRequest) (webchart.ChartResponse, error)
	ViewAll(customerId int) ([]webchart.ChartResponse, error)
	ViewOne(memberwebchartId int) (webchart.ChartResponse, error)
	Edit(req webchart.ChartUpdateRequest) (webchart.ChartResponse, error)
	Unreg(memberwebchartId int) (webchart.ChartResponse, error)
}
