#!/usr/bin/env bash
curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"id":"miquel"}' \
         localhost:3004/api/register

curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"id":"pere"}' \
         localhost:3004/api/register

curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"id":"lol"}' \
         localhost:3004/api/register

curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"id":"wtf"}' \
         localhost:3004/api/register

curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"trusts":{"id":"miquel"},"trusted":{"id":"pere"}}' \
         localhost:3004/api/relation

curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"trusts":{"id":"pere"},"trusted":{"id":"miquel"}}' \
         localhost:3004/api/relation

curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"trusts":{"id":"pere"},"trusted":{"id":"lol"}}' \
         localhost:3004/api/relation

curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"trusts":{"id":"wtf"},"trusted":{"id":"lol"}}' \
         localhost:3004/api/relation

curl localhost:3004/api/log
