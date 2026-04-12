package mdparser

import "sync"

var secrets []secretContainer

var secretMu sync.RWMutex
