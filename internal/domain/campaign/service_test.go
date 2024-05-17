package campaign

import (
	"campaign/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	service := Service{}

	newCampaign := contract.NewCampaign{
		Name:    "test",
		Content: "test content",
		Emails:  []string{"teste@example.com", "teste@example.com"},
	}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) Save(campaign *Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func Test_Create_SaveCampaign(t *testing.T) {

	newCampaign := contract.NewCampaign{
		Name:    "test",
		Content: "test content",
		Emails:  []string{"teste@example.com", "teste@example.com"},
	}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service := Service{Repository: repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}
