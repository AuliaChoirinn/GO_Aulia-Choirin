#!/bin/bash

# Check if arguments are provided
if [[ $# -ne 3 ]]; then
    echo "Usage: $0 <folder_name> <facebook_username> <linkedin_username>"
    exit 1
fi

# Create directory structure
mkdir -p "$1/about_me/personal"
mkdir -p "$1/about_me/professional"
mkdir -p "$1/my_friends"
mkdir -p "$1/my_system_info"

# Create files
touch "$1/about_me/personal/facebook.txt"
touch "$1/about_me/professional/linkedin.txt"
touch "$1/my_friends/list_of_my_friends.txt"
touch "$1/my_system_info/about_this_laptop.txt"
touch "$1/my_system_info/internet_connection.txt"

# Populate files with data
echo "https://www.facebook.com/$2" > "$1/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$3" > "$1/about_me/professional/linkedin.txt"
curl https://raw.githubusercontent.com/auliachoirin/friends/main/friends.txt > "$1/my_friends/list_of_my_friends.txt"
echo "My username: $(whoami)" > "$1/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$1/my_system_info/about_this_laptop.txt"
ping -c 3 google.com > "$1/my_system_info/internet_connection.txt"

# Print success message
echo "Folder structure created successfully in $(pwd)/$1"
