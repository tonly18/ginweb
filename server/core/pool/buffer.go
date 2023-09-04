package pool

import (
	"bytes"
	"sync"
)

var Buffer64Pool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

var Buffer512Pool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 512))
	},
}

var Buffer1024Pool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

var Buffer2408Pool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 2408))
	},
}

var Buffer4096Pool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 4096))
	},
}
