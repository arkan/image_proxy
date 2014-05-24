# Image Proxy #

Image Proxy is a proxy server for processing images on the fly.

It allows you to process images stored on s3 or anywhere on the Internet.

## How does it work ? ##

### Available options ###

    -addr="localhost:8080": Addr
    -aws_key_id="????": Your AWS_ACCESS_KEY_ID
    -aws_secret_access_key="???": Your AWS_SECRET_ACCESS_KEY
    -bucket_name="": Your bucket name
    -max_height=1000: Max height
    -max_width=1000: Max width
    -provider="http": Provider name - Available http, s3

### Providers ###

#### Http ####

The urls looks like this: `http://localhost:8080/http://server.com/path/to/the/image.jpg?w=800&h=600`

#### S3 ####

The urls looks like this: `http://localhost:8080/path/to/the/image.jpg?w=800&h=600`


### Available options

    w=800       # To specify the width
    h=600       # To specify the height
    rot=90      # To apply a rotation of the image - 90/180/270
    fh=1        # To apply a horizontal flip
    fv=1        # To apply a vertical flip
    fit=1       # To apply a fit constraint

## TODO ##

* Add a domain whitelist
