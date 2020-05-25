#!/bin/bash
set -ex

cp -r tests/testdata/input/* tests/testdata/output/

./bin/kpt fn run tests/testdata/output \
  --image kpt-remove-resource:latest \
  -- kind=Deployment name=matching-name namespace=matching-namespace

git diff --exit-code tests/testdata/output/
