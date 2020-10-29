package cmd

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use: "post",
	Run: func(cmd *cobra.Command, args []string) {
		Post()
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
}

// Post - gorequest post test
func Post() {
	// gorequest init
	req := gorequest.New()

	// API router
	url := "https://jsonplaceholder.typicode.com/posts"

	// Post
	req.Post(url)

	// Post data
	d := &TestBody{
		Title:  "YauTz",
		Body:   "YauTz Body",
		UserID: 123456,
	}

	// send 
	req.Send(d)

	// response
	v := new(TestBody)
	resp, _, errs := req.EndStruct(&v)

	// error check
	for _, err := range errs {
		if err != nil {
			logrus.Error(err)
			break
		}
	}

	fmt.Println(resp)
	fmt.Println()
	fmt.Println("res: ", v)
}
