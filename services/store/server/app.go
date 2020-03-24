package server

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	pb "store/protoc"
	"io/ioutil"
)

type UploadFileApp struct {
	
	Port string;
	SavePath string;
	pb.UnimplementedUploadServer
}

func (app *UploadFileApp)UploadFile(stream pb.Upload_UploadFileServer) error {
	log.Println("into func")
	req , err := stream.Recv()
	log.Println(req)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	//fileName := req.Name
	fmt.Println(app.SavePath + "ttt.txt")
	fmt.Println(app)
	//log.Fatal("file name is %s" , fileName)
	file := req.File
	err = ioutil.WriteFile( app.SavePath + req.Name , file , 0644)
	fmt.Println("write over")
	if err != nil {
		log.Println(err)
	}

	return nil
}

func NewUploadFileApp(port string , path string) *UploadFileApp {
	app := &UploadFileApp{
		Port:                      port,
		SavePath:                  path,
	}
	fmt.Println("app %v" , app)
		return app
}





func (app *UploadFileApp)Run() error{

	lis, err := net.Listen("tcp" , ":" + app.Port)

	if err != nil {
		log.Fatal(err)
		return err
	}

	
	s := grpc.NewServer()
	pb.RegisterUploadServer(s , app)
	s.Serve(lis)

	return nil

}
