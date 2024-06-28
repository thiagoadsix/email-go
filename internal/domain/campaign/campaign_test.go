package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "New Campaign"
	content = "Content"
	emails  = []string{"test1@email.com", "test2@email.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails)

	assert.Equal(name, campaign.Name, "Name must be equal")
	assert.Equal(content, campaign.Content, "Content must be equal")
	assert.Equal(len(emails), len(campaign.Contacts), "Contacts must be equal")
	assert.Equal(emails[0], campaign.Contacts[0].Email, "Emails must be equal")
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails)

	assert.NotNil(campaign.ID, "ID must not be empty")
}

func Test_NewCampaign_CreateOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, emails)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_NameMustNotBeEmpty(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, emails)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_ContentMustNotBeEmpty(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", emails)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_ContactsMustBeNotEmpty(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required", err.Error())
}
