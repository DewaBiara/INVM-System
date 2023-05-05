package impl

import (
	"context"

	"github.com/DewaBiara/INVM-System/internal/inventory/dto"
	"github.com/DewaBiara/INVM-System/internal/inventory/repository"
	"github.com/DewaBiara/INVM-System/internal/inventory/service"
	"github.com/google/uuid"
)

type (
	SaleServiceImpl struct {
		saleRepository repository.SaleRepository
	}
)

func NewSaleServiceImpl(saleRepository repository.SaleRepository) service.SaleService {
	return &SaleServiceImpl{
		saleRepository: saleRepository,
	}
}

func (u *SaleServiceImpl) CreateSale(ctx context.Context, sale *dto.CreateSaleRequest) error {

	saleEntity := sale.ToEntity()
	saleEntity.ID = uint(uuid.New().ID())

	err := u.saleRepository.CreateSale(ctx, saleEntity)
	if err != nil {
		return err
	}

	return nil
}

func (d *SaleServiceImpl) GetSingleSale(ctx context.Context, saleID string) (*dto.GetSingleSaleResponse, error) {
	sale, err := d.saleRepository.GetSingleSale(ctx, saleID)
	if err != nil {
		return nil, err
	}

	var saleResponse = dto.NewGetSingleSaleResponse(sale)

	return saleResponse, nil
}

func (u *SaleServiceImpl) GetPageSale(ctx context.Context, page int, limit int) (*dto.GetPageSalesResponse, error) {
	offset := (page - 1) * limit

	sales, err := u.saleRepository.GetPageSale(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.NewGetPageSalesResponse(sales), nil
}

func (u *SaleServiceImpl) UpdateSale(ctx context.Context, saleID uint, updateSale *dto.UpdateSaleRequest) error {
	sale := updateSale.ToEntity()
	sale.ID = saleID

	return u.saleRepository.UpdateSale(ctx, sale)
}

func (u *SaleServiceImpl) CreateSaleDetail(ctx context.Context, sale *dto.CreateSaleDetailRequest) error {

	saleDetailEntity := sale.ToEntity()
	saleDetailEntity.ID = uint(uuid.New().ID())

	err := u.saleRepository.CreateSaleDetail(ctx, saleDetailEntity)
	if err != nil {
		return err
	}

	return nil
}

func (d *SaleServiceImpl) GetSingleSaleDetail(ctx context.Context, saleDetailID string) (*dto.GetSingleSaleDetailResponse, error) {
	saleDetail, err := d.saleRepository.GetSingleSaleDetail(ctx, saleDetailID)
	if err != nil {
		return nil, err
	}

	var saleDetailResponse = dto.NewGetSingleSaleDetailResponse(saleDetail)

	return saleDetailResponse, nil
}

func (u *SaleServiceImpl) GetPageSaleDetail(ctx context.Context, page int, limit int) (*dto.GetPageSaleDetailsResponse, error) {
	offset := (page - 1) * limit

	saleDetails, err := u.saleRepository.GetPageSaleDetail(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.NewGetPageSaleDetailsResponse(saleDetails), nil
}

func (u *SaleServiceImpl) UpdateSaleDetail(ctx context.Context, saleDetailID uint, updateSaleDetail *dto.UpdateSaleDetailRequest) error {
	saleDetail := updateSaleDetail.ToEntity()
	saleDetail.ID = saleDetailID

	return u.saleRepository.UpdateSaleDetail(ctx, saleDetail)
}
