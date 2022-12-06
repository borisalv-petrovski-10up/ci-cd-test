SHELL := /bin/bash

run-and-deploy-flexible-service:
	bash scripts/run-docker-build-and-deploy.sh

generate-ci-image-signer:
	bash scripts/generate-and-push-ci-image-signer.sh