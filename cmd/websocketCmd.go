package cmd

import (
	"fmt"
	"os"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func recieveHandler(connection *websocket.Conn) {
	defer connection.Close()
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			var message string = fmt.Sprintf("Error in receiver:%s", err)
			log.Error().Msg(message)
			os.Exit(1)
		}
		log.Printf("Received: %s\n", msg)
	}
}

func connectToWebSocket(cmd *cobra.Command, args []string) {
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
	Long:    "sets a connection with a websocket and that's it",
	Run:     connectToWebSocket,
}
