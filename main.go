package main

import (
	"time"

	"github.com/mrpiggy97/piggyhttp/cmd"
	"github.com/mrpiggy97/piggyhttp/repository"
	"github.com/rs/zerolog/log"
)

func main() {
	var err error = cmd.RootCmd.Execute()
	if err != nil {
		log.Error().Msg(err.Error())
	}
	time.Sleep(time.Second * 1)
	repository.AppWaiter.Wait()
}
