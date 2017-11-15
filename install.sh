#!/bin/sh

cd ./tin
go install
cd ..

cd ./tout
go install
cd ..
