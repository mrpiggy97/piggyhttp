package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var postData *string = new(string)

func postRequest(cmd *cobra.Command, args []string) {
	var data map[string]string = map[string]string{
		"data": *postData,
	}
	jsonData, _ := json.Marshal(data)
	var buffer *bytes.Buffer = bytes.NewBuffer(jsonData)
	request, requestError := http.NewRequest("POST", *url, buffer)
	var client *http.Client = &http.Client{}
	if requestError != nil {
		log.Error().Msg(requestError.Error())
	}
	if len(*authorizationToken) > 0 {
		request.Header.Add("Authorization", *authorizationToken)
	}
	response, responseError := client.Do(request)
	if responseError != nil {
		log.Error().Msg(responseError.Error())
	} else {
		decodedResponse, _ := io.ReadAll(response.Body)
		log.Info().Msg(string(decodedResponse))
	}
}

var postCmd *cobra.Command = &cobra.Command{
	Use:     "post",
	Short:   "makes a post request",
	Long:    "makes a post request and prints response to terminal",
	Example: "piggyhttp post --url <target url> --data <your data>",
	Run:     postRequest,
}