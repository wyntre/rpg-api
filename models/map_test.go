package models

import (
	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) CreateMap(u4 uuid.UUID, name string, description string, q4 uuid.UUID, so int) *Map {
	rpg_map := &Map{
		UserID:      u4,
		Name:        name,
		Description: description,
		QuestID:     q4,
		SortOrder:   so,
	}

	verrs, err := ms.DB.ValidateAndCreate(rpg_map)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	return rpg_map
}

func (ms *ModelSuite) Test_Map_Create() {
	count, err := ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")
	c := ms.CreateCampaign(u.ID, "Test", "Test Description.")
	q := ms.CreateQuest(u.ID, "Test Quest", "Test Quest Description", c.ID, 1)

	rpg_map := ms.CreateMap(
		u.ID,
		"Test",
		"Test description.",
		q.ID,
		1,
	)
	ms.NotNil(rpg_map.ID)

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
	rpg_map := &Map{}

	verrs, err := ms.DB.ValidateAndCreate(rpg_map)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing UserID
	rpg_map = &Map{
		Name:        "Test Map",
		Description: "Test Description",
		QuestID:     q.ID,
		SortOrder:   1,
	}

	verrs, err = ms.DB.ValidateAndCreate(rpg_map)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing Name
	rpg_map = &Map{
		UserID:      u.ID,
		Description: "Test Description",
		QuestID:     q.ID,
		SortOrder:   1,
	}

	verrs, err = ms.DB.ValidateAndCreate(rpg_map)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing Description
	rpg_map = &Map{
		UserID:    u.ID,
		Name:      "Test Map",
		QuestID:   q.ID,
		SortOrder: 1,
	}

	verrs, err = ms.DB.ValidateAndCreate(rpg_map)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing QuestID
	rpg_map = &Map{
		UserID:      u.ID,
		Name:        "Test Map",
		Description: "Test Description",
		SortOrder:   1,
	}

	verrs, err = ms.DB.ValidateAndCreate(rpg_map)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing SortOrder
	rpg_map = &Map{
		UserID:      u.ID,
		Name:        "Test Map",
		Description: "Test Description",
		QuestID:     q.ID,
	}

	verrs, err = ms.DB.ValidateAndCreate(rpg_map)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("maps")
	ms.NoError(err)
	ms.Equal(0, count)
}
