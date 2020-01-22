package App


import (
	"github.com/urfave/cli"
	linkedList "algorithms/LinkedList"
)

type CliApp struct {
	App *cli.App
}


func (c *CliApp) RunApp(args []string)error {

	c.App = &cli.App {
		Name: "algorithms-cli",
		Commands: []cli.Command{
			{
				
				Name: "linked_list",
				Aliases: []string{"ll"},
				Description: "run linked list",
				Action: linkedList.Run,
			},
		},
	}

	return c.App.Run(args)


}
