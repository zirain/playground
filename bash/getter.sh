#!/bin/bash

function verify_geater() {
  local actual="$1"
  local expected="$2"
  if  [[ "$actual" -le "$expected" ]]; then
    echo "expected $actual to be greater than $expected"
    exit 1
  fi
}

verify_geater "2" "1"
verify_geater 2 1