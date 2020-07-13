package actions

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/wyntre/rpg_api/models"
)

type CreateTileTypeRequest struct {
	Name           string    `json:"name"`
	TileCategoryID uuid.UUID `json:"tile_category_id"`
}

type TileTypesListResponse struct {
	TileTypes models.TileTypes `json:"tile_types"`
}

func (as *ActionSuite) CreateTileType(name string, tileCategoryID uuid.UUID, token string) *models.TileType {
	// create Tile Category
	TileTypeRequest := &CreateTileTypeRequest{
		Name:           name,
		TileCategoryID: tileCategoryID,
	}

	req := as.JSON("/v1/tile_types/new")
	req.Headers["Authorization"] = token
	res := req.Post(TileTypeRequest)
	as.Equal(http.StatusCreated, res.Code)

	tc := &models.TileType{}
	res.Bind(tc)
	as.NotNil(tc)

	return tc
}

func (as *ActionSuite) Test_TileTypes_Create() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)

	// create Tile Category
	TileTypeRequest := &CreateTileTypeRequest{
		Name:           "Test TileType",
		TileCategoryID: tc.ID,
	}

	req := as.JSON("/v1/tile_types/new")
	req.Headers["Authorization"] = token
	res := req.Post(TileTypeRequest)
	as.Equal(http.StatusCreated, res.Code)

	tt := &models.TileType{}
	res.Bind(tt)
	as.NotNil(tt)
}

func (as *ActionSuite) Test_TileTypes_Create_Fail() {
	token := as.CreateUser("test@test.com", "test")

	// without name
	TileTypeRequest := &CreateTileTypeRequest{}

	req := as.JSON("/v1/tile_types/new")
	req.Headers["Authorization"] = token
	res := req.Post(TileTypeRequest)
	as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as *ActionSuite) Test_TileTypes_Create_No_Token() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)

	TileTypeRequest := &CreateTileTypeRequest{
		Name:           "Test TileType",
		TileCategoryID: tc.ID,
	}

	req := as.JSON("/v1/tile_types/new")
	res := req.Post(TileTypeRequest)
	as.Equal(http.StatusUnauthorized, res.Code)
}

func (as *ActionSuite) Test_TileTypes_Show() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)

	tt := as.CreateTileType("Test TileType", tc.ID, token)

	req := as.JSON("/v1/tile_types/" + tt.ID.String())
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)

	testType := &models.TileType{}
	res.Bind(testType)
	as.Equal(tt.ID, testType.ID)
	as.Equal(tt.Name, testType.Name)
}

func (as *ActionSuite) Test_TileTypes_Show_Fail() {
	token := as.CreateUser("test@test.com", "test")

	req := as.JSON("/v1/tile_types/show/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_TileTypes_Destroy() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)

	tt := as.CreateTileType("Test TileType", tc.ID, token)

	req := as.JSON("/v1/tile_types/" + tt.ID.String())
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)

	// delete quest
	req = as.JSON("/v1/tile_types/" + tt.ID.String())
	req.Headers["Authorization"] = token
	res = req.Delete()
	as.Equal(http.StatusAccepted, res.Code)

	req = as.JSON("/v1/tile_types/" + tt.ID.String())
	req.Headers["Authorization"] = token
	res = req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_TileTypes_Destroy_Fail() {
	token := as.CreateUser("test@test.com", "test")

	// delete missing quest id
	req := as.JSON("/v1/tile_types/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Delete()
	as.Equal(http.StatusNotFound, res.Code)

	// delete bad quest id
	req = as.JSON("/v1/tile_types/abcde")
	req.Headers["Authorization"] = token
	res = req.Delete()
	as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_TileTypes_List() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)

	as.CreateTileType("Test TileType", tc.ID, token)
	as.CreateTileType("Test TileType2", tc.ID, token)

	clr := &TileTypesListResponse{}
	req := as.JSON("/v1/tile_categories/" + tc.ID.String() + "/tile_types")
	req.Headers["Authorization"] = token
	res := req.Get()
	res.Bind(clr)

	as.Equal(2, len(clr.TileTypes))
}

func (as *ActionSuite) Test_TileTypes_Update() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)

	tt := as.CreateTileType("Test TileType", tc.ID, token)

	tt.Name = "Fake TileType"

	req := as.JSON("/v1/tile_types/" + tt.ID.String())
	req.Headers["Authorization"] = token
	res := req.Put(tt)
	as.Equal(http.StatusAccepted, res.Code)

	testType := &models.TileType{}
	res.Bind(testType)
	as.Equal("Fake TileType", testType.Name)
}

func (as *ActionSuite) Test_TileTypes_Update_Fail() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test Category", token)

	tt := as.CreateTileType("Test TileType", tc.ID, token)

	tt.Name = "Fake TileType"

	// fail on unknown uuid
	req := as.JSON("/v1/tile_types/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Put(tt)
	as.Equal(http.StatusNotFound, res.Code)

	// fail on bad uuid
	req = as.JSON("/v1/tile_types/aaaaaaaa")
	req.Headers["Authorization"] = token
	res = req.Put(tt)
	as.Equal(http.StatusInternalServerError, res.Code)
}
