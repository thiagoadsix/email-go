package campaign

type Repository interface {
	Create(campaign *Campaign) error
	FindAll() (*[]Campaign, error)
	FindByID(id string) (*Campaign, error)
	Update(campaign *Campaign) error
	Delete(campaign *Campaign) error
}
