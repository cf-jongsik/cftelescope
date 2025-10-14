#!/bin/bash
if [ "$(uname)" = "Darwin" ]; then
    echo "This script is not supported on mac os because macOS doesn't support grep -r --include"
    exit 1
fi
grep -r "client\." ./cloudflare-go --include="*_test.go" | sed 's/.*client\.//' | sed 's/(.*//' | sort | uniq | grep "Get|List" > command_list.txt