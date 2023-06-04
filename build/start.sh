#!/usr/bin/env bash

if [ -d "$HOME/.permastar/config.yaml" ]; then
  /opt/permastar-server daemon
else
  /opt/permastar-server init && /opt/permastar-server daemon
fi