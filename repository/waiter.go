package repository

import "sync"

var AppWaiter *sync.WaitGroup = new(sync.WaitGroup)
