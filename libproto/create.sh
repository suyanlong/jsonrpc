#!/usr/bin/env bash
protoc --go_out=import_path=libproto:. *.proto