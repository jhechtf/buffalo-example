package grifts

import (
	"fmt"

	"github.com/markbates/grift/grift"
	"github.com/website2/models"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here
		users := []models.User{
			models.User{
				Username: "mark",
				Title:    "Director of Engineering",
			},
			models.User{
				Username: "kenny",
				Title:    "Does a little bit of everything",
			},
			models.User{
				Username: "mark",
				Title:    "Senior Backend Guy",
			},
			models.User{
				Username: "dereck",
				Title:    "Senior Front-End Guy",
			},
			models.User{
				Username: "gordon",
				Title:    "I have no idea what Gordon's actual title is",
			},
			models.User{
				Username: "eric",
				Title:    "I'm beginning to wonder if he still exists tbh",
			},
			models.User{
				Username: "jon",
				Title:    "Senior Full-Stack Guy",
			},
			models.User{
				Username: "jim",
				Title:    "Other Senior Full-Stack Guy",
			},
			models.User{
				Username: "ibrahim",
				Title:    "Developer",
			},
		}
		for _, user := range users {
			err := models.DB.Create(&user)
			if err != nil {
				fmt.Println("ERROR!", err)
			} else {
				fmt.Println("Success!")
			}
		}
		return nil
	})

})
