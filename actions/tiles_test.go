package actions

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/wyntre/rpg_api/models"
)

type CreateTileRequest struct {
	LevelID    uuid.UUID `json:"level_id"`
	TileTypeID uuid.UUID `json:"tile_type_id"`
	X          int       `json:"x"`
	Y          int       `json:"y"`
}

type TilesListResponse struct {
	Tiles models.Tiles `json:"tiles"`
}

func (as *ActionSuite) CreateTile(levelID uuid.UUID, tileTypeID uuid.UUID, x int, y int, token string) *models.Tile {
	// create Tile Category
	TileRequest := &CreateTileRequest{
		LevelID:    levelID,
		TileTypeID: tileTypeID,
		X:          x,
		Y:          y,
	}

	req := as.JSON("/v1/tiles/new")
	req.Headers["Authorization"] = token
	res := req.Post(TileRequest)
	as.Equal(http.StatusCreated, res.Code)

	t := &models.Tile{}
	res.Bind(t)
	as.NotNil(t)

	return t
}

func (as *ActionSuite) Test_Tiles_Create() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)
	tt := as.CreateTileType("Test Type", tc.ID, token)

	c := as.CreateCampaign("Test Campaign", "Test Campaign Description", token)
	q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
	m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
	l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token)

	// create Tile Category
	TileRequest := &CreateTileRequest{
		LevelID:    l.ID,
		TileTypeID: tt.ID,
		X:          1,
		Y:          1,
	}

	req := as.JSON("/v1/tiles/new")
	req.Headers["Authorization"] = token
	res := req.Post(TileRequest)
	as.Equal(http.StatusCreated, res.Code)

	t := &models.Tile{}
	res.Bind(t)
	as.NotNil(t)
}

func (as *ActionSuite) Test_Tiles_Create_Fail() {
	token := as.CreateUser("test@test.com", "test")

	levelID, _ := uuid.FromString("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	tileTypeID, _ := uuid.FromString("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")

	// without name
	TileRequest := &CreateTileRequest{
		LevelID:    levelID,
		TileTypeID: tileTypeID,
		X:          1,
	}

	req := as.JSON("/v1/tiles/new")
	req.Headers["Authorization"] = token
	res := req.Post(TileRequest)
	as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as *ActionSuite) Test_Tiles_Create_No_Token() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)
	tt := as.CreateTileType("Test Type", tc.ID, token)

	c := as.CreateCampaign("Test Campaign", "Test Campaign Description", token)
	q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
	m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
	l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token)

	// create Tile Category
	TileRequest := &CreateTileRequest{
		LevelID:    l.ID,
		TileTypeID: tt.ID,
		X:          1,
		Y:          1,
	}

	req := as.JSON("/v1/tiles/new")
	res := req.Post(TileRequest)
	as.Equal(http.StatusUnauthorized, res.Code)
}

func (as *ActionSuite) Test_Tiles_Show() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)
	tt := as.CreateTileType("Test Type", tc.ID, token)

	c := as.CreateCampaign("Test Campaign", "Test Campaign Description", token)
	q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
	m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
	l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token)
	t := as.CreateTile(l.ID, tt.ID, 1, 1, token)

	req := as.JSON("/v1/tiles/" + t.ID.String())
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)

	testTile := &models.Tile{}
	res.Bind(testTile)
	as.Equal(t.ID, testTile.ID)
	as.Equal(t.X, testTile.X)
	as.Equal(t.Y, testTile.Y)
}

func (as *ActionSuite) Test_Tiles_Show_Fail() {
	token := as.CreateUser("test@test.com", "test")

	req := as.JSON("/v1/tiles/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Tiles_Destroy() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)
	tt := as.CreateTileType("Test Type", tc.ID, token)

	c := as.CreateCampaign("Test Campaign", "Test Campaign Description", token)
	q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
	m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
	l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token)
	t := as.CreateTile(l.ID, tt.ID, 1, 1, token)

	req := as.JSON("/v1/tiles/" + t.ID.String())
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)

	// delete quest
	req = as.JSON("/v1/tiles/" + t.ID.String())
	req.Headers["Authorization"] = token
	res = req.Delete()
	as.Equal(http.StatusAccepted, res.Code)

	req = as.JSON("/v1/tiles/" + t.ID.String())
	req.Headers["Authorization"] = token
	res = req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Tiles_Destroy_Fail() {
	token := as.CreateUser("test@test.com", "test")

	// delete missing quest id
	req := as.JSON("/v1/tiles/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Delete()
	as.Equal(http.StatusNotFound, res.Code)

	// delete bad quest id
	req = as.JSON("/v1/tiles/abcde")
	req.Headers["Authorization"] = token
	res = req.Delete()
	as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_Tiles_List() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)
	tt := as.CreateTileType("Test Type", tc.ID, token)

	c := as.CreateCampaign("Test Campaign", "Test Campaign Description", token)
	q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
	m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
	l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token)

	as.CreateTile(l.ID, tt.ID, 1, 1, token)
	as.CreateTile(l.ID, tt.ID, 1, 2, token)
	as.CreateTile(l.ID, tt.ID, 1, 3, token)
	as.CreateTile(l.ID, tt.ID, 2, 1, token)
	as.CreateTile(l.ID, tt.ID, 2, 2, token)
	as.CreateTile(l.ID, tt.ID, 2, 3, token)

	clr := &TilesListResponse{}
	req := as.JSON("/v1/levels/" + l.ID.String() + "/tiles")
	req.Headers["Authorization"] = token
	res := req.Get()
	res.Bind(clr)

	as.Equal(6, len(clr.Tiles))
}

func (as *ActionSuite) Test_Tiles_Update() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)
	tt1 := as.CreateTileType("Test Type", tc.ID, token)
	tt2 := as.CreateTileType("Test Type 2", tc.ID, token)

	c := as.CreateCampaign("Test Campaign", "Test Campaign Description", token)
	q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
	m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
	l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token)
	t := as.CreateTile(l.ID, tt1.ID, 1, 1, token)

	t.TileTypeID = tt2.ID

	req := as.JSON("/v1/tiles/" + t.ID.String())
	req.Headers["Authorization"] = token
	res := req.Put(t)
	as.Equal(http.StatusAccepted, res.Code)

	testTile := &models.Tile{}
	res.Bind(testTile)
	as.Equal(tt2.ID, testTile.TileTypeID)
}

func (as *ActionSuite) Test_Tiles_Update_Fail() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)
	tt1 := as.CreateTileType("Test Type", tc.ID, token)
	tt2 := as.CreateTileType("Test Type 2", tc.ID, token)

	c := as.CreateCampaign("Test Campaign", "Test Campaign Description", token)
	q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
	m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
	l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token)
	t := as.CreateTile(l.ID, tt1.ID, 1, 1, token)

	t.TileTypeID = tt2.ID

	// fail on unknown uuid
	req := as.JSON("/v1/tiles/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Put(t)
	as.Equal(http.StatusNotFound, res.Code)

	// fail on bad uuid
	req = as.JSON("/v1/tiles/aaaaaaaa")
	req.Headers["Authorization"] = token
	res = req.Put(t)
	as.Equal(http.StatusInternalServerError, res.Code)
}
