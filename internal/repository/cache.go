package repository

type CodeCache interface {
	SetCode(codekey string, code string, expire int64) error
	GetCode(codekey string) (string, error)
}
