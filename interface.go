package service

// CustomerUsecase 
type CustomerUsecase interface {
	Fetch(context.Context) ([]Customer, error)
	GetByID(context.Context, int64) (Customer, error)
	Update(context.Context, int64, map[string]interface{}) error
	Store(context.Context, *Customer) error
	Delete(context.Context, int64) error
}
