#!/bin/bash

APPDIR=gamepc_monitor

# home
HOME_PATH=/home/work/$APPDIR

# set exec 
chmod +x $HOME_PATH/bin/appexe

cd $HOME_PATH/

# start command
bin/appexe