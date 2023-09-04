//go:build zombiezen_sqlite

package sqlite

import "zombiezen.com/go/sqlite"

type (
	Conn         = sqlite.Conn
	Stmt         = sqlite.Stmt
	FunctionImpl = sqlite.FunctionImpl
	Context      = sqlite.Context
	Value        = sqlite.Value
	ResultCode   = sqlite.ResultCode
)

const (
	TypeNull = sqlite.TypeNull

	OpenNoMutex     = sqlite.OpenNoMutex
	OpenReadOnly    = sqlite.OpenReadOnly
	OpenURI         = sqlite.OpenURI
	OpenWAL         = sqlite.OpenWAL
	OpenCreate      = sqlite.OpenCreate
	OpenReadWrite   = sqlite.OpenReadWrite
	OpenSharedCache = sqlite.OpenSharedCache

	ResultCodeConstraintUnique = sqlite.ResultConstraintUnique
	ResultCodeInterrupt        = sqlite.ResultInterrupt
)

var (
	BlobValue = sqlite.BlobValue
	OpenConn  = sqlite.OpenConn
)

// This produces an error code even if it's not an underlying sqlite error. This could differ from
// the crawshaw implementation.
func GetResultCode(err error) (ResultCode, bool) {
	return sqlite.ErrCode(err), true
}

// Error code extraction looks something like this:
//sqlite.ErrCode(err) == sqlite.ResultInterrupt
