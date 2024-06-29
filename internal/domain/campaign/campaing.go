package campaign

import (
	internalerros "emailn/internal/internal-erros"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

const (
	Pending string = "pending"
	Started string = "started"
	Done    string = "done"
)

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `validate:"required"`
	CreatedOn time.Time `validate:"required"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
		CreatedOn: time.Now(),
	}

	err := internalerros.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err
}
