package destroyer

import (
	"log"

	"github.com/nubunto/vise/persistence"
	"github.com/robfig/cron"
)

func Scan() {
	c := cron.New()
	c.AddFunc("@daily", func() {
		err := persistence.CheckExpiration()
		if err != nil {
			log.Println("Failed to run today's daily scanner:", err)
		}
	})
	c.Start()
}
