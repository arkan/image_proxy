package main

import (
  "io/ioutil"
  "log"
  "net/http"

  "github.com/crowdmob/goamz/aws"
  "github.com/crowdmob/goamz/s3"
)

type FileFetcher interface {
  Fetch(string) ([]byte, error)
}

type HttpFileFetcher struct{}

func (httpFileFetcher *HttpFileFetcher) Fetch(url string) ([]byte, error) {
  res, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()
  bytes, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }
  return bytes, nil
}

type S3FileFetcher struct {
  AwsKeyId     string
  AwsSecretKey string
  BucketName   string
}

func (s3FileFetcher *S3FileFetcher) Fetch(url string) ([]byte, error) {
  log.Printf("Fetching s3 file [%s]\n", url)
  auth := aws.Auth{AccessKey: s3FileFetcher.AwsKeyId, SecretKey: s3FileFetcher.AwsSecretKey}
  store := s3.New(auth, aws.EUWest)
  log.Printf("Buckername: [%s]\n", s3FileFetcher.BucketName)
  bucket := store.Bucket(s3FileFetcher.BucketName)
  s3_reader, err := bucket.GetReader(url)

  if err != nil {
    log.Printf("Error while getting: %s\n", err.Error())
    return nil, err
  }

  b, err := ioutil.ReadAll(s3_reader)
  if err != nil {
    log.Printf("Error while reading: %s\n", err.Error())
    return nil, err
  }

  return b, nil
}
