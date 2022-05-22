package main

import (
	"github.com/mrpiggy97/piggyhttp/cmd"
	"github.com/mrpiggy97/piggyhttp/repository"
	"github.com/rs/zerolog/log"
)

func main() {
	repository.AppWaiter.Add(1)
	var err error = cmd.RootCmd.Execute()
	if err != nil {
		log.Error().Msg(err.Error())
	}
	repository.AppWaiter.Wait()
}
