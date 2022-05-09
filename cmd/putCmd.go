package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var putData *string = new(string)

func putRequest(cmd *cobra.Command, args []string) {
	//set data to send
	var keysAndValues []string = strings.Split(*putData, ",")
	var data map[string]string = make(map[string]string)
	if len(keysAndValues)%2 != 0 {
		var err error = errors.New("each key of data must have a value")
		log.Error().Msg(err.Error())
	} else {
		var keys []string = []string{}
		var values []string = []string{}

		for index, member := range keysAndValues {
			if index == 0 {
				keys = append(keys, member)
			} else {
				if index%2 == 0 {
					keys = append(keys, member)
				} else {
					values = append(values, member)
				}
			}
		}

		for i := 0; i < len(keys); i++ {
			data[keys[i]] = values[i]
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
			message := fmt.Sprintf(
				"%d,%s,%s",
				response.StatusCode,
				response.Status,
				string(decodedResponse),
			)
			log.Info().Msg(message)
		}
		defer response.Body.Close()
	}
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
