package db

type IDB interface {
	RawScan(sql string, dest interface{}, values ...interface{}) error
}
