#! /bin/bash

# Exit on any error
set -e

echo "Copy the input text to your clipboard, and hit enter to continue"
read

day="day$(ls | grep -E '^day[0-9]{2}' | sort | tail -n 1 | awk -F'day' '{printf("%02d\n", $2 + 1)}')"
mkdir ${day}

# Get the input text from the clipboard
pbpaste > ${day}/input.txt

cd ${day}
go mod init github.com/boffbowsh/aoc2022/${day}

git add .
git commit -m "Input for ${day}"
