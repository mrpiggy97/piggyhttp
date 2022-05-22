package cmd

import "github.com/spf13/cobra"

var url *string = new(string)
var authorizationToken *string = new(string)

var RootCmd *cobra.Command = &cobra.Command{
	Use:     "piggyhttp",
	Version: "v1.0.5",
	Short:   "makes http requests and can make websocket connections",
	Long:    "will make an http request based on url given and print the result, you can also make a websocket request",
}

func init() {
	RootCmd.AddCommand(getCmd, postCmd, putCmd, deleteCmd, webSocketCmd)
	RootCmd.PersistentFlags().StringVar(url, "url", "default", "will set target url")
	RootCmd.PersistentFlags().StringVar(authorizationToken, "authorization", "", "sets jwt token for authorization header")
}
