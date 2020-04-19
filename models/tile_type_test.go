package models

import (
	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) CreateTileType(name string, tc4 uuid.UUID) *TileType {
	t := &TileType{
		Name:           name,
		TileCategoryID: tc4,
	}

	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	return t
}

func (ms *ModelSuite) Test_TileType_Create() {
	count, err := ms.DB.Count("tile_types")
	ms.NoError(err)
	ms.Equal(0, count)

	tc := ms.CreateTileCategory("Test Tile Category")

	t := &TileType{
		Name:           "Test Tile Type",
		TileCategoryID: tc.ID,
	}

	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotNil(t.ID)

	count, err = ms.DB.Count("tile_types")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_TileType_ValidationErrors() {
	count, err := ms.DB.Count("tile_types")
	ms.NoError(err)
	ms.Equal(0, count)

	// test empty map
	t := &TileType{
		Name: "Test Tile Type",
	}

	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("tile_types")
	ms.NoError(err)
	ms.Equal(0, count)

	tc := ms.CreateTileCategory("Test Tile Category")
	// test empty map
	t = &TileType{
		TileCategoryID: tc.ID,
	}

	verrs, err = ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("tile_types")
	ms.NoError(err)
	ms.Equal(0, count)
}
