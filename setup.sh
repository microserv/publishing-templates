#!/bin/bash

echo "Installing Go libs..."
go get -v github.com/gin-gonic/gin
go get -v gopkg.in/mgo.v2
echo "Finished installing Go libs!"
