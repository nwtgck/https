package cmd

import (
	"fmt"
	"github.com/nwtgck/https/util"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var getFlags struct {
	insecure bool
	http3    bool
}

func init() {
	RootCmd.AddCommand(getCmd)

	// Flags
	// NOTE: --insecure, -k is inspired by curl
	getCmd.Flags().BoolVarP(&getFlags.insecure, "insecure", "k", false, "Allow insecure server connections when using SSL/TLS")
	getCmd.Flags().BoolVarP(&getFlags.http3, "http3", "", false, "HTTP3 (experimental)")
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "HTTP GET method",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("URL should be specified")
		}
		// Get server URL and fill https if need
		serverUrlStr, err := util.FillHttps(args[0])
		if err != nil {
			return err
		}
		// Get HTTP client
		httpClient := util.GetHttpClient(getFlags.insecure, getFlags.http3)
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
