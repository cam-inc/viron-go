#!/bin/bash

echo "validate commit message start"
echo "Running script for commit file: $1"

MSG_FILE="$1"
MSG="$(head -n1 "$MSG_FILE")"

REGEX="^(feat|fix|chore|docs|style|refactor|perf|test|revert|ci|build)(\([a-zA-Z0-9_-]+\))?: .+"

if ! echo "$MSG" | grep -Eq "$REGEX"; then
  echo "❌ The commit message does not comply with Conventional Commits"
  echo "Example: feat(lib): add endpoint"
  echo "Allowed types: feat, fix, chore, docs, style, refactor, perf, test, revert, ci, build"
  exit 1
fi