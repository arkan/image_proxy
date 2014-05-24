package main

import (
  "net/http"
  "net/url"
  "strconv"
)

type Request struct {
  Url            string         // The URL to transform
  Transformation Transformation // The transformation to do
}

func parseAsInt(values []string) int64 {
  if len(values) == 1 {
    c, err := strconv.Atoi(values[0])
    if err == nil {
      return int64(c)
    }
  }
  return int64(0)
}

// We will parse the HTTP Request and extract the url containing the picture as well
// as the requested transformtions(resize, flip, rotation)
func NewRequest(r *http.Request) (*Request, error) {
  values := r.URL.Query()
  width := parseAsInt(values["w"])
  height := parseAsInt(values["h"])
  fit := parseAsInt(values["fit"]) == 1
  rotation := parseAsInt(values["rot"])
  flipVertical := parseAsInt(values["fv"]) == 1
  flipHorizontal := parseAsInt(values["fh"]) == 1

  path := r.URL.Path[1:] // strip leading slash
  u, _ := url.Parse(path)
  us := u.String()
  if !u.IsAbs() {
    us = "/" + u.String()
  }
  return &Request{
    Url: us,
    Transformation: Transformation{
      Width:          int64(width),
      Height:         int64(height),
      FlipVertical:   flipVertical,
      FlipHorizontal: flipHorizontal,
      Fit:            fit,
      Rotate:         rotation,
    },
  }, nil
}
