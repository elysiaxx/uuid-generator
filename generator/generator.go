package generator

import "fmt"

type Generator interface {
	UUID() ([]byte, error)
	String([]byte) string
	Version() string
}

func String(u []byte) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[:4], u[4:6], u[6:8], u[8:10], u[10:16])
}