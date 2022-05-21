package cmd

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func connectToWebSocket(cmd *cobra.Command, args []string) {
	_, _, socketErr := websocket.DefaultDialer.Dial(*url, nil)
	if socketErr != nil {
		log.Error().Msg(socketErr.Error())
	} else {
		var message string = fmt.Sprintf("successfuly connected to %s with a websocket", *url)
		log.Info().Msg(message)
	}
}

var webSocketCmd *cobra.Command = &cobra.Command{
	Use:     "websocket",
	Example: "piggyhttp websocket --url <url here>",
	Short:   "sets a connection with a websocket",
	Long:    "sets a connection with a websocket and that's it",
	Run:     connectToWebSocket,
}
