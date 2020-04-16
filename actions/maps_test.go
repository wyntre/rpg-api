package actions

import (
  "net/http"
  "github.com/wyntre/rpg_api/models"
  "github.com/gofrs/uuid"
)

type CreateMapRequest struct {
  Name 				string 		 `json:"name"`
  Description string 		 `json:"description"`
  QuestID     uuid.UUID  `json:"quest_id"`
}

type MapsListResponse struct {
  Maps models.Maps `json:maps`
}

func (as *ActionSuite) CreateMap(name string, description string, quest_id uuid.UUID, token string) *models.Map {
  // create quest
  mapRequest := &CreateMapRequest{
    Name: "Test Map",
    Description: "This is a test quest.",
    QuestID: quest_id,
  }

  req := as.JSON("/v1/maps/new")
  req.Headers["Authorization"] = token
  res := req.Post(mapRequest)
  as.Equal(http.StatusCreated, res.Code)

  rpg_map := &models.Map{}
  res.Bind(rpg_map)
  as.NotNil(rpg_map)

  return rpg_map
}

func (as *ActionSuite) Test_Maps_Create() {
  token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)

  // create quest
  mapRequest := &CreateMapRequest{
    Name: "Test Map",
    Description: "This is a test quest.",
    QuestID: q.ID,
  }

  req := as.JSON("/v1/maps/new")
  req.Headers["Authorization"] = token
  res := req.Post(mapRequest)
  as.Equal(http.StatusCreated, res.Code)

  rpg_map := &models.Map{}
  res.Bind(rpg_map)
  as.NotNil(rpg_map)
}

func (as *ActionSuite) Test_Maps_Create_Fail() {
  token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)

  // without description
  mapRequest := &CreateMapRequest{
    Name: "Test Map",
    QuestID: q.ID,
  }

  req := as.JSON("/v1/maps/new")
  req.Headers["Authorization"] = token
  res := req.Post(mapRequest)
  as.Equal(http.StatusUnprocessableEntity, res.Code)

  // without name
  mapRequest = &CreateMapRequest{
    Description: "This is a test map.",
    QuestID: q.ID,
  }

  req = as.JSON("/v1/maps/new")
  req.Headers["Authorization"] = token
  res = req.Post(mapRequest)
  as.Equal(http.StatusUnprocessableEntity, res.Code)

  // without campaign_id
  mapRequest = &CreateMapRequest{
    Name: "Test Map",
    Description: "This is a test map.",
  }

  req = as.JSON("/v1/maps/new")
  req.Headers["Authorization"] = token
  res = req.Post(mapRequest)
  as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as *ActionSuite) Test_Maps_Create_No_Token() {
  questRequest := &CreateMapRequest{
    Name: "Test Map",
    Description: "This is a test quest.",
  }

  req := as.JSON("/v1/maps/new")
  res := req.Post(questRequest)
  as.Equal(http.StatusUnauthorized, res.Code)
}

func (as *ActionSuite) Test_Maps_Show() {
  token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)

  req := as.JSON("/v1/maps/show/" + m.ID.String())
  req.Headers["Authorization"] = token
  res := req.Get()
  as.Equal(http.StatusOK, res.Code)

  test_quest := &models.Map{}
  res.Bind(test_quest)
  as.Equal(m.ID, test_quest.ID)
  as.Equal(m.Name, test_quest.Name)
}

func (as *ActionSuite) Test_Maps_Show_Fail() {
  token := as.CreateUser("test@test.com", "test")

  req := as.JSON("/v1/maps/show/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
  req.Headers["Authorization"] = token
  res := req.Get()
  as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Maps_Show_Other_User() {
  token1 := as.CreateUser("test@test.com", "test")
  token2 := as.CreateUser("fake@fake.com", "fake")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token1)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token1)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token1)

  // request quest with token2
  req := as.JSON("/v1/maps/" + m.ID.String())
  req.Headers["Authorization"] = token2
  res := req.Get()
  as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Maps_Destroy() {
  token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)

  // delete quest
  req := as.JSON("/v1/maps/" + m.ID.String())
  req.Headers["Authorization"] = token
  res := req.Delete()
  as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_Maps_Destroy_Fail() {
  token := as.CreateUser("test@test.com", "test")

  // delete missing quest id
  req := as.JSON("/v1/maps/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
  req.Headers["Authorization"] = token
  res := req.Delete()
  as.Equal(http.StatusNotFound, res.Code)

  // delete bad quest id
  req = as.JSON("/v1/maps/abcde")
  req.Headers["Authorization"] = token
  res = req.Delete()
  as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_Maps_Destroy_Other_User() {
  token1 := as.CreateUser("test@test.com", "test")
  token2 := as.CreateUser("fake@fake.com", "fake")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token1)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token1)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token1)

  // delete quest
  req := as.JSON("/v1/maps/" + m.ID.String())
  req.Headers["Authorization"] = token2
  res := req.Delete()
  as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Maps_List() {
  token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  as.CreateMap("Test Map", "Test Map Description", q.ID, token)
  as.CreateMap("Test Map 2", "Test Map Description", q.ID, token)

  clr := &MapsListResponse{}
  req := as.JSON("/v1/maps/" + q.ID.String())
  req.Headers["Authorization"] = token
  res := req.Get()
  res.Bind(clr)

  as.Equal(2, len(clr.Maps))
}

func (as *ActionSuite) Test_Maps_Update() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)

  m.Name = "Fake Map"

  req := as.JSON("/v1/maps/" + m.ID.String())
  req.Headers["Authorization"] = token
  res := req.Put(m)
  as.Equal(http.StatusOK, res.Code)

  test_map := &models.Map{}
  res.Bind(test_map)
  as.Equal("Fake Map", test_map.Name)
}

func (as *ActionSuite) Test_Maps_Update_Fail() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token)

  m.Name = "Fake Map"

  // fail on unknown uuid
  req := as.JSON("/v1/maps/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
  req.Headers["Authorization"] = token
  res := req.Put(m)
  as.Equal(http.StatusNotFound, res.Code)

  // fail on bad uuid
  req = as.JSON("/v1/maps/aaaaaaaa")
  req.Headers["Authorization"] = token
  res = req.Put(m)
  as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_Maps_Update_Other_User() {
  // create valid user
	token1 := as.CreateUser("test@test.com", "test")
  token2 := as.CreateUser("fake@fake.com", "fake")

  c := as.CreateCampaign("Test Campagin", "Test Campaign Description", token1)
  q := as.CreateQuest("Test Quest", "Test Quest Description", c.ID, token1)
  m := as.CreateMap("Test Map", "Test Map Description", q.ID, token1)

  m.Name = "Fake Map"

  req := as.JSON("/v1/maps/" + m.ID.String())
  req.Headers["Authorization"] = token2
  res := req.Put(m)
  as.Equal(http.StatusNotFound, res.Code)
}
