#!/bin/bash

pushd network

./testfeed.sh

sleep 5

./testfungi.sh

sleep 5

./testfungi2.sh
popd 