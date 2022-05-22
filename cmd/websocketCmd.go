package cmd

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/mrpiggy97/piggyhttp/repository"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func recieveHandler(connection *websocket.Conn) {
	defer connection.Close()
	defer repository.AppWaiter.Done()
	log.Info().Msg("waiting for socket messages")
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			var message string = fmt.Sprintf("Error in receiver:%s", err)
			log.Error().Msg(message)
			break
		}
		log.Info().Msg(string(msg))
	}
}

func connectToWebSocket(cmd *cobra.Command, args []string) {
	repository.AppWaiter.Add(1)
	socketConn, _, socketErr := websocket.DefaultDialer.Dial(*url, nil)
	if socketErr != nil {
		log.Error().Msg(socketErr.Error())
	} else {
		var message string = fmt.Sprintf("successfuly connected to %s with a websocket", *url)
		log.Info().Msg(message)
		go recieveHandler(socketConn)
	}
}

var webSocketCmd *cobra.Command = &cobra.Command{
	Use:     "websocket",
	Example: "piggyhttp websocket --url <url here>",
	Short:   "sets a connection with a websocket",
	Long:    "sets a connection with a websocket and listenes for message from that connection",
	Run:     connectToWebSocket,
}
