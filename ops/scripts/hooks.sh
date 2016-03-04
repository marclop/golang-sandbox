#!/usr/bin/env bash

if [[ ! -d .git/hooks ]]; then
  mkdir .git/hooks
fi

cp ops/scripts/pre-commit.sh .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
