package grifts

import (
	"github.com/markbates/grift/grift"
	"github.com/wyntre/rpg_api/models"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		tile_category := &models.TileCategory{
			Name: "Castle",
		}
		err := models.DB.Create(tile_category)
		if err != nil {
			return err
		}

		tile_category = &models.TileCategory{
			Name: "Forest",
		}
		err = models.DB.Create(tile_category)
		if err != nil {
			return err
		}

		tile_category = &models.TileCategory{
			Name: "Dungeon",
		}
		err = models.DB.Create(tile_category)
		if err != nil {
			return err
		}

		tile_category = &models.TileCategory{
			Name: "Trap",
		}
		err = models.DB.Create(tile_category)
		if err != nil {
			return err
		}

		tile_category = &models.TileCategory{
			Name: "Town",
		}
		err = models.DB.Create(tile_category)
		if err != nil {
			return err
		}

		tile_category = &models.TileCategory{
			Name: "Road",
		}
		err = models.DB.Create(tile_category)
		if err != nil {
			return err
		}

		return nil
	})

})
