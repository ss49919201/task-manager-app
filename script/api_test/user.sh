#!/bin/bash

function postUser() {
    curl -w'\n' -i -X POST -H "Content-Type: application/json" -d '{"name":"太郎"}' localhost:12345/users
}

function getUser() {
    curl -w'\n' -i localhost:12345/users
}

echo "=========POST /users========="
postUser
echo "============================="
echo "=========Get /users========="
getUser
echo "============================="
