#!/bin/bash

if [ -d "./discord-api-docs" ]; then
  cd "./discord-api-docs"
  git pull
  cd ..
else
  git clone "https://github.com/discordapp/discord-api-docs"
fi

python converter.py