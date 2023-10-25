package hook

import (
	"github.com/rs/zerolog"
)

// hook
type ZeroLogHook struct{}

func (h *ZeroLogHook) Run(logger *zerolog.Event, level zerolog.Level, msg string) {}
