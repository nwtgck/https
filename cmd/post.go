package cmd

import (
	"fmt"
	"github.com/nwtgck/https/util"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var postFlag struct {
	insecure bool
	http3    bool
}

func init() {
	RootCmd.AddCommand(postCmd)

	// Flags
	// NOTE: --insecure, -k is inspired by curl
	postCmd.Flags().BoolVarP(&postFlag.insecure, "insecure", "k", false, "Allow insecure server connections when using SSL/TLS")
	postCmd.Flags().BoolVarP(&postFlag.http3, "http3", "", false, "HTTP3 (experimental)")
}

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "HTTP POST method",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("URL should be specified")
		}
		// Get server URL and fill https if need
		serverUrlStr, err := util.FillHttps(args[0])
		if err != nil {
			return err
		}
		// Read from stdin
		// NOTE: reader can be change in the future by options
		reader := os.Stdin
		// TODO: Hard code
		contentType := "application/octet-stream"
		// Get HTTP client
		httpClient := util.GetHttpClient(postFlag.insecure, postFlag.http3)
		// Post
		resp, err := httpClient.Post(serverUrlStr, contentType, reader)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Output body as a stream
		io.Copy(os.Stdout, resp.Body)

		return nil
	},
}
