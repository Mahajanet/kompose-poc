package main

import (
	"fmt"

	"github.com/kernelpanic77/poc/pkg/convert"
)

func main() {
	client, err := convert.NewClient(convert.WithErrorOnWarning())
	if err != nil {
		fmt.Print(err)
		return
	}
	_, err = client.Convert(convert.ConvertOptions{
		OutFile: "",
		InputFiles: []string{
			"docker-compose.yaml",
		},
	})
	if err != nil {
		fmt.Print(err)
		return
	}

}
