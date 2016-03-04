#!/usr/bin/env bash

result=$(make unit | tee /dev/tty)

if [[ $result =~ "FAILED" ]]
then
  echo -e "\tUnit tests failed ! Aborting commit." >&2
  exit 1;
fi
