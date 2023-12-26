#!/bin/bash

mkdir -p $HOME/.config/syncdeck
cp ../configs/client.json $HOME/.config/syncdeck/config.json
cp ../configs/units.json $HOME/.config/syncdeck/units.json
echo ":: Default config located at \"$HOME/.config/syncdeck/config.json\", recommended to change before using tool"

cd ../cmd

go build 
sudo mv cmd /usr/bin/syncdeck

echo ":: Installation complete, to start use script run 'syncdeck'"
cd -
