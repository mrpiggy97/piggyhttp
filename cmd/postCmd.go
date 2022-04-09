package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var postData *string = new(string)

func postRequest(cmd *cobra.Command, args []string) {

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
		fmt.Println(data)
		jsonData, _ := json.Marshal(data)
		fmt.Println(string(jsonData))
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
		if response.StatusCode != http.StatusAccepted {
			log.Error().Msg(response.Status)
		} else if responseError != nil {
			log.Error().Msg(responseError.Error())
		} else {
			decodedResponse, _ := io.ReadAll(response.Body)
			log.Info().Msg(string(decodedResponse))
		}
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
