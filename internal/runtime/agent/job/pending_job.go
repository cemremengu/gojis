package job

import (
	"github.com/gojisvm/gojis/internal/runtime/lang"
	"github.com/gojisvm/gojis/internal/runtime/realm"
)

// PendingJob represents a pending job that is to be executed by an agent.
type PendingJob struct {
	Job            string
	Arguments      []lang.Value
	Realm          *realm.Realm
	ScriptOrModule interface{} // FIXME: 8.4, Table 24
	HostDefined    lang.InternalValue
}
