#!/bin/bash

BASE_URL="http://localhost:7000"

# Endpoints
endpoints=(
  "GET /user"
  "POST /user"
  "PUT /user"
  "DELETE /user"
  "POST /user-disable"
  "POST /email-verify"
  "POST /login"
  "POST /forgot-password"
)

# Request data for POST and PUT requests
userData='{"first_name":"Samarth","last_name":"Asthan","email":"samarthasthan25@gmail.com","password":"Password123!","phone_no":"19557037766","role_name":"user"}'
loginData='{"email":"samarthasthan25@gmail.com","password":"Password123!"}'

# Function to perform a curl request
perform_request() {
  local method=$1
  local endpoint=$2
  local data=$3

  case $method in
    GET|DELETE)
      curl -X $method "${BASE_URL}${endpoint}" -H "Content-Type: application/json" -w "\n";;
    POST|PUT)
      curl -X $method "${BASE_URL}${endpoint}" -H "Content-Type: application/json" -d "$data" -w "\n";;
  esac
}

# Loop through endpoints and perform requests
for i in {1..500}; do
  for endpoint in "${endpoints[@]}"; do
    IFS=' ' read -r method path <<< "$endpoint"
    
    # Use different data for specific endpoints
    case $path in
      "/user")
        perform_request $method $path "$userData";;
      "/login")
        perform_request $method $path "$loginData";;
      *)
        perform_request $method $path "";;
    esac
  done
done
