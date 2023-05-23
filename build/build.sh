#!/bin/bash
APPDIR=gamepc_monitor
BASEPATH=/root/workspace/$APPDIR
mkdir $BASEPATH/target
cd $BASEPATH
go build -o appexe
cp -rf $BASEPATH/bin $BASEPATH/target/bin
mv $BASEPATH/appexe $BASEPATH/target/bin/appexe
cp -rf $BASEPATH/target $BASEPATH/build/$APPDIR
ls $BASEPATH/target
