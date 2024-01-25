#!/bin/bash

res=$(kubectl get pods -n swh)
status=$(word 8, $(res))
echo $status
