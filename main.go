package main

import (
	"github.com/mrpiggy97/piggyhttp/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	var err error = cmd.RootCmd.Execute()
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
