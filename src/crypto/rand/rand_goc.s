// +build js,wasm,goc

// Copyright 2019 Andreas T Jonsson.

#include "textflag.h"

TEXT ·getRandomValues(SB), NOSPLIT, $0
  CallImport
  RET
