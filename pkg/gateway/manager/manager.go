package manager

import (
	cmap "github.com/nevercase/concurrent-map"
)

type Manager struct {
	workers  *workerHub
	handlers cmap.ConcurrentMap
}
