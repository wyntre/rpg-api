package models

func (ms *ModelSuite) CreateTileCategory(name string) *TileCategory {
  t := &TileCategory{
    Name:        name,
  }

  verrs, err := ms.DB.ValidateAndCreate(t)
  ms.NoError(err)
  ms.False(verrs.HasAny())

  return t
}

func (ms *ModelSuite) Test_TileCategory_Create() {
  count, err := ms.DB.Count("tile_categories")
	ms.NoError(err)
	ms.Equal(0, count)

  t := ms.CreateTileCategory("Test Tile Category")
  ms.NotNil(t.ID)

  count, err = ms.DB.Count("tile_categories")
  ms.NoError(err)
  ms.Equal(1, count)
}

func (ms *ModelSuite) Test_TileCategory_ValidationErrors() {
  count, err := ms.DB.Count("tile_categories")
	ms.NoError(err)
	ms.Equal(0, count)

  // test empty map
  t := &TileCategory{}

  verrs, err := ms.DB.ValidateAndCreate(t)
  ms.NoError(err)
  ms.True(verrs.HasAny())

  count, err = ms.DB.Count("tile_categories")
  ms.NoError(err)
  ms.Equal(0, count)
}
