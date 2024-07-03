package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name             = "New Campaign"
	createdBy        = "test_created_by@email.com"
	content          = "Content"
	emails           = []string{"test1@email.com", "test2@email.com"}
	status    string = Pending

	fake = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails, createdBy)

	assert.Equal(name, campaign.Name, "Name must be equal")
	assert.Equal(content, campaign.Content, "Content must be equal")
	assert.Equal(len(emails), len(campaign.Contacts), "Contacts must be equal")
	assert.Equal(emails[0], campaign.Contacts[0].Email, "Emails must be equal")
	assert.Equal(createdBy, campaign.CreatedBy)
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails, createdBy)

	assert.NotNil(campaign.ID, "ID must not be empty")
}

func Test_NewCampaign_CreateOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, emails, createdBy)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(4), content, emails, createdBy)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(25), content, emails, createdBy)

	assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", emails, createdBy)

	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), emails, createdBy)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{}, createdBy)

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"invalid_email"}, createdBy)

	assert.Equal("email is invalid", err.Error())
}

func Test_NewCampaign_StatusMustBePending(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails, createdBy)

	assert.Equal(campaign.Status, status)
}

func Test_NewCampaign_MustValidateCreatedBy(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, emails, "")

	assert.Equal("createdby is invalid", err.Error())
}
