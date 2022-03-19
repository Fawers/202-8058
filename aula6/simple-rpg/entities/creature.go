package entities

import "rpg/dice"

type Creature struct {
	maxHp, hp, gold uint
	attack, defense uint
}

type MonsterCreationOptions struct {
	MinHP uint
	HPDie dice.Die

	MinAttDef uint
	AttDefDie dice.Die

	MinGold uint
	GoldDie dice.Die
}

func NewCreature(maxHp, attack, defense uint) *Creature {
	p := new(Creature)
	p.maxHp = maxHp
	p.hp = maxHp
	p.gold = 0

	p.attack = attack
	p.defense = defense
	return p
}

func NewMonster(options *MonsterCreationOptions) *Creature {
	c := NewCreature(
		options.HPDie.Roll()+options.MinHP,
		options.AttDefDie.Roll()+options.MinAttDef,
		options.AttDefDie.Roll()+options.MinAttDef,
	)

	c.SetGold(options.GoldDie.Roll() + options.MinGold)
	return c
}

func (p *Creature) GetHP() uint {
	return p.hp
}

func (p *Creature) SetHP(newHp uint) {
	if newHp > p.maxHp {
		newHp = p.maxHp
	}

	p.hp = newHp
}

func (p *Creature) IsDead() bool {
	return p.hp == 0
}

func (p *Creature) GetGold() uint {
	return p.gold
}

func (p *Creature) SetGold(g uint) {
	p.gold = g
}

func (p *Creature) Attack(e Entity, d dice.Die) {
	roll := d.Roll()
	rate := roll * 100 / d.GetMax()
	attack := p.attack * rate / 100

	defense := e.Defend(d)

	if attack > defense {
		damage := attack - defense
		curHp := e.GetHP()

		if damage >= curHp {
			e.SetHP(0)
		} else {
			e.SetHP(curHp - damage)
		}
	}
}

func (p *Creature) Defend(d dice.Die) uint {
	roll := d.Roll()
	rate := roll * 100 / d.GetMax()
	defense := p.defense * rate / 100
	return defense
}

func init() {
	var _ Entity = &Creature{}
}
