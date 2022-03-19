package entities

import "rpg/dice"

type Entity interface {
	GetHP() uint
	SetHP(uint)
	IsDead() bool
	GetGold() uint
	SetGold(uint)
	Attack(Entity, dice.Die)
	Defend(dice.Die) uint
}
