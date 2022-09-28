#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

echo "running 'make regen'"
make regen

echo "checking for diffs"
diffs="$(git diff --stat)"
if [[ "$diffs" != "" ]]; then
	echo "found diffs after running 'make regen'"
	echo "please run & commit the changed files"
	exit 1
fi
