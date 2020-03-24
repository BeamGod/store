package upload

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	pb "store/protoc"
	"time"
)

type UploadFileClient struct {
	Address string ;

}

func NewUploadFileClient(address string) *UploadFileClient  {
	fmt.Println("new client")
	return &UploadFileClient{Address:address}
}

func (cli *UploadFileClient)UploadFile(filePath string , fileName string)  error{
	fmt.Println("upload file %v" , cli.Address)
	conn , err := grpc.Dial(cli.Address ,  grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after conn")
	defer conn.Close()
	out := pb.NewUploadClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req , err := out.UploadFile(ctx)

	file , err := ioutil.ReadFile(filePath)
	fmt.Println(len(file))
	if err != nil {
		log.Fatal(err)
		return err
	}

	// /Users/bynn/uploadFileTest.txt

	err = req.Send(&pb.UploadFileReq{
		Name: fileName,
		File: file,
	})
	if err != nil {
		log.Fatal("upload file err %v" , err)
		return err
	}

	return nil

}
