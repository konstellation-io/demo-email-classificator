#!/bin/bash

echo "Tidy email_filter node go.mod..."
cd classificator/src/email_filter
go mod tidy

echo "Tidy stats_storer node go.mod...."
cd ../stats_storer
go mod tidy

echo "Tidy exitpoint node go.mod..."
cd ../exitpoint
go mod tidy

echo "Done"