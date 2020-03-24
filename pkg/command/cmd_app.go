package command

import "C"
import (
	"fmt"
	"github.com/urfave/cli/v2"
	"store/clients/upload"
	server2 "store/services/store/server"
)

func NewUrfaveApp(types string) *cli.App {
	if types == "uploadCli"{

	}
	switch types {
	case "uploadCli":
		return newUploadUrfaveApp()
	case "uploadServer":
		return NewUploadFileServerUrfave()
	default:
		return newUploadUrfaveApp()
	}
}

func newUploadUrfaveApp() *cli.App {

	config := &uploadConfig{}

	return &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "address",
				Usage: "language for the greeting",
			},

			&cli.StringFlag{
				Name: "path",
				Usage: "path of file to upload ",
			},
			&cli.StringFlag{
				Name: "name",
				Usage: "fileName for new project",
			},
			&cli.StringFlag{
				Name: "tag",
				Usage: "tag for project in store",
			},
		},
		Action: func(c *cli.Context) error {
			config.address = c.String("address")
			config.path = c.String("path")
			config.fileName = c.String("name")
			config.tag = c.String("tag")
			fmt.Println(config)
			cli := upload.NewUploadFileClient(config.address)
			cli.UploadFile(config.path , config.fileName)
			return nil
		},
	}
}

func NewUploadFileServerUrfave()  *cli.App{
	return &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "port",
				Usage: "port",
			},

			&cli.StringFlag{
				Name: "path",
				Usage: "path of file to upload ",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println(c.String("port"))
			fmt.Println(c.String("path"))
			app := server2.NewUploadFileApp(c.String("port") , c.String("path"))
			app.Run()
			return nil
		},
	}
}

type uploadConfig struct {
	address string
	fileName string
	path string
	tag string
}