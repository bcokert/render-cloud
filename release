#!/usr/bin/env bash

if [ -z "$1" ]; then echo "Missing release type as first argument. Should be one of: 'major, minor, patch'"; exit 1; fi
if [ $1 != "major" ] && [ $1 != "minor" ] && [ $1 != "patch" ]; then echo "Release type must be one of 'major, minor, patch'"; exit 1; fi
if ! git branch | grep -q '* master'; then echo "You must be on the master branch"; exit 1; fi
if ! git status | grep "working directory clean"; then echo "You must be on a clean copy of master, with no staged or unstaged code"; exit 1; fi

VERSION=$(git describe --abbrev=0 --tags | sed 's/v//g')

MAJOR=$(echo "${VERSION}" | cut -d '.' -f 1)
MINOR=$(echo "${VERSION}" | cut -d '.' -f 2)
PATCH=$(echo "${VERSION}" | cut -d '.' -f 3)

echo "Previous Version: $MAJOR.$MINOR.$PATCH"

if   [ $1 = "major" ]; then MAJOR=$((MAJOR+=1)); MINOR=0; PATCH=0;
elif [ $1 = "minor" ]; then MINOR=$((MINOR+=1)); PATCH=0;
elif [ $1 = "patch" ]; then PATCH=$((PATCH+=1));
fi

echo "New Version: $MAJOR.$MINOR.$PATCH"

git tag "v$MAJOR.$MINOR.$PATCH"

git push --tags
