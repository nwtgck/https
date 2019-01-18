package cmd

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

		// TODO: Hard code
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

		// TODO: Output body as a stream
		fileIdBytes, _ := ioutil.ReadAll(resp.Body)
		fileId := strings.TrimRight(string(fileIdBytes), "\n")
		fmt.Println(fileId)

		return nil
	},
}
