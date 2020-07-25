package models

import (
	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) CreateMap(u4 uuid.UUID, name string, description string, q4 uuid.UUID, so int) *Map {
	rpgMap := &Map{
		UserID:      u4,
		Name:        name,
		Description: description,
		QuestID:     q4,
		SortOrder:   so,
	}

	verrs, err := ms.DB.ValidateAndCreate(rpgMap)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	return rpgMap
}

func (ms *ModelSuite) Test_Map_Create() {
	count, err := ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")
	c := ms.CreateCampaign(u.ID, "Test", "Test Description.")
	q := ms.CreateQuest(u.ID, "Test Quest", "Test Quest Description", c.ID, 1)

	rpgMap := ms.CreateMap(
		u.ID,
		"Test",
		"Test description.",
		q.ID,
		1,
	)
	ms.NotNil(rpgMap.ID)

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_Map_ValidationErrors() {
	count, err := ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")
	c := ms.CreateCampaign(u.ID, "Test", "Test Description.")
	q := ms.CreateQuest(u.ID, "Test Quest", "Test Quest Description", c.ID, 1)

	// test empty map
	rpgMap := &Map{}

	verrs, err := ms.DB.ValidateAndCreate(rpgMap)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing UserID
	rpgMap = &Map{
		Name:        "Test Map",
		Description: "Test Description",
		QuestID:     q.ID,
		SortOrder:   1,
	}

	verrs, err = ms.DB.ValidateAndCreate(rpgMap)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing Name
	rpgMap = &Map{
		UserID:      u.ID,
		Description: "Test Description",
		QuestID:     q.ID,
		SortOrder:   1,
	}

	verrs, err = ms.DB.ValidateAndCreate(rpgMap)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing Description
	rpgMap = &Map{
		UserID:    u.ID,
		Name:      "Test Map",
		QuestID:   q.ID,
		SortOrder: 1,
	}

	verrs, err = ms.DB.ValidateAndCreate(rpgMap)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing QuestID
	rpgMap = &Map{
		UserID:      u.ID,
		Name:        "Test Map",
		Description: "Test Description",
		SortOrder:   1,
	}

	verrs, err = ms.DB.ValidateAndCreate(rpgMap)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing SortOrder
	rpgMap = &Map{
		UserID:      u.ID,
		Name:        "Test Map",
		Description: "Test Description",
		QuestID:     q.ID,
	}

	verrs, err = ms.DB.ValidateAndCreate(rpgMap)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)
}
