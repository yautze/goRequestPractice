package cmd

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Do - 
var Do string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {
		if Do != "1" {
			GetAll()
		}else{
			GetOne()
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&Do, "do", "d", "1", "get one or all")
}

// TestBody -
type TestBody struct {
	UserID int64  `json:"userId"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// GetOne - gorequest get one test
func GetOne() {
	// gorequest init
	req := gorequest.New()

	// API router
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// GET
	req.Get(url)

	// request set
	req.Header.Add("Content-Type", "application/json")
	req.Type("json")

	// Unmarshal body 
	v := new(TestBody)
	_, _, errs := req.EndStruct(v)

	// error check
	for _, err := range errs {
		if err != nil {
			logrus.Error(err)
			break
		}
	}

	fmt.Println(v)
}

// GetAll - gorequest get lsit test
func GetAll() {
	// gorequest init
	req := gorequest.New()

	// API router
	url := "https://jsonplaceholder.typicode.com/posts"

	// GET
	req.Get(url)

	// request set
	req.Header.Add("Content-Type", "application/json")
	req.Type("json")

	// Unmarshal body 
	v := make([]*TestBody, 0)
	_, _, errs := req.EndStruct(&v)

	// error check
	for _, err := range errs {
		if err != nil {
			logrus.Error(err)
			break
		}
	}

	for _, val := range v {
		fmt.Println(val)
		fmt.Println()
	}
}
