package cmd

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func deleteRequest(cmd *cobra.Command, args []string) {

	//set request
	var client *http.Client = &http.Client{}
	request, requestError := http.NewRequest("DELETE", *url, nil)
	if requestError != nil {
		log.Error().Msg(requestError.Error())
		panic(requestError.Error())
	}
	if len(*authorizationToken) > 0 {
		request.Header.Add("Authorization", *authorizationToken)
	}

	//make request
	response, responseError := client.Do(request)
	if responseError != nil {
		log.Error().Msg(responseError.Error())
	} else {
		decodedResponse, _ := io.ReadAll(response.Body)
		log.Info().Msg(string(decodedResponse))
	}
	defer response.Body.Close()
}

var deleteCmd *cobra.Command = &cobra.Command{
	Use:     "delete",
	Short:   "makes a delete request",
	Long:    "makes a delete http request and prints response to terminal",
	Example: "piggyhttp delete --url <target url>",
	Run:     deleteRequest,
}
