#!/usr/bin/env bash
if [ ! -e ./etcd ]; then
  echo "Downloading v0.3"
  wget https://github.com/coreos/etcd/releases/download/v0.3.0/etcd-v0.3.0-darwin-amd64.zip
  unzip etcd-v0.3.0-darwin-amd64.zip
  mv ./etcd-v0.3.0-darwin-amd64/etcd ./etcd
  rm -rf ./etcd-v0.3.0-darwin-amd64 etcd-v0.3.0-darwin-amd64.zip
fi
