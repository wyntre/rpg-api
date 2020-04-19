package models

import (
	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) CreateTile(u4 uuid.UUID, l4 uuid.UUID, x int, y int, t4 uuid.UUID) *Tile {
	t := &Tile{
		UserID:     u4,
		LevelID:    l4,
		X:          x,
		Y:          y,
		TileTypeID: t4,
	}

	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	return t
}

func (ms *ModelSuite) Test_Tile_Create() {
	count, err := ms.DB.Count("tile_types")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")
	c := ms.CreateCampaign(u.ID, "Test", "Test Description.")
	q := ms.CreateQuest(u.ID, "Test Quest", "Test Quest Description", c.ID, 1)
	m := ms.CreateMap(u.ID, "Test Map", "Test Map Description", q.ID, 1)
	l := ms.CreateLevel(u.ID, "Test Level", "Test Level Description", m.ID, 1)

	tc := ms.CreateTileCategory("Test Tile Category")
	tt := ms.CreateTileType("Test", tc.ID)

	t := &Tile{
		UserID:     u.ID,
		LevelID:    l.ID,
		X:          1,
		Y:          1,
		TileTypeID: tt.ID,
	}

	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotNil(t.ID)

	count, err = ms.DB.Count("tiles")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_Tile_ValidationErrors() {
	count, err := ms.DB.Count("tiles")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")
	c := ms.CreateCampaign(u.ID, "Test", "Test Description.")
	q := ms.CreateQuest(u.ID, "Test Quest", "Test Quest Description", c.ID, 1)
	m := ms.CreateMap(u.ID, "Test Map", "Test Map Description", q.ID, 1)
	l := ms.CreateLevel(u.ID, "Test Level", "Test Level Description", m.ID, 1)

	tc := ms.CreateTileCategory("Test Tile Category")
	tt := ms.CreateTileType("Test", tc.ID)

	// test empty map
	t := &Tile{}

	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("tiles")
	ms.NoError(err)
	ms.Equal(0, count)

	// test empty user_id
	t = &Tile{
		LevelID:    l.ID,
		X:          1,
		Y:          1,
		TileTypeID: tt.ID,
	}

	verrs, err = ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("tiles")
	ms.NoError(err)
	ms.Equal(0, count)

	// test empty level_id
	t = &Tile{
		UserID:     u.ID,
		X:          1,
		Y:          1,
		TileTypeID: tt.ID,
	}

	verrs, err = ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("tiles")
	ms.NoError(err)
	ms.Equal(0, count)

	// test empty x
	t = &Tile{
		UserID:     u.ID,
		LevelID:    l.ID,
		Y:          1,
		TileTypeID: tt.ID,
	}

	verrs, err = ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("tiles")
	ms.NoError(err)
	ms.Equal(0, count)

	// test empty y
	t = &Tile{
		UserID:     u.ID,
		LevelID:    l.ID,
		X:          1,
		TileTypeID: tt.ID,
	}

	verrs, err = ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("tiles")
	ms.NoError(err)
	ms.Equal(0, count)

	// test empty tile type
	t = &Tile{
		UserID:  u.ID,
		LevelID: l.ID,
		X:       1,
		Y:       1,
	}

	verrs, err = ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("tiles")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_Tile_Create_Duplicate() {
	count, err := ms.DB.Count("tiles")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")
	c := ms.CreateCampaign(u.ID, "Test", "Test Description.")
	q := ms.CreateQuest(u.ID, "Test Quest", "Test Quest Description", c.ID, 1)
	m := ms.CreateMap(u.ID, "Test Map", "Test Map Description", q.ID, 1)
	l := ms.CreateLevel(u.ID, "Test Level", "Test Level Description", m.ID, 1)

	tc := ms.CreateTileCategory("Test Tile Category")
	tt := ms.CreateTileType("Test", tc.ID)

	t := &Tile{
		UserID:     u.ID,
		LevelID:    l.ID,
		X:          1,
		Y:          1,
		TileTypeID: tt.ID,
	}

	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotNil(t.ID)

	count, err = ms.DB.Count("tiles")
	ms.NoError(err)
	ms.Equal(1, count)

	d := &Tile{
		UserID:     u.ID,
		LevelID:    l.ID,
		X:          1,
		Y:          1,
		TileTypeID: tt.ID,
	}

	verrs, err = ms.DB.ValidateAndCreate(d)
	ms.NotNil(err)
	ms.False(verrs.HasAny())

	count, err = ms.DB.Count("tiles")
	ms.NoError(err)
	ms.Equal(1, count)
}
