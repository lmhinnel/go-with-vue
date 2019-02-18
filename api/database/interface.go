package database

type Database interface {
	IsAlive() bool
}
