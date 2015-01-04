package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/shunsugai/android/patternlock"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "ptrn"
	app.Usage = "Decrypt gesture.key"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			cli.ShowAppHelp(c)
			return
		}
		unlock(c)
	}
	app.Run(os.Args)
}

func unlock(c *cli.Context) {
	hash, err := ioutil.ReadFile(c.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hash Key : %x\n", hash)
	p := patternlock.Calc(hash)
	patternlock.PrintResult(p)
}
