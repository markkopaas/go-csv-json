#!/usr/bin/env bash

echo "LatD;LatM;LatS;NS;LonD;LonM;LonS;EW;City;State"

for i in `seq -f %0.0f 1 $1`
do
    echo "$i;52;48;N;97;23;23;W;Yankton;SD"
done