#!/bin/bash

# Friendly colors
red=$'\e[1;31m'
green=$'\e[1;32m'
yellow=$'\e[1;33m'
blue=$'\e[1;34m'
magenta=$'\e[1;35m'
cyan=$'\e[1;36m'
dim=$'\e[2m'
end=$'\e[0m'

ROOT="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

YEAR=${1:-$(date +%Y)}
DAY=${2:-$(date +%d)}

YEAR_DIR="${ROOT}/${YEAR}"
DAY_DIR="${ROOT}/${YEAR}/${DAY}"

echo "${yellow}Preparing: (${YEAR}/${DAY}) ...${end}"

# Make sure the dirs exist
[[ ! -d $YEAR_DIR ]] && mkdir -p $YEAR_DIR
[[ -d $DAY_DIR ]] && echo "${yellow}Day is already prepared - ${cyan}Good luck${end}" && exit 0

# Setup base project
cp -r ${ROOT}/utils/starter_go ${DAY_DIR}
cd ${DAY_DIR}
go mod tidy

echo "${yellow}GOOD LUCK - AND ${cyan}GOD SPEED!${end}"
