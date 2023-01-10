package module

type PaginationParams struct {
	page     uint
	quantity uint
}

func NewPaginationParam(page uint, quantity uint) PaginationParams {
	if page == 0 {
		page = 1
	}

	if quantity == 0 {
		quantity = 1
	}

	return PaginationParams{
		page:     page,
		quantity: quantity,
	}
}
