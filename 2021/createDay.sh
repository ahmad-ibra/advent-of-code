#!/bin/bash

# Check if a day argument is provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

day="$1"

# Validate the day input
if ! [[ "$day" =~ ^[0-9]{2}$ ]]; then
    echo "Invalid day format. Please enter a valid day (dd)."
    exit 1
fi

# Prepare the source and destination directory names
src_dir="src/dayXX"
dst_dir="src/day$day"

# Check if the source directory exists
if [ ! -d "$src_dir" ]; then
    echo "Source directory $src_dir not found."
    exit 1
fi

# Check if the destination directory already exists
if [ -d "$dst_dir" ]; then
    echo "Destination directory $dst_dir already exists."
    exit 1
fi

# Copy the source directory to the destination and rename files
cp -r "$src_dir" "$dst_dir"
mv "$dst_dir/dayXX.ts" "$dst_dir/day$day.ts"

# Replace XX with the day in file contents using sed
sed -i '' -e "s|XX|$day|g" "$dst_dir/day$day.ts"

echo "Directory copied, files renamed, and contents updated successfully."


