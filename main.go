package main

import (
	"log"
	"os"
	"store/clients"
	"store/pkg/command"
	"store/server"
	//"store/pkg/test"
)


func main()  {
	types := "uploadServer"
	app := command.NewUrfaveApp(types) // uploadServer  uploadCli
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	//servers()
	//c := test.NewCmd()
	//c.Run()
}

func clien()  {
	cli := clients.NewUploadFileClient("127.0.0.1:8081")
	cli.UploadFile("/Users/bynn/uploadFileTest.txt" , "uploadFileTest.txt")
}

func servers()  {
	app := server.NewUploadFileApp("8888" , "/Users/bynn/")
	app.Run()
}