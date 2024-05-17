package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Campaign struct {
	ID        string
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contact
}

type Contact struct {
	Email string
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("name must not be empty")
	} else if content == "" {
		return nil, errors.New("content must not be empty")
	} else if len(emails) == 0 {
		return nil, errors.New("emails must not be empty")
	}

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}, nil
}
