#!/bin/bash

echo "Tidy repairs_handler node go.mod..."
cd classificator/src/common/repairs_handler
go mod tidy

echo "Tidy stats_storer node go.mod...."
cd ../stats_storer
go mod tidy

echo "Tidy exitpoint node go.mod..."
cd ../exitpoint
go mod tidy

echo "Tidy email_classificator node go.mod..."
cd ../../go-classificator/email_classificator
go mod tidy

echo "Tidy etl node go.mod..."
cd ../etl
go mod tidy


echo "Done"