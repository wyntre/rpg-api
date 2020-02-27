package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/wyntre/rpg_api/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
