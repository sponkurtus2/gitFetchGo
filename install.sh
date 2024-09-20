#!/bin/bash


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
