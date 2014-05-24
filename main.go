package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var (
		providerName = flag.String("provider", "http", "Provider name - Available http")
		maxWidth     = flag.Int("max_width", 1000, "Max width")
		maxHeight    = flag.Int("max_height", 1000, "Max height")
		addr         = flag.String("addr", "localhost:8080", "Addr")
		bucketName   = flag.String("bucket_name", "", "Your bucket name")
		awsKeyID     = flag.String("aws_key_id", os.Getenv("AWS_ACCESS_KEY_ID"), "Your AWS_ACCESS_KEY_ID")
		awsSecretKey = flag.String("aws_secret_access_key", os.Getenv("AWS_SECRET_ACCESS_KEY"), "Your AWS_SECRET_ACCESS_KEY")
	)

	flag.Parse()

	log.Printf("Provider: [%s]\n", *providerName)
	log.Printf("Max Width: [%d]\n", *maxWidth)
	log.Printf("Max Height: [%d]\n", *maxHeight)
	if *providerName == "s3" {
		log.Printf("Bucket Name: [%s]\n", *bucketName)
		log.Printf("Aws Key ID: [%s]\n", *awsKeyID)
		log.Printf("Aws Secret Key: [%s]\n", "***")
	}

	log.Printf("Service started on %s\n", *addr)

	imageProxy, err := NewImageProxy(Config{
		ProviderName: *providerName,
		MaxWidth:     int64(*maxWidth),
		MaxHeight:    int64(*maxHeight),
		AwsKeyID:     *awsKeyID,
		AwsSecretKey: *awsSecretKey,
		BucketName:   *bucketName,
		Addr:         *addr,
	})

	if err != nil {
		log.Panicf("Unable to start the image proxy: %s\n", err.Error())
		return
	}

	imageProxy.Run()
}
