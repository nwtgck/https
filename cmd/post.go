package cmd

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

var postInsecure bool

func init() {
	RootCmd.AddCommand(postCmd)

	// Flags
	// NOTE: --insecure, -k is inspired by curl
	postCmd.Flags().BoolVarP(&postInsecure, "insecure", "k", false, "Allow insecure server connections when using SS")
}

// Generate HTTP client
func getHttpClient() *http.Client {
	// Set insecure or not
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{ InsecureSkipVerify: postInsecure },
	}
	return &http.Client{Transport: tr}

}

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "HTTP POST method",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("URL should be specified")
		}
		// Get server URL
		// TODO: improve usability to fill "https://" automatically
		serverUrlStr := args[0]
		// Read from stdin
		// NOTE: reader can be change in the future by options
		reader := os.Stdin
		// TODO: Hard code
		contentType := "application/octet-stream"
		// Get HTTP client
		httpClient := getHttpClient()
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
