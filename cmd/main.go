package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"sort"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isophtalic/License/internal/configs"
	"github.com/isophtalic/License/internal/initdata"
	"github.com/isophtalic/License/internal/persistence"
	"github.com/isophtalic/License/internal/routes"
	"github.com/urfave/cli/v2"
)

var (
	ctx    gin.Context
	config *configs.Configure
	err    error
	mode   string
	serve  *gin.Engine
)

func InfoCommand() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		fmt.Printf("*** %s - Version %s ***\n", c.App.Name, c.App.Version)
		fmt.Printf("*** %s ***\n", c.App.Copyright)
		return nil
	}
}
func RunServerCommand() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		config, err = configs.GetConfig()
		handleError()
		migrateDB()
		if config.Mode == "release" {
			fmt.Printf("Service is running on http://127.0.0.1%s/api/v1\n", config.ServerPort)
		}
		serve := routes.NewAPIv1(config, config.Mode)

		serve.Run(config.ServerPort)
		return nil
	}
}
func MigrateDBCommand() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		config, err = configs.GetConfig()
		if err != nil {
			panic(err)
		}

		persistence.ConnectDatabase(config)
		persistence.MigrateDatabase()
		initdata.InitAccount()
		log.Println("Migrate successfully")
		return nil
	}
}
func main() {
	app := cli.NewApp()
	app.Name = "CyRadar License Management Backend API"
	app.Usage = "CyRadar License Management Backend"
	app.Copyright = "Copyright © 2023 CyRadar. All Rights Reserved."
	currentTime := time.Now()
	app.Version = "1.1.1-" + currentTime.Format("2006-01-02 17:06:06")
	// app.Compiled = currentTime.Format("2017-09-07 17:06:06")
	// flag.Parse()
	// flag.StringVar(&mode, "mode", "debug", "Mode : debug | release")
	app.Commands = []*cli.Command{
		{
			Name:   "info",
			Usage:  "print application name and version",
			Action: InfoCommand(),
		},
		{
			Name:   "runserver",
			Usage:  "Run server the web server",
			Action: RunServerCommand(),
		},
		{
			Name:   "migrate",
			Usage:  "Run server the web server",
			Action: MigrateDBCommand(),
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("err:", err)
	}
}

func handleError() {
	errSignal := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(errSignal, syscall.SIGTERM)

	go func(c *gin.Context) {
		err := c.Errors.Last()
		errChan <- err
	}(&ctx)

	select {
	case sign := <-errSignal:
		log.Println("Shutting down : ", sign)
		return
	case err := <-errChan:
		if err == nil || (reflect.ValueOf(err).Kind() == reflect.Ptr && reflect.ValueOf(err).IsNil()) {
			return
		}
		log.Println("ERROR: ", err)

		return
	}
}

func migrateDB() {
	config, err = configs.GetConfig()
	if err != nil {
		panic(err)
	}

	persistence.ConnectDatabase(config)
	persistence.MigrateDatabase()
	initdata.InitAccount()

	flag.StringVar(&mode, "mode", "debug", "Mode : debug | release")
}
