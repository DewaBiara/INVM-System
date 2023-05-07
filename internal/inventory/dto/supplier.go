package dto

import "github.com/DewaBiara/INVM-System/pkg/entity"

type CreateSupplierRequest struct {
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address"`
	Telp      string `json:"telp"`
	CreatedBy string `json:"createdby"`
}

func (u *CreateSupplierRequest) ToEntity() *entity.Supplier {
	return &entity.Supplier{
		Name:      u.Name,
		Address:   u.Address,
		Telp:      u.Telp,
		CreatedBy: u.CreatedBy,
	}
}

type UpdateSupplierRequest struct {
	ID        uint   `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address"`
	Telp      string `json:"telp"`
	UpdatedBy string `json:"updatedby"`
}

func (u *UpdateSupplierRequest) ToEntity() *entity.Supplier {
	return &entity.Supplier{
		Name:      u.Name,
		Address:   u.Address,
		Telp:      u.Telp,
		UpdatedBy: u.UpdatedBy,
	}
}

type GetSingleSupplierResponse struct {
	ID        uint   `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address"`
	Telp      string `json:"telp"`
	CreatedBy string `json:"createdby"`
	UpdatedBy string `json:"updatedby"`
}

func NewGetSingleSupplierResponse(supplier *entity.Supplier) *GetSingleSupplierResponse {
	return &GetSingleSupplierResponse{
		ID:        supplier.ID,
		Name:      supplier.Name,
		Address:   supplier.Address,
		Telp:      supplier.Telp,
		CreatedBy: supplier.CreatedBy,
		UpdatedBy: supplier.UpdatedBy,
	}
}

type GetPageSupplierResponse struct {
	ID        uint   `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address"`
	Telp      string `json:"telp"`
	CreatedBy string `json:"createdby"`
	UpdatedBy string `json:"updatedby"`
}

func NewGetPageSupplierResponse(supplier *entity.Supplier) *GetPageSupplierResponse {
	return &GetPageSupplierResponse{
		ID:        supplier.ID,
		Name:      supplier.Name,
		Address:   supplier.Address,
		Telp:      supplier.Telp,
		CreatedBy: supplier.CreatedBy,
		UpdatedBy: supplier.UpdatedBy,
	}
}

type GetPageSuppliersResponse []GetPageSupplierResponse

func NewGetPageSuppliersResponse(suppliers *entity.Suppliers) *GetPageSuppliersResponse {
	var getPageSuppliers GetPageSuppliersResponse
	for _, suppliers := range *suppliers {
		getPageSuppliers = append(getPageSuppliers, *NewGetPageSupplierResponse(&suppliers))
	}
	return &getPageSuppliers
}
