package generator

import (
	"crypto/rand"
)

type GeneratorV4 struct {
	Generator
}

var _ Generator = (*GeneratorV4)(nil)

func (g *GeneratorV4) UUID() ([]byte, error) {
	res := make([]byte, 16)
	if _, err := rand.Reader.Read(res); err != nil {
		return nil, err
	}
	res[6] = (res[6] & 0x0F) | 0x40 // version number
	res[8] &= ^uint8(1 << 6)
	res[8] |= (1 << 7)
	return res, nil
}

func (g *GeneratorV4) Version() string {
	return "v4"
}

func (g *GeneratorV4) String(u []byte) string {
	return String(u)
}


