#!/bin/bash

if [ -z "$1" ]
then
  echo "Please provide a service name"
  exit 1
fi

service_name=$1
app_dir='../internal/app/'

if [ -d $app_dir$service_name ]
then
  echo "Service "$service_name" already exists"
  exit 1
fi

mkdir -p $app_dir$service_name
touch $app_dir$service_name/{routes.go,handlers.go,service.go,middleware.go,model.go,repository.go}

echo -e "Service \033[1;32m$service_name\033[0m created successfully"
