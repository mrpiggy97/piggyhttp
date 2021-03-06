package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mrpiggy97/piggyhttp/repository"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var postData *string = new(string)

func postRequest(cmd *cobra.Command, args []string) {
	repository.AppWaiter.Add(1)
	//set data to send
	var keyValues []string = strings.Split(*postData, ",")
	var keys []string = []string{}
	var values []string = []string{}
	var data map[string]string = make(map[string]string)
	if len(keyValues)%2 != 0 {
		log.Error().Msg("all keys need a value")
	} else {
		//separate keys and values
		for index, member := range keyValues {
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
		//each key
		for i := 0; i < len(keys); i++ {
			data[keys[i]] = values[i]
		}
		jsonData, _ := json.Marshal(data)
		var buffer *bytes.Buffer = bytes.NewBuffer(jsonData)

		//set request
		request, requestError := http.NewRequest("POST", *url, buffer)
		request.Header.Add("Content-type", "application/json")
		var client *http.Client = &http.Client{}
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
}

var postCmd *cobra.Command = &cobra.Command{
	Use:     "post",
	Short:   "makes a post request",
	Long:    "makes a post request and prints response to terminal",
	Example: "piggyhttp post --url <target url> --data <your data>",
	Run:     postRequest,
}

func init() {
	postCmd.Flags().StringVar(postData, "data", "", "will set keys and values for body length, its length must be a divisible by 2")
	postCmd.MarkFlagRequired("data")
}
