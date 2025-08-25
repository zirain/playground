#!/bin/bash
set -e

echo "Requesting new envelope session..."
SESSION_ID=$(curl -s -D - 'http://localhost:8080/?envelope=true' | grep -i 'x-session-id:' | awk '{print $2}' | tr -d '\r\n')
echo "SESSION_ID: $SESSION_ID"

echo "Sending request with x-session-id header..."
curl -v 'http://localhost:8080/?envelope=true' -H "x-session-id: $SESSION_ID" | jq
