package models

import (
	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) CreateCharacter(u4 uuid.UUID, name string, description string) *Character {
	character := &Character{
		UserID:      u4,
		Name:        name,
		Description: description,
	}

	verrs, err := ms.DB.ValidateAndCreate(character)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	return character
}

func (ms *ModelSuite) Test_Character_Create() {
	count, err := ms.DB.Count("characters")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")

	character := ms.CreateCharacter(
		u.ID,
		"Test",
		"Test description.",
	)
	ms.NotNil(character.ID)

	count, err = ms.DB.Count("characters")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_Character_ValidationErrors() {
	count, err := ms.DB.Count("characters")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")

	// test empty character
	character := &Character{}

	verrs, err := ms.DB.ValidateAndCreate(character)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("characters")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing UserID
	character = &Character{
		Name:        "Test Character",
		Description: "Test description.",
	}

	verrs, err = ms.DB.ValidateAndCreate(character)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("characters")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing Name
	character = &Character{
		UserID:      u.ID,
		Description: "Test description.",
	}

	verrs, err = ms.DB.ValidateAndCreate(character)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("characters")
	ms.NoError(err)
	ms.Equal(0, count)

	// test missing Description
	character = &Character{
		UserID: u.ID,
		Name:   "Test Character",
	}

	verrs, err = ms.DB.ValidateAndCreate(character)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("characters")
	ms.NoError(err)
	ms.Equal(0, count)
}
