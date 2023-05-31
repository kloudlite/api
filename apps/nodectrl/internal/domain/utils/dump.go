package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Upload(bucket *gridfs.Bucket) error {
	fileName := "ram.zip"
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{"metadata tag", "first"}})
	objectID, err := bucket.UploadFromStream(fileName, io.Reader(file),
		uploadOpts)
	if err != nil {
		return err
	}
	fmt.Printf("New file uploaded with ID %s", objectID)
	return nil

}

func getFiles(bucket *gridfs.Bucket) error {
	filter := bson.D{{}}
	cursor, err := bucket.Find(filter)
	if err != nil {
		return err
	}
	type gridfsFile struct {
		Id     string `bson:"_id"`
		Name   string `bson:"filename"`
		Length int64  `bson:"length"`
	}
	var foundFiles []gridfsFile
	if err = cursor.All(context.TODO(), &foundFiles); err != nil {
		panic(err)
	}
	for _, file := range foundFiles {
		fmt.Printf("filename: %s, length: %d, id: %s\n", file.Name, file.Length, file.Id)
	}

	return nil
}

func Delete(bucket *gridfs.Bucket, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if err := bucket.Delete(_id); err != nil {
		return err
	}
	return nil
}

func Download(bucket *gridfs.Bucket, id string) error {

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	fileBuffer := bytes.NewBuffer(nil)
	if _, err := bucket.DownloadToStream(_id, fileBuffer); err != nil {
		panic(err)
	}

	if err := os.WriteFile("out/hello.zip", fileBuffer.Bytes(), os.ModePerm); err != nil {
		return err
	}
	return nil

}

func Update(bucket *gridfs.Bucket, id string) error {
	if err := Delete(bucket, id); err != nil {
		return err
	}
	// bucket.GetChunksCollection().UpdateByID(ctx, id , update , opts ...*options.UpdateOptions)
	return nil
}

func M() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return err
	}

	db := client.Database("myfiles")
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		return err
	}

	// if err := Upload(bucket); err != nil {
	// 	return err
	// }

	if err := getFiles(bucket); err != nil {
		return err
	}

	if err := Download(bucket, "6475cdb55687ddf79ce71c6b"); err != nil {
		return err
	}

	return nil
}

func DumpSample() error {
	return M()
}
