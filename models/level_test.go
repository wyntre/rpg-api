package models

import (
  "github.com/gofrs/uuid"
)

func (ms *ModelSuite) CreateLevel(u4 uuid.UUID, name string, description string, m4 uuid.UUID, so int) *Level {
  l := &Level{
    UserID:      u4,
    Name:        name,
    Description: description,
    MapID:     m4,
    SortOrder:   so,
  }

  verrs, err := ms.DB.ValidateAndCreate(l)
  ms.NoError(err)
  ms.False(verrs.HasAny())

  return l
}

func (ms *ModelSuite) Test_Level_Create() {
  count, err := ms.DB.Count("levels")
	ms.NoError(err)
	ms.Equal(0, count)

  u := ms.CreateUser("test@test.com", "test")
  c := ms.CreateCampaign(u.ID, "Test", "Test Description.")
  q := ms.CreateQuest(u.ID, "Test Quest", "Test Quest Description", c.ID, 1)
  m := ms.CreateMap(u.ID, "Test Map", "Test Map Description", q.ID, 1)

  l := ms.CreateLevel(
    u.ID,
    "Test",
    "Test description.",
    m.ID,
    1,
  )
  ms.NotNil(l.ID)

  count, err = ms.DB.Count("levels")
  ms.NoError(err)
  ms.Equal(1, count)
}

func (ms *ModelSuite) Test_Level_ValidationErrors() {
  count, err := ms.DB.Count("levels")
	ms.NoError(err)
	ms.Equal(0, count)

  u := ms.CreateUser("test@test.com", "test")
  c := ms.CreateCampaign(u.ID, "Test", "Test Description.")
  q := ms.CreateQuest(u.ID, "Test Quest", "Test Quest Description", c.ID, 1)
  m := ms.CreateMap(u.ID, "Test Map", "Test Map Description", q.ID, 1)

  // test empty level
  l := &Level{}

  verrs, err := ms.DB.ValidateAndCreate(l)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("levels")
  ms.NoError(err)
  ms.Equal(0, count)

  // test missing UserID
  l = &Level{
    Name:        "Test Level",
    Description: "Test Description",
    MapID:       m.ID,
    SortOrder:   1,
  }

  verrs, err = ms.DB.ValidateAndCreate(l)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("levels")
  ms.NoError(err)
  ms.Equal(0, count)

  // test missing Name
  l = &Level{
    UserID:      u.ID,
    Description: "Test Description",
    MapID:       m.ID,
    SortOrder:   1,
  }


  verrs, err = ms.DB.ValidateAndCreate(l)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("levels")
  ms.NoError(err)
  ms.Equal(0, count)

  // test missing Description
  l = &Level{
    UserID:      u.ID,
    Name:        "Test Level",
    MapID:       m.ID,
    SortOrder:   1,
  }


  verrs, err = ms.DB.ValidateAndCreate(l)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("levels")
  ms.NoError(err)
  ms.Equal(0, count)

  // test missing QuestID
  l = &Level{
    UserID:      u.ID,
    Name:        "Test Level",
    Description: "Test Description",
    SortOrder:   1,
  }


  verrs, err = ms.DB.ValidateAndCreate(l)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("levels")
  ms.NoError(err)
  ms.Equal(0, count)

  // test missing SortOrder
  l = &Level{
    UserID:      u.ID,
    Name:        "Test Level",
    Description: "Test Description",
    MapID:       m.ID,
  }


  verrs, err = ms.DB.ValidateAndCreate(l)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("levels")
  ms.NoError(err)
  ms.Equal(0, count)
}
