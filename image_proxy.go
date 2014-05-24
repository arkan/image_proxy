package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

type ImageProxy struct {
	Config      Config
	FileFetcher FileFetcher
}

func NewImageProxy(config Config) (*ImageProxy, error) {
	var fileFetcher FileFetcher
	if config.ProviderName == "http" {
		fileFetcher = &HttpFileFetcher{}
	} else if config.ProviderName == "s3" {
		fileFetcher = &S3FileFetcher{
			BucketName:   config.BucketName,
			AwsSecretKey: config.AwsSecretKey,
			AwsKeyId:     config.AwsKeyID,
		}
	} else {
		return nil, fmt.Errorf("The provider %s is not supported.", config.ProviderName)
	}
	return &ImageProxy{
		Config:      config,
		FileFetcher: fileFetcher,
	}, nil
}

func rootHandler(image_proxy *ImageProxy) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		start := time.Now()
		defer func() {
			w.Header().Add("X-RUNTIME", time.Since(start).String())
			log.Printf("Request took %s", time.Since(start))
		}()

		if r.Method != "GET" && r.Method != "HEAD" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// time.Sleep(2 * time.Second)

		var t1 time.Time
		t1 = time.Now()
		req, err := NewRequest(r)
		log.Printf(" -> NewRequest Took %s", time.Since(t1))

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error while parsing the request: %s", err.Error())
			return
		}

		t1 = time.Now()
		originalData, err := image_proxy.FileFetcher.Fetch(req.Url)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error while fetching the file: %s", err.Error())
			return
		}
		log.Printf(" -> Fetch Took %s", time.Since(t1))

		t1 = time.Now()
		transformedData, err := applyTransformation(originalData, req.Transformation)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error while applying the transformtion: %s", err.Error())
			return
		}
		log.Printf(" -> applyTransformation Took %s", time.Since(t1))

		buf := new(bytes.Buffer)
		buf.Write(transformedData)

		w.Header().Add("Cache-Control", "public, max-age=315576000")
		w.Header().Add("Content-Length", strconv.Itoa(len(transformedData)))
		w.Header().Add("Expires", time.Now().Add(time.Hour*24*365*10).In(time.UTC).Format(time.RFC1123))

		if r.Method == "GET" {
			io.CopyN(w, bufio.NewReader(buf), int64(len(transformedData)))
		}
	})
}

func (image_proxy *ImageProxy) Run() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Fatal(http.ListenAndServe(image_proxy.Config.Addr, rootHandler(image_proxy)))
}
