package dto

import "github.com/DewaBiara/INVM-System/pkg/entity"

type CreateSupplierRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Telp    string `json:"telp"`
}

func (u *CreateSupplierRequest) ToEntity() *entity.Supplier {
	return &entity.Supplier{
		Name:    u.Name,
		Address: u.Address,
		Telp:    u.Telp,
	}
}

type UpdateSupplierRequest struct {
	ID      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Telp    string `json:"telp"`
}

func (u *UpdateSupplierRequest) ToEntity() *entity.Supplier {
	return &entity.Supplier{
		Name:    u.Name,
		Address: u.Address,
		Telp:    u.Telp,
	}
}

type GetSingleSupplierResponse struct {
	ID      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Telp    string `json:"telp"`
}

func NewGetSingleSupplierResponse(supplier *entity.Supplier) *GetSingleSupplierResponse {
	return &GetSingleSupplierResponse{
		ID:      supplier.ID,
		Name:    supplier.Name,
		Address: supplier.Address,
		Telp:    supplier.Telp,
	}
}

type GetPageSupplierResponse struct {
	ID      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Telp    string `json:"telp"`
}

func NewGetPageSupplierResponse(supplier *entity.Supplier) *GetPageSupplierResponse {
	return &GetPageSupplierResponse{
		ID:      supplier.ID,
		Name:    supplier.Name,
		Address: supplier.Address,
		Telp:    supplier.Telp,
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
