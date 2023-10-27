#!/usr/bin/env bash

docker run --rm -v "$(pwd)/vmware-rest-proxy:/helm-docs" -u $(id -u) jnorwood/helm-docs:latest
