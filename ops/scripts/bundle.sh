#!/usr/bin/env bash

TARGET_DIR="tmp/artifacts"
PARENT_DIR=$(dirname $(pwd))
BASE_NAME=$(pwd)

if [[ ! -d "${TARGET_DIR}" ]]; then
  mkdir -p "${TARGET_DIR}"
fi

if [[ $(uname -s) != 'Linux' ]]; then
  sed -i '' 's/VIRTUAL_ENV=.*/VIRTUAL_ENV=NEW_VENV_PATH/' venv/bin/activate
else
  sed -i 's/VIRTUAL_ENV=.*/VIRTUAL_ENV=NEW_VENV_PATH/' venv/bin/activate
fi

tar --exclude='tmp' --exclude='.git' -cvzf "${TARGET_DIR}/panorama.tar.gz" -C ${PARENT_DIR} ${BASE_NAME}
