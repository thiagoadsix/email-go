package campaign

import (
	internalerros "emailn/internal/internal-erros"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50"`
}

const (
	Pending   string = "pending"
	Started   string = "started"
	Done      string = "done"
	Cancelled string = "cancelled"
	Fail      string = "fail"
)

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50;not null"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100;not null"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024;not null"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `validate:"required" gorm:"size:20;not null"`
	CreatedOn time.Time `validate:"required" gorm:"not null"`
	UpdatedOn time.Time
	CreatedBy string `validate:"email" gorm:"size:100;not null"`
}

func NewCampaign(name string, content string, emails []string, createdBy string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
		contacts[index].ID = xid.New().String()
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
		CreatedOn: time.Now(),
		CreatedBy: createdBy,
	}

	err := internalerros.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err
}

func (campaign *Campaign) CancelCampaign() {
	campaign.Status = Cancelled
	campaign.UpdatedOn = time.Now()
}

func (campaign *Campaign) DoneCampaign() {
	campaign.Status = Done
	campaign.UpdatedOn = time.Now()
}

func (campaign *Campaign) FailCampaign() {
	campaign.Status = Fail
	campaign.UpdatedOn = time.Now()
}

func (campaign *Campaign) StartCampaign() {
	campaign.Status = Started
	campaign.UpdatedOn = time.Now()
}
