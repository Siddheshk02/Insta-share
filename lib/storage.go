package lib

import (
	"context"
	"fmt"
	"io"
	"log"
	"path/filepath"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

func UploadFile(c *fiber.Ctx) error {
	fmt.Println("done")
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	filename := file.Filename

	ctx := context.Background()
	conf := &firebase.Config{
		StorageBucket: "instashare-fe6e6.appspot.com",
	}

	opt := option.WithCredentialsFile("instashare-fe6e6-firebase-adminsdk-jl7ew-c1732f98fb.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return err
	}

	client, err := app.Storage(ctx)
	if err != nil {
		return err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return err
	}
	// file1 := strings.Itoa(file)

	// data, err := ioutil.ReadFile(file)
	// if err != nil {
	// 	log.Fatalf("Failed to read file: %v", err)
	// }
	obj := bucket.Object("Files" + "/" + filepath.Base(filename))
	wc := obj.NewWriter(ctx)
	// if _, err = wc.Write(data); err != nil {
	// 	log.Fatalf("Failed to write file to Firebase Cloud Storage: %v", err)
	// }
	// if err := wc.Close(); err != nil {
	// 	log.Fatalf("Failed to close Firebase Cloud Storage writer: %v", err)
	// }
	fileReader, err := file.Open()
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer fileReader.Close()

	if _, err := io.Copy(wc, fileReader); err != nil {
		log.Fatalf("Failed to write file to Firebase Cloud Storage: %v", err)
	}
	fmt.Println("Done")

	return nil
}
