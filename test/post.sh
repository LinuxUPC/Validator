#!/usr/bin/env bash
curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"trusts":{"name":"Pedro Sanchez"},{"trusted":{"name":"Quim torri"}}' \
         localhost:3004/api/register/