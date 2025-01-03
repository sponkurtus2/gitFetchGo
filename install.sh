#!/bin/bash

# Check if the files exist before installation
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
	echo "Installing on Linux..."
	
	file="./builds/gitFetchGo"
	if [ ! -f "$file" ]; then
		echo "Error: $file not found"
		exit 1
	fi

	sudo mv "$file" /usr/local/bin/ || {
		echo "Error: Failed to move file to /usr/local/bin/"
		exit 1
	}
	sudo chmod +x /usr/local/bin/"$file"

	echo "Installation successful!"
	echo "Usage: gitFetchGo <github-name>"

elif [[ "$OSTYPE" == "darwin"* ]]; then
	echo "Installing on macOS..."
	
	file="./builds/gitFetchGo"
	if [ ! -f "$file" ]; then
		echo "Error: $file not found"
		exit 1
	fi

	sudo mv "$file" /usr/local/bin/ || {
		echo "Error: Failed to move file to /usr/local/bin/"
		exit 1
	}
	sudo chmod +x /usr/local/bin/"$file"

	echo "Installation successful!"
	echo "Usage: gitFetchGo <github-name>"

elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" ]]; then
	echo "Installing on Windows..."
	
	file="./builds/gitFetchGo.exe"
	if [ ! -f "$file" ]; then
		echo "Error: $file not found"
		exit 1
	fi

	# Check if running with admin privileges
	if ! net session &>/dev/null; then
		echo "Error: Please run as Administrator"
		exit 1
	fi

	install_dir="C:/Program Files/GitFetchGo"
	mkdir -p "$install_dir" || {
		echo "Error: Failed to create installation directory"
		exit 1
	}
	
	mv "$file" "$install_dir/" || {
		echo "Error: Failed to move file to $install_dir"
		exit 1
	}

	# Add to PATH if not already present
	if [[ ":$PATH:" != *":$install_dir:"* ]]; then
		setx PATH "$PATH;$install_dir" /M
	fi

	echo "Installation successful!"
	echo "Usage: gitFetchGo <github-name>"

else
	echo "Error: Unsupported operating system: $OSTYPE"
	exit 1
fi
