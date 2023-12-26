#!/bin/bash

mkdir -p $HOME/.config/syncdeck
cp ../configs/server.json $HOME/.config/syncdeck/server.json
echo ":: Default config located at \"$HOME/.config/syncdeck/config.json\", recommended to change before using tool"

cd ../api

go build 
sudo mv api /usr/bin/syncdeck-api

echo ":: Installation complete, to start api run 'syncdeck-api'"
cd -
