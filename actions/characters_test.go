package actions

import (
  "net/http"

  "github.com/wyntre/rpg_api/models"
)

type CharactersListResponse struct {
  Characters models.Characters `json:characters`
}

type BadCharacter struct {
  Name string `json:name`
  Description string `json:description`
  ExtraData string `json:extradata`
}

func (as *ActionSuite) Test_Characters_Create() {
  // create valid user
  token := as.CreateUser("test@test.com", "test")

  // create character data
  character := &models.Character{
    Name: "test",
    Description: "test",
  }

  // request character creation
	req := as.JSON("/v1/characters/new")
	req.Headers["Authorization"] = token
	res := req.Post(character)
	as.Equal(http.StatusCreated, res.Code)

  // validate character id
  data := map[string]string{}
  res.Bind(data)
  as.NotNil(data["id"])
}

func (as *ActionSuite) Test_Characters_Create_Fail() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  // create invalid character data
  character := &models.Character{
    Name: "test",
  }

  // request character creation
	req := as.JSON("/v1/characters/new")
	req.Headers["Authorization"] = token
	res := req.Post(character)
	as.Equal(http.StatusConflict, res.Code)
}

func (as *ActionSuite) Test_Characters_Create_Extra() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  // create invalid character data
  character := &BadCharacter{
    Name: "Test",
    Description: "Test description",
    ExtraData: "Extra Data",
  }

  // request character creation
	req := as.JSON("/v1/characters/new")
	req.Headers["Authorization"] = token
	res := req.Post(character)
	as.Equal(http.StatusCreated, res.Code)
}

func (as *ActionSuite) Test_Characters_Destroy() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  // create character data
  character := &models.Character{
    Name: "test",
    Description: "test",
  }

  // request character creation
	req := as.JSON("/v1/characters/new")
	req.Headers["Authorization"] = token
	res := req.Post(character)
	as.Equal(http.StatusCreated, res.Code)

  // validate character id
  res.Bind(character)
  as.NotNil(character.ID)

  // request character deletion
  req = as.JSON("/v1/characters/" + character.ID.String())
  req.Headers["Authorization"] = token
  res = req.Delete()
  as.Equal(http.StatusOK, res.Code)

  // request character deletion again
  req = as.JSON("/v1/characters/" + character.ID.String())
  req.Headers["Authorization"] = token
  res = req.Delete()
  as.Equal(http.StatusNotFound, res.Code)

  // request character deletion with bad uuid
  req = as.JSON("/v1/characters/" + "abce")
  req.Headers["Authorization"] = token
  res = req.Delete()
  as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_Characters_List() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  // create character data
  character1 := &models.Character{
    Name: "test",
    Description: "test",
  }

  character2 := &models.Character{
    Name: "test",
    Description: "test",
  }

  // request character creation
	req := as.JSON("/v1/characters/new")
	req.Headers["Authorization"] = token
	res := req.Post(character1)
	as.Equal(http.StatusCreated, res.Code)
  res = req.Post(character2)
	as.Equal(http.StatusCreated, res.Code)

  characters := &CharactersListResponse{}
  req = as.JSON("/v1/characters")
  req.Headers["Authorization"] = token
  res = req.Get()
  res.Bind(characters)

  as.Equal(2, len(characters.Characters))
}

func (as *ActionSuite) Test_Characters_Show() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  // create character data
  character := &models.Character{
    Name: "test",
    Description: "test",
  }

  req := as.JSON("/v1/characters/new")
  req.Headers["Authorization"] = token
	res := req.Post(character)
	as.Equal(http.StatusCreated, res.Code)

  res.Bind(character)

  // request character data
  req = as.JSON("/v1/characters/" + character.ID.String())
  req.Headers["Authorization"] = token
  res = req.Get()
  as.Equal(http.StatusOK, res.Code)

  // test for character id
  test_character := &models.Character{}
  res.Bind(test_character)
  as.NotNil(test_character.ID)
}

func (as *ActionSuite) Test_Characters_Show_Fail() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  // create character data
  character := &models.Character{
    Name: "test",
    Description: "test",
  }

  req := as.JSON("/v1/characters/new")
  req.Headers["Authorization"] = token
	res := req.Post(character)
	as.Equal(http.StatusCreated, res.Code)

  // non-existant character id
  req = as.JSON("/v1/characters/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
  req.Headers["Authorization"] = token
  res = req.Get()
  as.Equal(http.StatusNotFound, res.Code)

  // bad character id
  req = as.JSON("/v1/characters/" + "abce")
  req.Headers["Authorization"] = token
  res = req.Get()
  as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_Characters_Show_Other_User() {
  // create valid user
	token1 := as.CreateUser("test@test.com", "test")
  token2 := as.CreateUser("fake@fake.com", "fake")

  // create character data
  character := &models.Character{
    Name: "test",
    Description: "test",
  }

  req := as.JSON("/v1/characters/new")
  req.Headers["Authorization"] = token1
	res := req.Post(character)
	as.Equal(http.StatusCreated, res.Code)

  res.Bind(character)
  as.NotNil(character.ID)

  req = as.JSON("/v1/characters/" + character.ID.String())
  req.Headers["Authorization"] = token2
  res = req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Characters_Update() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  // create character data
  character := &models.Character{
    Name: "test",
    Description: "test",
  }

  req := as.JSON("/v1/characters/new")
  req.Headers["Authorization"] = token
	res := req.Post(character)
	as.Equal(http.StatusCreated, res.Code)

  res.Bind(character)
  as.Equal("test", character.Name)

  character.Name = "fake"

  req = as.JSON("/v1/characters/" + character.ID.String())
  req.Headers["Authorization"] = token
  res = req.Put(character)
  as.Equal(http.StatusOK, res.Code)

  test_character := &models.Character{}
  res.Bind(test_character)
  as.Equal("fake", test_character.Name)
}

func (as *ActionSuite) Test_Characters_Update_Fail() {
  // create valid user
	token := as.CreateUser("test@test.com", "test")

  // create character data
  character := &models.Character{
    Name: "test",
    Description: "test",
  }

  req := as.JSON("/v1/characters/new")
  req.Headers["Authorization"] = token
	res := req.Post(character)
	as.Equal(http.StatusCreated, res.Code)

  res.Bind(character)
  as.Equal("test", character.Name)

  character.Name = "fake"

  // bad character id
  req = as.JSON("/v1/characters/abcd")
  req.Headers["Authorization"] = token
  res = req.Put(character)
  as.Equal(http.StatusInternalServerError, res.Code)

  // non-existant character id
  req = as.JSON("/v1/characters/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
  req.Headers["Authorization"] = token
  res = req.Put(character)
  as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Characters_Update_Other_User() {
  // create valid user
	token1 := as.CreateUser("test@test.com", "test")
  token2 := as.CreateUser("fake@fake.com", "fake")

  // create character data
  character := &models.Character{
    Name: "test",
    Description: "test",
  }

  req := as.JSON("/v1/characters/new")
  req.Headers["Authorization"] = token1
	res := req.Post(character)
	as.Equal(http.StatusCreated, res.Code)

  res.Bind(character)
  as.NotNil(character.ID)

  character.Name = "fake"

  req = as.JSON("/v1/characters/" + character.ID.String())
  req.Headers["Authorization"] = token2
  res = req.Put(character)
	as.Equal(http.StatusNotFound, res.Code)
}
