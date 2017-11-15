#!/bin/sh

pushd .
cd tin
go install
popd

pushd .
cd tout
go install
popd
