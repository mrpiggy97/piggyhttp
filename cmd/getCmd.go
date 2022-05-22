package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mrpiggy97/piggyhttp/repository"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func getRequest(cmd *cobra.Command, args []string) {

	//set request
	request, requestError := http.NewRequest("GET", *url, nil)
	var client *http.Client = &http.Client{}
	if requestError != nil {
		log.Error().Msg(requestError.Error())
		panic(requestError.Error())
	}
	if len(*authorizationToken) > 0 {
		request.Header.Add("Authorization", *authorizationToken)
	}

	//make request
	response, responseErr := client.Do(request)
	if responseErr != nil {
		log.Error().Msg(responseErr.Error())
	} else {
		decodedResponse, _ := io.ReadAll(response.Body)
		message := fmt.Sprintf(
			"%d,%s,%s",
			response.StatusCode,
			response.Status,
			string(decodedResponse),
		)
		log.Info().Msg(message)
	}
	defer repository.AppWaiter.Done()
	defer response.Body.Close()
}

var getCmd *cobra.Command = &cobra.Command{
	Use:     "get",
	Example: "piggyhttp get --url <target url>",
	Run:     getRequest,
}
