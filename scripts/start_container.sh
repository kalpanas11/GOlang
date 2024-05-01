#!/bin/bash
set -e
echo "------starting container -----"
# Pull the Docker image from Docker Hub
 docker pull kalpana111/go-template-simple-form-app

# Run the Docker image as a container
 docker run -d -p 5000:5000 kalpana111/go-template-simple-form-app

