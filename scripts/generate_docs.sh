#!/bin/bash

forge doc

cat docs/src/README.md docs/src/SUMMARY.md > docs/src/index.md

echo -e "---\ntitle: \"Protocol contracts\"\n---\n" | cat - docs/src/index.md > temp && mv temp docs/src/index.md

if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's|contracts/|protocol/contracts/|g' docs/src/index.md
else
    sed -i 's|contracts/|protocol/contracts/|g' docs/src/index.md
fi

find docs/src/ -type f -name "*.md" -exec sed -i.bak -E 's|(https://github.com/RWAs-labs/protocol-contracts/blob/)[^/]+/|\1main/v2/|g' {} +

find docs/src/ -type f -name "*.bak" -exec rm {} +
