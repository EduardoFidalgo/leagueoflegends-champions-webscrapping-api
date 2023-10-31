package models

type Character struct {
	Name       string
	Url        string
	History    string
	Role       string
	Difficulty string
	Skins      []string
	Skills     []string
}

var Characters []Character
