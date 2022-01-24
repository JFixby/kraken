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

if [ ! -d kraken ] then
   git clone https://github.com/krcod/kraken kraken
if

cd kraken
git reset HEAD --hard

dir
