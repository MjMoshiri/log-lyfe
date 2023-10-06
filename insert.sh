#!/bin/bash

# Total number of JSON objects
count=$(jq '. | length' insert.json)

for i in $(seq 0 $((count-1)))
do
    jq ".[$i]" insert.json > temp.json
    curl -X POST -H "Authorization: canon" -H "Content-Type: application/json" -d @temp.json http://localhost:8080/insert
    echo ""
    sleep 1
done

rm temp.json
