package cmd

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

func init() {
	RootCmd.AddCommand(postCmd)
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
		// TODO: Hard code
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{ InsecureSkipVerify: true },
		}
		httpClient := http.Client{Transport: tr}
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
