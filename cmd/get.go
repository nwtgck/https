package cmd

import (
	"fmt"
	"github.com/nwtgck/https/util"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var getInsecure bool

func init() {
	RootCmd.AddCommand(getCmd)

	// Flags
	// NOTE: --insecure, -k is inspired by curl
	getCmd.Flags().BoolVarP(&getInsecure, "insecure", "k", false, "Allow insecure server connections when using SS")
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "HTTP GET method",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("URL should be specified")
		}
		// Get server URL
		// TODO: improve usability to fill "https://" automatically
		serverUrlStr := args[0]
		// Get HTTP client
		httpClient := util.GetHttpClient(getInsecure)
		// Post
		resp, err := httpClient.Get(serverUrlStr)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Output body as a stream
		io.Copy(os.Stdout, resp.Body)

		return nil
	},
}
