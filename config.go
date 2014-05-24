package main

type Config struct {
  ProviderName string
  MaxWidth     int64
  MaxHeight    int64
  BucketName   string
  AwsKeyID     string
  AwsSecretKey string
  Addr         string
}
