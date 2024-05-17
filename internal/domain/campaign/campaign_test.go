package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var name string = "Campaign X"
var content string = "Body"
var contacts = []string{"john@example.com", "email@example.com"}
var now time.Time = time.Now().Add(-time.Minute)

// AAA arrange act assert
func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func TestNewCampaignIDNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func TestNewCampaignCreatedAtGreaterThanNow(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedAt, now)
}

func TestNewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.NotNil("name is required", err.Error())
}
