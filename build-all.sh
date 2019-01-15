#!/bin/bash

# build BE.
cd be
./build.sh
cd ..

# build FE
cd fe/melapoly-tracker
./build.sh