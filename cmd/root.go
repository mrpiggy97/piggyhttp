package cmd

import "github.com/spf13/cobra"

var url *string = new(string)
var authorizationToken *string = new(string)

var RootCmd *cobra.Command = &cobra.Command{
	Use:     "piggyhttp",
	Version: "v1.0.2",
	Short:   "makes http requests",
	Long:    "will make an http request based on url given and print the result",
}

func init() {
	RootCmd.AddCommand(getCmd, postCmd, putCmd, deleteCmd)
	RootCmd.PersistentFlags().StringVar(url, "url", "default", "will set target url")
	RootCmd.PersistentFlags().StringVar(authorizationToken, "withAuthorizationToken", "", "sets token for authorization header")
}
