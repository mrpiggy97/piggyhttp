package cmd

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func getRequest(cmd *cobra.Command, args []string) {
	request, requestError := http.NewRequest("GET", *url, nil)
	var client *http.Client = &http.Client{}
	if requestError != nil {
		log.Error().Msg(requestError.Error())
	}
	if len(*authorizationToken) > 0 {
		request.Header.Add("Authorization", *authorizationToken)
	}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		log.Error().Msg(responseErr.Error())
	} else {
		decodedResponse, _ := io.ReadAll(response.Body)
		log.Info().Msg(string(decodedResponse))
	}
	defer response.Body.Close()
}

var getCmd *cobra.Command = &cobra.Command{
	Use:     "get",
	Example: "piggyhttp get --url <target url>",
	Run:     getRequest,
}
