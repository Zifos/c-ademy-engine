

---
URL - /authn/signup
METHOD - POST
REQUEST BODY 
{
  "username": "feco",
  "password": "balde",
}

RESPONSE CODE - 200 OK

---
URL - /authn/get-token
METHOD POST
REQUEST BODY
{
  "username": "feco",
  "password": "balde",
}


RESPONSE CODE - 200 OK
RESPONSE BODY 
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiZmVjbyIsImlhdCI6MTYxNjIzOTAyMn0.1J7"
}

---
URL - /exec
METHOD POST
HEADER "Authorization": "Bearer <token>"
REQUEST BODY
{
  "language": 0,
  "code": "base64"
  "input": "base64"
  "expected_output": "base64"
  // al momento en el que finalice, va a enviar los resultados a esta url con el formato de /get-result
  "webhook_url": "https://api.c-ademy.com/solution/123321/result/123321"
}

RESPONSE CODE - 200 OK
RESPONSE BODY
{
  // uint64 with a unique 
  "exec_code": 321123
}

---
URL - /get-result/:exec_code
METHOD GET
HEADER "Authorization": "Bearer <token>"

RESPONSE STATUS - 404 "Not found"

RESPONSE STATUS - 425 "Too early"
RESPONSE BODY
{
  "status": "in_queue" | "running"
}

RESPONSE STATUS - 200 "OK"
RESPONSE BODY
{
  "status": "error" | "success"
  "stdout": "base64"

  // Tengo que ver como hacer esta parte porque mansa paja
  "diff": "base64"
  "execution_time": uint64 // time in miliseconds
}


