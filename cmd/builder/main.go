package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	start := time.Now()
	jsx := `
        import * as React from 'react'
        import * as ReactDOM from 'react-dom'

        ReactDOM.render(
            <h1>Hello, world!</h1>,
            document.getElementById('root')
        );
    `

	result := api.Transform(jsx, api.TransformOptions{
		Loader: api.LoaderJSX,
	})

	finish := time.Now()

	fmt.Printf("%d errors and %d warnings\n",
		len(result.Errors), len(result.Warnings))

	os.Stdout.Write(result.Code)

	log.Printf("Took %f seconds to generate", finish.Sub(start).Seconds())
}