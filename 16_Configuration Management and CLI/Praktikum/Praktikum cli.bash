#!/bin/bash

# Get arguments
dir_name="$1 Sunday at $(date)"
facebook_username="$2"
linkedin_username="$3"

# Create directories
mkdir -p "$dir_name/about_me/personal"
mkdir -p "$dir_name/about_me/professional"
mkdir -p "$dir_name/my_friends"
mkdir -p "$dir_name/my_system_info"

# Create files
echo "https://www.facebook.com/$facebook_username" > "$dir_name/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$linkedin_username" > "$dir_name/about_me/professional/linkedin.txt"
curl -o "$dir_name/my_friends/list_of_my_friends.txt" https://example.com/list_of_my_friends.txt
echo "My username: $(whoami)" > "$dir_name/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$dir_name/my_system_info/about_this_laptop.txt"
ping -c 3 google.com > "$dir_name/my_system_info/internet_connection.txt"
