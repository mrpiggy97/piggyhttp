package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var putData *string = new(string)

func putRequest(cmd *cobra.Command, args []string) {
	//set data to send
	var data map[string]string = map[string]string{
		"data": *putData,
	}

	jsonData, _ := json.Marshal(data)
	var buffer *bytes.Buffer = bytes.NewBuffer(jsonData)

	//set request
	request, requestError := http.NewRequest("PUT", *url, buffer)
	var client *http.Client = &http.Client{}
	if requestError != nil {
		log.Err(requestError).Msg(requestError.Error())
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

var putCmd *cobra.Command = &cobra.Command{
	Use:     "put",
	Short:   "makes a put request",
	Long:    "makes a put request and then prints response to terminal",
	Example: "piggyhttp put --url <target url> --data <data to send>",
	Run:     putRequest,
}

func init() {
	putCmd.Flags().StringVar(putData, "data", "", "sets data to send")
	putCmd.MarkFlagRequired("data")
}
