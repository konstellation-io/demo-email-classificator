#!/bin/bash

echo "Tidy repairs_handler node go.mod..."
cd classificator/src/repairs_handler
go mod tidy

echo "Tidy stats_storer node go.mod...."
cd ../stats_storer
go mod tidy

echo "Tidy exitpoint node go.mod..."
cd ../exitpoint
go mod tidy

echo "Done"