#!/bin/bash

set -xe

root=$(git rev-parse --show-toplevel)
cd "$root/backend" && go fmt './...'
cd "$root/frontend" && prettier --write .