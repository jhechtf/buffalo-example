package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/website2/models"
)

func SomethingDifferent(c buffalo.Context) error {
	c.Set("name", c.Param("name"))
	return c.Render(201, r.HTML("test.html"))
}

func DevTeam(c buffalo.Context) error {
	users := []models.User{}
	err := models.DB.All(&users)
	fmt.Println("JIM, CAN YOU SEE ME!?", err)
	if err != nil {
		c.Set("error", err)
	} else {
		c.Set("users", users)
	}

	return c.Render(200, r.HTML("devteam.html"))
}
