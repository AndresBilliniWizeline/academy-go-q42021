package structs

import (
	"strconv"
)

type Pokemon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (p *Pokemon) Odd() bool {
	return p.Id%2 == 0
}

func (p *Pokemon) Even() bool {
	return p.Id%2 == 1
}

func (p *Pokemon) SetInfoCSV(line []string) {
	id, _ := strconv.Atoi(line[0])
	p.Id = id
	p.Name = line[1]
	p.Url = line[2]
}
