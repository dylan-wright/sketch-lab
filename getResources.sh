#!/bin/bash
mkdir resources
for line in $(cat resources.csv); do
	IFS=, read -a array <<< $line
	curl ${array[1]} > resources/${array[0]}.txt
done
