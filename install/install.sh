#!/bin/sh

if ! git --version
then
    echo "git could not be found"
    exit
fi

if ! go version
then
    echo "go could not be found"
    exit
fi

git clone https://github.com/krcod/kraken kraken