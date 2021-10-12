#!/bin/bash

echo "=> /opt/elk contents"
ls -l /opt/elk

echo "=> setting up ${BROKERS}"

IFS=','
read -ra ADDR <<<"$BROKERS"
for BROKER in "${ADDR[@]}"; do
    /opt/elk/bootstrap --broker=${BROKER}
done
