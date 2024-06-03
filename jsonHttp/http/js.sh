#!/bin/sh

JSON=$(curl -s "https://jsonplaceholder.typicode.com/posts/")

printf "%s\n" "$JSON"

echo "$JSON" | jq '.data[] | {email: .email, name: .name}'