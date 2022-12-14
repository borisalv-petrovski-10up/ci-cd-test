#!/bin/bash

IMAGE_NAME=eu.gcr.io/sk-borislav/flexible:latest
docker build -t ${IMAGE_NAME} -f ././Dockerfile .
docker push ${IMAGE_NAME}
gcloud app deploy --image-url=${IMAGE_NAME} --appyaml=services/app-ae-flexible/app.yaml
