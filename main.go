package main

import (
	"bufio"
	"os"
	"io"
	"log"
	"fmt"
	"os/exec"
	"github.com/urfave/cli"
)

func getFrameworkType(f string) *exec.Cmd {
	switch f {
	case "rails":
		return exec.Command("bundle", "exec", "rails", "s")
	case "yarn":
		return exec.Command("yarn", "run", "start")
	}
	return nil
}

func printOutStdout(stdout io.ReadCloser) {
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	app := cli.NewApp()
	
	app.Name = "myCommand"
	app.Usage = "This application is only for my practice."
	app.Version = "0.0.1"
	
	app.Action = func (context *cli.Context) error {


		switch {
		case context.Bool("framework"):
			var err error

			cmd := getFrameworkType(context.Args().Get(0))
			if err != nil {
				log.Fatal(err)
			}

			stdout, err := cmd.StdoutPipe()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		
			cmd.Start()
		
			printOutStdout(stdout)
		
			cmd.Wait()
		default:
			cli.ShowAppHelp(context)
		}			
		return nil
	}
		
		app.Flags = []cli.Flag{
			cli.BoolFlag {
				Name: "framework, f",
				Usage: "specify framework",
			},
		}
		
		err := app.Run(os.Args)
		
		if err != nil {
			log.Fatal(err)
		}
	}