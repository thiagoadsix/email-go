package campaign

type Repository interface {
	Save(campaign *Campaign) error
	FindAll() ([]Campaign, error)
	FindByID(id string) (struct {
		ID      string
		Name    string
		Content string
		Status  string
	}, error)
}
