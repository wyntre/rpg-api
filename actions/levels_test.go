package actions

import (
  "net/http"
  "github.com/wyntre/rpg_api/models"
  "github.com/gofrs/uuid"
)

type CreateLevelRequest struct {
  Name 				string 		 `json:"name"`
  Description string 		 `json:"description"`
  MapID       uuid.UUID  `json:"map_id"`
}

type LevelsListResponse struct {
  Levels models.Levels `json:levels`
}

func (as *ActionSuite) CreateLevel(name string, description string, map_id uuid.UUID, token string) *models.Level {
  // create quest
  levelRequest := &CreateLevelRequest{
    Name:        name,
    Description: description,
    MapID:       map_id,
  }

  req := as.JSON("/v1/levels/new")
  req.Headers["Authorization"] = token
  res := req.Post(levelRequest)
  as.Equal(http.StatusCreated, res.Code)

  l := &models.Level{}
  res.Bind(l)
  as.NotNil(l)

  return l
}

func (as *ActionSuite) Test_Levels_Create() {
  token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)

  // create quest
  levelRequest := &CreateLevelRequest{
    Name:        "Test Level",
    Description: "This is a test quest.",
    MapID:       m.ID,
  }

  req := as.JSON("/v1/levels/new")
  req.Headers["Authorization"] = token
  res := req.Post(levelRequest)
  as.Equal(http.StatusCreated, res.Code)

  l := &models.Level{}
  res.Bind(l)
  as.NotNil(l)
}

func (as *ActionSuite) Test_Levels_Create_Fail() {
  token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)

  // without description
  levelRequest := &CreateLevelRequest{
    Name:  "Test Level",
    MapID: m.ID,
  }

  req := as.JSON("/v1/levels/new")
  req.Headers["Authorization"] = token
  res := req.Post(levelRequest)
  as.Equal(http.StatusUnprocessableEntity, res.Code)

  // without name
  levelRequest = &CreateLevelRequest{
    Description: "This is a test level.",
    MapID:       m.ID,
  }

  req = as.JSON("/v1/levels/new")
  req.Headers["Authorization"] = token
  res = req.Post(levelRequest)
  as.Equal(http.StatusUnprocessableEntity, res.Code)

  // without campaign_id
  levelRequest = &CreateLevelRequest{
    Name:        "Test Level",
    Description: "This is a test level.",
  }

  req = as.JSON("/v1/levels/new")
  req.Headers["Authorization"] = token
  res = req.Post(levelRequest)
  as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as *ActionSuite) Test_Levels_Create_No_Token() {
  levelRequest := &CreateLevelRequest{
    Name:        "Test Level",
    Description: "This is a test quest.",
  }

  req := as.JSON("/v1/levels/new")
  res := req.Post(levelRequest)
  as.Equal(http.StatusUnauthorized, res.Code)
}

func (as *ActionSuite) Test_Levels_Show() {
  token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
  l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token)

  req := as.JSON("/v1/levels/show/" + l.ID.String())
  req.Headers["Authorization"] = token
  res := req.Get()
  as.Equal(http.StatusOK, res.Code)

  test_quest := &models.Level{}
  res.Bind(test_quest)
  as.Equal(l.ID, test_quest.ID)
  as.Equal(l.Name, test_quest.Name)
}

func (as *ActionSuite) Test_Levels_Show_Fail() {
  token := as.CreateUser("test@test.com", "test")

  req := as.JSON("/v1/levels/show/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
  req.Headers["Authorization"] = token
  res := req.Get()
  as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Levels_Show_Other_User() {
  token1 := as.CreateUser("test@test.com", "test")
  token2 := as.CreateUser("fake@fake.com", "fake")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token1)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token1)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token1)
  l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token1)

  // request quest with token2
  req := as.JSON("/v1/levels/" + l.ID.String())
  req.Headers["Authorization"] = token2
  res := req.Get()
  as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Levels_Destroy() {
  token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
  l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token)

  // delete quest
  req := as.JSON("/v1/levels/" + l.ID.String())
  req.Headers["Authorization"] = token
  res := req.Delete()
  as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_Levels_Destroy_Fail() {
  token := as.CreateUser("test@test.com", "test")

  // delete missing quest id
  req := as.JSON("/v1/levels/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
  req.Headers["Authorization"] = token
  res := req.Delete()
  as.Equal(http.StatusNotFound, res.Code)

  // delete bad quest id
  req = as.JSON("/v1/levels/abcde")
  req.Headers["Authorization"] = token
  res = req.Delete()
  as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_Levels_Destroy_Other_User() {
  token1 := as.CreateUser("test@test.com", "test")
  token2 := as.CreateUser("fake@fake.com", "fake")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token1)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token1)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token1)
  l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token1)

  // delete quest
  req := as.JSON("/v1/levels/" + l.ID.String())
  req.Headers["Authorization"] = token2
  res := req.Delete()
  as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Levels_List() {
  token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
  as.CreateLevel("Test Level", "Test Level Description", m.ID, token)
  as.CreateLevel("Test Level 2", "Test Level Description", m.ID, token)

  clr := &LevelsListResponse{}
  req := as.JSON("/v1/levels/" + m.ID.String())
  req.Headers["Authorization"] = token
  res := req.Get()
  res.Bind(clr)

  as.Equal(2, len(clr.Levels))
}

func (as *ActionSuite) Test_Levels_Update() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
  l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token)

  l.Name = "Fake Level"

  req := as.JSON("/v1/levels/" + l.ID.String())
  req.Headers["Authorization"] = token
  res := req.Put(l)
  as.Equal(http.StatusOK, res.Code)

  test_level := &models.Level{}
  res.Bind(test_level)
  as.Equal("Fake Level", test_level.Name)
}

func (as *ActionSuite) Test_Levels_Update_Fail() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)
  l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token)

  l.Name = "Fake Level"

  // fail on unknown uuid
  req := as.JSON("/v1/levels/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
  req.Headers["Authorization"] = token
  res := req.Put(l)
  as.Equal(http.StatusNotFound, res.Code)

  // fail on bad uuid
  req = as.JSON("/v1/levels/aaaaaaaa")
  req.Headers["Authorization"] = token
  res = req.Put(l)
  as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_Levels_Update_Other_User() {
  // create valid user
	token1 := as.CreateUser("test@test.com", "test")
  token2 := as.CreateUser("fake@fake.com", "fake")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token1)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token1)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token1)
  l := as.CreateLevel("Test Level", "Test Level Description", m.ID, token1)

  l.Name = "Fake Level"

  req := as.JSON("/v1/levels/" + l.ID.String())
  req.Headers["Authorization"] = token2
  res := req.Put(l)
  as.Equal(http.StatusNotFound, res.Code)
}
