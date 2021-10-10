#!/bin/bash

echo "=> /opt/elk contents"
ls -l /opt/elk

echo "=> setting up rabbitmq"
/opt/elk/bootstrap --broker=rabbitmq

echo "=> setting up kafka"
/opt/elk/bootstrap --broker=kafka
