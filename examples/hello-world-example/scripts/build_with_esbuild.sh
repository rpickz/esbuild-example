#!/usr/bin/env bash

time esbuild --bundle ./src/index.js --loader:.js=jsx --outfile=out.js

time yarn build
