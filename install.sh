#!/bin/bash

if [[ $EUID -ne 0 ]]; then
	echo "This must be executed as root in order to install"
	exit 1
fi

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
	echo "Installing on linux device"

	file="gitFetchGo"
	sudo mv "$file" /usr/local/bin/
	sudo chmod +x /usr/local/bin/"$file"

	echo "Done"
	echo "Usage: ./gitFetchGo <github-name>"

elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" ]]; then
	echo "Installing on windows"

	file="gitFetchGo.git.exe"
	mv "$file" "C:/Program Files/$file"

	echo "Done"
	echo "Usage: ./gitFetchGo <github-name>"

else
	echo "Error detecting os type"
	exit 1
fi

# Install needed library
echo "Downloading neccesary library..."
go install github.com/qeesung/image2ascii@latest

if [ $? -eq 0 ]; then
	echo "Library installed succesfully"
else
	echo "Error installing library"
	exit 1
fi

