package models

import (
	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) CreateCampaign(u4 uuid.UUID, name string, description string) *Campaign {
	campaign := &Campaign{
		UserID:      u4,
		Name:        name,
		Description: description,
	}

	verrs, err := ms.DB.ValidateAndCreate(campaign)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	return campaign
}

func (ms *ModelSuite) Test_Campaign_Create() {
	count, err := ms.DB.Count("campaigns")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")

	campaign := ms.CreateCampaign(
		u.ID,
		"Test",
		"Test description.",
	)
	ms.NotNil(campaign.ID)

	count, err = ms.DB.Count("campaigns")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_Campaign_ValidationErrors() {
	count, err := ms.DB.Count("campaigns")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")

	// test empty campaign
	campaign := &Campaign{}

	verrs, err := ms.DB.ValidateAndCreate(campaign)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("campaigns")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing UserID
	campaign = &Campaign{
		Name:        "Test campaign",
		Description: "Test description.",
	}

	verrs, err = ms.DB.ValidateAndCreate(campaign)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("campaigns")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing Name
	campaign = &Campaign{
		UserID:      u.ID,
		Description: "Test description.",
	}

	verrs, err = ms.DB.ValidateAndCreate(campaign)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("campaigns")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing Description
	campaign = &Campaign{
		UserID: u.ID,
		Name:   "Test campaign",
	}

	verrs, err = ms.DB.ValidateAndCreate(campaign)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("campaigns")
	ms.NoError(err)
	ms.Equal(0, count)
}
