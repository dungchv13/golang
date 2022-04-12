package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"os"
)

func main() {
	endpoint := "171.244.133.228:30292"
	accessKeyID := "iot"
	secretAccessKey := "iot@2022"

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	options := minio.GetObjectOptions{}
	err = options.SetRange(0, 10000)
	object, err := minioClient.GetObject(context.Background(), "test1", "cat.jpg", options)
	if err != nil {
		fmt.Println(err)
		return
	}
	localFile, err := os.Create("/home/dung/GolandProjects/awesomeProject/cat.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err = io.Copy(localFile, object); err != nil {
		fmt.Println(err)
		return
	}
}
