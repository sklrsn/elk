#!/bin/bash

echo "=> /opt/elk contents"
ls -l /opt/elk

echo "=> setting up ${BROKER}"
/opt/elk/bootstrap --broker=${BROKER}
