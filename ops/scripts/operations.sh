#!/usr/bin/env bash

if [[ "${1}" == 'unit' ]]; then
    make unit
elif [[ "${1}" == 'coverage' ]]; then
    make coverage
elif [[ "${1}" == 'lint' ]]; then
    make lint
elif [[ "${1}" == 'functional' ]]; then
    make functional
elif [[ "${1}" == 'fixtures' ]]; then
    flag=true
    echo "Waiting for DB to be up..."
    while $flag; do
      nc -vz mysql 3306 > /dev/null 2>&1
      if [[ $? -eq 0 ]]; then
        flag=false
      fi
    done
    make fixtures
elif [[ "${1}" == 'dependencies' ]]; then
  make deps
fi
