package chartservice

import (
	"fmt"
	"log"
	"pancakaki/internal/domain/entity"
	webchart "pancakaki/internal/domain/web/chart"
	"pancakaki/utils/helper"

	chartrepository "pancakaki/internal/repository/chart"
	productrepository "pancakaki/internal/repository/product"
)

type ChartServiceImpl struct {
	ChartRepository   chartrepository.ChartRepository
	ProductRepository productrepository.ProductRepository
}

func NewChartService(chartRepository chartrepository.ChartRepository, productRepository productrepository.ProductRepository) ChartService {
	return &ChartServiceImpl{
		ChartRepository:   chartRepository,
		ProductRepository: productRepository,
	}
}

func (chartService *ChartServiceImpl) Register(req webchart.ChartCreateRequest) (webchart.ChartResponse, error) {
	log.Println(req, "req")
	fmt.Scanln()
	product, _ := chartService.ProductRepository.FindProductById(req.ProductId)
	totalPrice := product.Price * req.Qty

	chart := entity.Chart{
		Qty:        req.Qty,
		Total:      float64(totalPrice),
		CustomerId: req.CustomerId,
		ProductId:  req.ProductId,
	}

	log.Println(chart, "chart service")
	fmt.Scanln()

	chartData, _ := chartService.ChartRepository.Create(&chart)

	chartResponse := webchart.ChartResponse{
		Id:         chartData.Id,
		Qty:        chartData.Qty,
		Total:      chartData.Total,
		CustomerId: chartData.CustomerId,
		ProductId:  chartData.ProductId,
	}
	return chartResponse, nil
}

func (chartService *ChartServiceImpl) ViewAll(customerId int) ([]webchart.ChartResponse, error) {

	chartData, err := chartService.ChartRepository.FindAll(customerId)
	helper.PanicErr(err)

	chartResponse := make([]webchart.ChartResponse, len(chartData))
	for i, chart := range chartData {
		chartResponse[i] = webchart.ChartResponse{
			Id:         chart.Id,
			Qty:        chart.Qty,
			Total:      chart.Total,
			CustomerId: chart.CustomerId,
			ProductId:  chart.ProductId,
		}
	}
	return chartResponse, nil
}

func (chartService *ChartServiceImpl) ViewOne(chartId int) (webchart.ChartResponse, error) {
	chart, err := chartService.ChartRepository.FindById(chartId)
	helper.PanicErr(err)

	chartResponse := webchart.ChartResponse{
		Id:         chart.Id,
		Qty:        chart.Qty,
		Total:      chart.Total,
		CustomerId: chart.CustomerId,
		ProductId:  chart.ProductId,
	}

	return chartResponse, nil
}

func (chartService *ChartServiceImpl) Edit(req webchart.ChartUpdateRequest) (webchart.ChartResponse, error) {
	chart, _ := chartService.ChartRepository.FindById(req.Id)
	product, _ := chartService.ProductRepository.FindProductById(chart.ProductId)
	totalPrice := product.Price * req.Qty

	chartUpdate := entity.Chart{
		Id:    chart.Id,
		Qty:   req.Qty,
		Total: float64(totalPrice),
	}

	chartData, err := chartService.ChartRepository.Update(&chartUpdate)
	helper.PanicErr(err)

	chartResponse := webchart.ChartResponse{
		Id:         chartData.Id,
		Qty:        chartData.Qty,
		Total:      float64(totalPrice),
		CustomerId: chart.CustomerId,
		ProductId:  chart.ProductId,
	}

	return chartResponse, nil
}

func (chartService *ChartServiceImpl) Unreg(chartId int) (webchart.ChartResponse, error) {

	chartData, err := chartService.ChartRepository.FindById(chartId)
	helper.PanicErr(err)

	err = chartService.ChartRepository.Delete(chartId)
	helper.PanicErr(err)

	chartResponse := webchart.ChartResponse{
		Id:         chartData.Id,
		Qty:        chartData.Qty,
		Total:      chartData.Total,
		CustomerId: chartData.CustomerId,
		ProductId:  chartData.ProductId,
	}

	return chartResponse, nil
}
