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

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "enter",
	Long:  "enter",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: impl
		// TODO: Hard code
		reader := os.Stdin
		// TODO: Hard code
		serverUrlStr := "https://trans.cf"
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
		fileIdBytes, _ := ioutil.ReadAll(resp.Body)
		fileId := strings.TrimRight(string(fileIdBytes), "\n")
		fmt.Println(fileId)

		return nil
	},
}

func init() {
	cobra.OnInitialize()
}
