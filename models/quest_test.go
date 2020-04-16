package models

import (
  "github.com/gofrs/uuid"
)

func (ms *ModelSuite) CreateQuest(u4 uuid.UUID, name string, description string, c4 uuid.UUID, so int) *Quest {
  quest := &Quest{
    UserID:      u4,
    Name:        name,
    Description: description,
    CampaignID:  c4,
    SortOrder:   so,
  }

  verrs, err := ms.DB.ValidateAndCreate(quest)
  ms.NoError(err)
  ms.False(verrs.HasAny())

  return quest
}

func (ms *ModelSuite) Test_Quest_Create() {
  count, err := ms.DB.Count("quests")
	ms.NoError(err)
	ms.Equal(0, count)

  u := ms.CreateUser("test@test.com", "test")
  c := ms.CreateCampaign(u.ID, "Test", "Test Description.")

  quest := ms.CreateQuest(
    u.ID,
    "Test",
    "Test description.",
    c.ID,
    1,
  )
  ms.NotNil(quest.ID)

  count, err = ms.DB.Count("quests")
  ms.NoError(err)
  ms.Equal(1, count)
}

func (ms *ModelSuite) Test_Quest_ValidationErrors() {
  count, err := ms.DB.Count("quests")
	ms.NoError(err)
	ms.Equal(0, count)

  u := ms.CreateUser("test@test.com", "test")
  c := ms.CreateCampaign(u.ID, "Test", "Test Description.")

  // test empty quest
  quest := &Quest{}

  verrs, err := ms.DB.ValidateAndCreate(quest)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("quests")
  ms.NoError(err)
  ms.Equal(0, count)

  // test missing UserID
  quest = &Quest{
    Name:        "Test Quest",
    Description: "Test Description",
    CampaignID:  c.ID,
    SortOrder:   1,
  }

  verrs, err = ms.DB.ValidateAndCreate(quest)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("quests")
  ms.NoError(err)
  ms.Equal(0, count)

  // test missing Name
  quest = &Quest{
    UserID:      u.ID,
    Description: "Test Description",
    CampaignID:  c.ID,
    SortOrder:   1,
  }


  verrs, err = ms.DB.ValidateAndCreate(quest)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("quests")
  ms.NoError(err)
  ms.Equal(0, count)

  // test missing Description
  quest = &Quest{
    UserID:      u.ID,
    Name:        "Test Quest",
    CampaignID:  c.ID,
    SortOrder:   1,
  }


  verrs, err = ms.DB.ValidateAndCreate(quest)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("quests")
  ms.NoError(err)
  ms.Equal(0, count)

  // test missing CampaignID
  quest = &Quest{
    UserID:      u.ID,
    Name:        "Test Quest",
    Description: "Test Description",
    SortOrder:   1,
  }


  verrs, err = ms.DB.ValidateAndCreate(quest)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("quests")
  ms.NoError(err)
  ms.Equal(0, count)

  // test missing SortOrder
  quest = &Quest{
    UserID:      u.ID,
    Name:        "Test Quest",
    Description: "Test Description",
    CampaignID:  c.ID,
  }


  verrs, err = ms.DB.ValidateAndCreate(quest)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("quests")
  ms.NoError(err)
  ms.Equal(0, count)
}
