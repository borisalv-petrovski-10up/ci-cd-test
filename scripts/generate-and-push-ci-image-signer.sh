#!/bin/bash

_CI_IMAGE_SIGNER=eu.gcr.io/sk-borislav/ci-image-signer:latest
docker build --tag=${_CI_IMAGE_SIGNER} -f ././ci-image-signer/Dockerfile ci-image-signer
docker push ${_CI_IMAGE_SIGNER}