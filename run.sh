#!/bin/bash
docker run -it --rm -p 8000:3000 -v ~/work/golang/go-ffmpeg:/go/src/go-ffmpeg -w /go/src/go-ffmpeg ffmpeg