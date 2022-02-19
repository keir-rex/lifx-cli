package main

import (
	"github.com/urfave/cli"
	"github.com/keir-rex/lifx-cli/cmd/lifx/client"
	"github.com/keir-rex/lifx-cli/cmd/macos/screens"
	"os"
	"fmt"
)

func main() {
	var display, light, displayState, lightState string
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Name = "Lifx"
	app.Usage = "lifx [command] [...params]"
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "debug, d", Usage: "Enable debugging"},
	}
	
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true

	app.Commands = []cli.Command{
		{
			Name:  "config",
			Usage: "Gets the configuration",
			Action: func(c *cli.Context) {
				client.Config()
			},
		},
		{
			Name:  "select",
			Usage: "Select a light",
			Action: func(c *cli.Context) {
				var selector string
				if c.NArg() > 0 {
					selector = c.Args()[0]
				}
				debug := c.GlobalBool("debug")
				client.SelectLight(debug, selector)
			},
		},
		{
			Name:  "list",
			Usage: "List all bulbs",
			Action: func(c *cli.Context) {
				debug := c.GlobalBool("debug")
				client.List(debug)
			},
		},
		{
			Name:  "toggle",
			Usage: "Toggle selected/all bulbs",
			Action: func(c *cli.Context) {
				debug := c.GlobalBool("debug")
				client.Toggle(debug)
			},
		},
		{
			Name:  "on",
			Usage: "Turn light(s) on",
			Action: func(c *cli.Context) {
				debug := c.GlobalBool("debug")
				client.On(debug)
			},
		},
		{
			Name:  "off",
			Usage: "Turn light(s) off",
			Action: func(c *cli.Context) {
				debug := c.GlobalBool("debug")
				client.Off(debug)
			},
		},
		{
			Name:  "brightness",
			Usage: "Change the brightness of a light",
			Action: func(c *cli.Context) {
				var brightness string
				if c.NArg() > 0 {
					brightness = c.Args()[0]
					debug := c.GlobalBool("debug")
					client.Brightness(debug, brightness)
				} else {
					println("Enter brighness level 0.0 - 1.0")
				}
			},
		},
		{
			Name:  "color",
			Usage: "Change the color of a light",
			Action: func(c *cli.Context) {
				var color string
				if c.NArg() > 0 {
					color = c.Args()[0]
					debug := c.GlobalBool("debug")
					client.Color(debug, color)
				} else {
					println("Enter color string")
				}
			},
		},
		{
			Name:        "displays",
			Usage:       "Commands relating to macos displays and triggering lifx bulbs",
			Subcommands: []cli.Command{
			  {
				Name:  "list",
				Usage: "Print a list of macos display names",
				Action: func(c *cli.Context) error {
					screenList := screens.List()
					for _, screen := range screenList {
						println(screen)
					}
					return nil
				},
			  },
			  {
				Name:  "active",
				Usage: "Print name of active display",
				Action: func(c *cli.Context) error {
					println(screens.Active())
					return nil
				},
			  },
			  {
				Name:  "if",
				Usage: "If display is in desired state set desired light state",
				Flags: []cli.Flag {
					cli.StringFlag{
					  Name: "display, d",
					  Value: "",
					  Usage: "--display `NAME` (required)",
					  Destination: &display,
					},
					cli.StringFlag{
						Name: "light, l",
						Value: "",
						Usage: "--light `UID` (defaults to selected light)",
						Destination: &light,
					  },
					  cli.StringFlag{
						Name: "display-state",
						Value: "active",
						Usage: "--display-state	`STATE`",
						Destination: &displayState,
					  },
					  cli.StringFlag{
						Name: "light-state",
						Value: "on",
						Usage: "--light-state	`STATE`",
						Destination: &lightState,
					  },
				},
				Action: func(c *cli.Context) error {
					isActive := screens.IsActive(display)
					if display == "" {
						fmt.Println("Provide display name")
					}
					if light == "" {
						fmt.Println("Provide light id")
					}
					if displayState == "active" && isActive {
						fmt.Println("display active")
						client.Set(light, lightState)
					}
					if displayState == "inactive" && !isActive {
						fmt.Println("display inactive")
						client.Set(light, lightState)
					}
					return nil
				},
			  },
			},
		},
	}
	app.Run(os.Args)
}
