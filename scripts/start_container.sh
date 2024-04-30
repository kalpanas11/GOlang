#!/bin/bash
set -e

# Pull the Docker image from Docker Hub
docker pull kalpanas11/GOlang

# Run the Docker image as a container
docker run -d -p 5000:5000 kalpanas11/GOlang
