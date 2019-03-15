// +build js,wasm,goc

// Copyright 2019 Andreas T Jonsson.

#include "textflag.h"

TEXT ·writeFile(SB), NOSPLIT, $0
  CallImport
  RET

TEXT ·readFile(SB), NOSPLIT, $0
  CallImport
  RET

TEXT ·openFile(SB), NOSPLIT, $0
  CallImport
  RET

TEXT ·closeFile(SB), NOSPLIT, $0
  CallImport
  RET
