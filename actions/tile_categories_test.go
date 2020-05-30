package actions

import (
	"github.com/wyntre/rpg_api/models"
	"net/http"
)

type CreateTileCategoryRequest struct {
	Name string `json:"name"`
}

type TileCategoriesListResponse struct {
	TileCategories models.TileCategories `json:"tile_categories"`
}

func (as *ActionSuite) CreateTileCategory(name string, token string) *models.TileCategory {
	// create quest
	tileCategoryRequest := &CreateTileCategoryRequest{
		Name: name,
	}

	req := as.JSON("/v1/tile_categories/new")
	req.Headers["Authorization"] = token
	res := req.Post(tileCategoryRequest)
	as.Equal(http.StatusCreated, res.Code)

	tc := &models.TileCategory{}
	res.Bind(tc)
	as.NotNil(tc)

	return tc
}

func (as *ActionSuite) Test_TileCategories_Create() {
	token := as.CreateUser("test@test.com", "test")

	// create quest
	tileCategoryRequest := &CreateTileCategoryRequest{
		Name: "Test TileCategory",
	}

	req := as.JSON("/v1/tile_categories/new")
	req.Headers["Authorization"] = token
	res := req.Post(tileCategoryRequest)
	as.Equal(http.StatusCreated, res.Code)

	tc := &models.TileCategory{}
	res.Bind(tc)
	as.NotNil(tc)
}

func (as *ActionSuite) Test_TileCategories_Create_Fail() {
	token := as.CreateUser("test@test.com", "test")

	// without description
	tileCategoryRequest := &CreateTileCategoryRequest{}

	req := as.JSON("/v1/tile_categories/new")
	req.Headers["Authorization"] = token
	res := req.Post(tileCategoryRequest)
	as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as *ActionSuite) Test_TileCategories_Create_No_Token() {
	tileCategoryRequest := &CreateTileCategoryRequest{
		Name: "Test TileCategory",
	}

	req := as.JSON("/v1/tile_categories/new")
	res := req.Post(tileCategoryRequest)
	as.Equal(http.StatusUnauthorized, res.Code)
}

func (as *ActionSuite) Test_TileCategories_Show() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test TileCategory", token)

	req := as.JSON("/v1/tile_categories/show/" + tc.ID.String())
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)

	test_category := &models.TileCategory{}
	res.Bind(test_category)
	as.Equal(tc.ID, test_category.ID)
	as.Equal(tc.Name, test_category.Name)
}

func (as *ActionSuite) Test_TileCategories_Show_Fail() {
	token := as.CreateUser("test@test.com", "test")

	req := as.JSON("/v1/tile_categories/show/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_TileCategories_Destroy() {
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test TileCategory", token)

	req := as.JSON("/v1/tile_categories/show/" + tc.ID.String())
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)

	// delete quest
	req = as.JSON("/v1/tile_categories/" + tc.ID.String())
	req.Headers["Authorization"] = token
	res = req.Delete()
	as.Equal(http.StatusOK, res.Code)

	req = as.JSON("/v1/tile_categories/show/" + tc.ID.String())
	req.Headers["Authorization"] = token
	res = req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_TileCategories_Destroy_Fail() {
	token := as.CreateUser("test@test.com", "test")

	// delete missing quest id
	req := as.JSON("/v1/tile_categories/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Delete()
	as.Equal(http.StatusNotFound, res.Code)

	// delete bad quest id
	req = as.JSON("/v1/tile_categories/abcde")
	req.Headers["Authorization"] = token
	res = req.Delete()
	as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_TileCategories_List() {
	token := as.CreateUser("test@test.com", "test")

	as.CreateTileCategory("Test TileCategory", token)
	as.CreateTileCategory("Test TileCategory2", token)

	clr := &TileCategoriesListResponse{}
	req := as.JSON("/v1/tile_categories")
	req.Headers["Authorization"] = token
	res := req.Get()
	res.Bind(clr)

	as.Equal(2, len(clr.TileCategories))
}

func (as *ActionSuite) Test_TileCategories_Update() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test TileCategory", token)

	tc.Name = "Fake TileCategory"

	req := as.JSON("/v1/tile_categories/" + tc.ID.String())
	req.Headers["Authorization"] = token
	res := req.Put(tc)
	as.Equal(http.StatusOK, res.Code)

	test_level := &models.TileCategory{}
	res.Bind(test_level)
	as.Equal("Fake TileCategory", test_level.Name)
}

func (as *ActionSuite) Test_TileCategories_Update_Fail() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	tc := as.CreateTileCategory("Test TileCategory", token)

	tc.Name = "Fake TileCategory"

	// fail on unknown uuid
	req := as.JSON("/v1/tile_categories/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Put(tc)
	as.Equal(http.StatusNotFound, res.Code)

	// fail on bad uuid
	req = as.JSON("/v1/tile_categories/aaaaaaaa")
	req.Headers["Authorization"] = token
	res = req.Put(tc)
	as.Equal(http.StatusInternalServerError, res.Code)
}
