#!/bin/bash

pushd network

./deployfungi.sh 1

sleep 5

./deployfeed.sh 1

popd 