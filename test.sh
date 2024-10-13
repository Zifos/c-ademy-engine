#!/bin/bash

# Set the base URL for the API
BASE_URL="http://localhost:1323"

# Function to print section headers
print_header() {
    echo
    echo "===== $1 ====="
    echo
}

# Test root endpoint
print_header "Testing root endpoint"
curl $BASE_URL/

# Test signup endpoint
print_header "Testing signup endpoint"
curl -X POST $BASE_URL/authn/signup \
  -H "Content-Type: application/json" \
  -d '{"username":"example_user", "password":"testpassword"}'

# Test get-token endpoint
print_header "Testing get-token endpoint"
TOKEN=$(curl -s -X POST $BASE_URL/authn/get-token \
  -H "Content-Type: application/json" \
  -d '{"username":"example_user", "password":"testpassword"}' | jq -r .token)

echo "Received token: $TOKEN"

# Test exec endpoint
print_header "Testing exec endpoint"
curl -X POST $BASE_URL/exec \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "language": "python3.8.1",
    "file_name": "test.py",
    "code": "'"$(echo 'print("Hello, World!")' | base64)"'",
    "input": "'"$(echo -n '' | base64)"'",
    "expected_output": "Hello, World!",
    "webhook_url": "http://example.com/webhook"
  }'

echo
echo "API testing completed."