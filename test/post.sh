#!/usr/bin/env bash
curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"name":"gnuillem"}' \
         localhost:3004/api/register/