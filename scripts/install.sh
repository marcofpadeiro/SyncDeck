#!/bin/bash


default_config='{
    "server_ip": "127.0.0.1",
    "server_port": "5137",
    "unit_metadata": "'$HOME'/.config/syncdeck/units.json",
    "api_key": "",
    "backup_path": "'$HOME'/.cache/syncdeck",
    "backup_size": 0 
}
'

mkdir -p $HOME/.config/syncdeck
echo $default_config > $HOME/.config/syncdeck/config.json
echo "[]" > $HOME/.config/syncdeck/units.json
cd ../cmd

go build 
sudo mv cmd /usr/bin/syncdeck

echo ":: Default config located at \"$HOME/.config/syncdeck/config.json\", recommended to change before using tool"
echo ":: Looking for server 127.0.0.1 on port 5137 (change on config.json)"
echo ":: Change the 'api_key' to the same one as server to enable authentication"

echo ":: Installation complete, to start 'syncdeck'"
cd - &> /dev/null
