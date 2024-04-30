#!/bin/bash
set -e

# Pull the Docker image from Docker Hub
sudo docker pull kalpana111/go-template-simple-form-app

# Run the Docker image as a container
sudo docker run -d -p 5000:5000 kalpana111/go-template-simple-form-app
