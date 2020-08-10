package identita

import (
	"io"
	"time"
)

type Persona struct {
	Version        int
	Creation       time.Time
	Country        int // ISO 3166-1 numeric
	Pin            string
	BirthDate      time.Time
	BirthZip       string
	Zip            string
	ChecksumPerson int
	Expires        int
	Verified       bool
	Padding        int
	Checksum       int
}

type SpainPersona Persona

func (p *SpainPersona) Decode(r io.Reader) (err error) {
	return
}

func (p *SpainPersona) Encode(w io.Writer) (err error) {
	return
}
