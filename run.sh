#!/bin/bash
set -e

#Bulding Docker image
echo "Bulding a docker image for Hex to RGB color code"
docker build -t hex-rgb-color:v1 .
echo ""

#Spawning a container
echo "Spinning a container for Hex to RGB color code"
docker run -d --name hex-to-rgb -p 8080:8080 hex-rgb-color:v1
echo ""

#Making a curl request
echo "Making some HTTP GET requests to the container"
echo ""
curl http://localhost:8080/convert/345675
echo ""
curl http://localhost:8080/convert/345676
echo ""
curl http://localhost:8080/convert/345677
echo ""
curl http://localhost:8080/convert/345678
echo ""
echo "Happy color coding... !!!"
