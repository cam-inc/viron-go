#!/bin/bash

MSG_FILE="$1"
MSG="$(head -n1 "$MSG_FILE")"

REGEX="^(feat|fix|chore|docs|style|refactor|perf|test|revert|ci|build)(\([a-zA-Z0-9_-]+\))?: .+"

if ! echo "$MSG" | grep -Eq "$REGEX"; then
  echo "❌ コミットメッセージが Conventional Commits に準拠していません"
  echo "例: feat(api): add user endpoint"
  echo "許可タイプ: feat, fix, chore, docs, style, refactor, perf, test, revert, ci, build"
  exit 1
fi