package main


import (
	"algorithms/App"
	"github.com/urfave/cli"
	"os"
)

func main(){

	cliApp := App.CliApp{
		App: cli.NewApp(),
	}

	if err := cliApp.RunApp(os.Args); err != nil {
		os.Exit(1)
	}

}