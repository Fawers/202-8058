package entities

import (
	"rpg/dice"
	"testing"
)

// new*
func TestNewCreatureReturnsCreatureWithGivenAttributes(t *testing.T) {
	expected := Creature{
		maxHp: 100, hp: 100,
		gold:   0,
		attack: 10, defense: 8,
	}
	actual := NewCreature(100, 10, 8)

	if *actual != expected {
		t.Fatalf("\n%#v\n!= %#v", *actual, expected)
	}
}

func TestNewMonsterReturnsCreatureWithAttributesRolled(t *testing.T) {
	expected := Creature{
		maxHp: 70, hp: 70, gold: 30,
		attack: 9, defense: 9,
	}
	actual := NewMonster(&MonsterCreationOptions{
		50, dice.NewLoaded(20),
		6, dice.NewLoaded(3),
		20, dice.NewLoaded(10),
	})

	if *actual != expected {
		t.Fatalf("\n%#v\n!= %#v", *actual, expected)
	}
}

// gethp
func TestGetHPReturnsCorrectHP(t *testing.T) {
	c := NewCreature(100, 10, 10)

	if hp := c.GetHP(); hp != 100 {
		t.Fatalf("%d != 100", hp)
	}
}

// sethp
func TestSetHPSetsHPCorrectly(t *testing.T) {
	c := NewCreature(100, 10, 10)
	c.SetHP(80)

	if hp := c.GetHP(); hp != 80 {
		t.Fatalf("%d != 80", hp)
	}

	c.SetHP(120)

	if hp := c.GetHP(); hp != 100 {
		t.Fatalf("%d != 100", hp)
	}
}

// isdead
func TestIsDeadReturnsCorrectStatus(t *testing.T) {
	c := NewCreature(100, 10, 10)

	if c.IsDead() {
		t.Fatalf("criatura não deveria estar morta. HP == %d", c.GetHP())
	}

	c.SetHP(0)

	if !c.IsDead() {
		t.Fatalf("criatura deveria estar morta. HP == %d", c.GetHP())
	}
}

// getgold
func TestGetGoldReturnsZeroForNewCreature(t *testing.T) {
	c := NewCreature(100, 10, 10)

	if g := c.GetGold(); g != 0 {
		t.Fatalf("criatura deveria estar pobre, mas tem %d gold", g)
	}
}

// setgold
func TestSetGoldSetsGoldCorrectly(t *testing.T) {
	c := NewCreature(100, 10, 10)
	c.SetGold(50)

	if g := c.GetGold(); g != 50 {
		t.Fatalf("criatura deveria ter 50 gold, mas tem %d", g)
	}

	c.SetGold(30)

	if g := c.GetGold(); g != 30 {
		t.Fatalf("criatura deveria ter 30 gold, mas tem %d", g)
	}
}

// attack
func TestAttackCausesDamageToWeakerEntity(t *testing.T) {
	at := NewCreature(100, 10, 10)
	df := NewCreature(100, 10, 8)

	expectedHP := uint(98)

	at.Attack(df, dice.NewLoaded(6))

	if hp := df.GetHP(); hp != expectedHP {
		t.Fatalf("HP do defensor deveria ser %d; é %d", expectedHP, hp)
	}
}

func TestAttackDoesNothingWhenDefenseIsEqualOrBetterThanAttack(t *testing.T) {
	at := NewCreature(100, 10, 10)
	df := NewCreature(100, 10, 10)
	d := dice.NewLoaded(6)

	expectedHP := uint(100)

	at.Attack(df, d)

	if hp := df.GetHP(); hp != expectedHP {
		t.Fatalf("HP do defensor deveria ser %d; é %d", expectedHP, hp)
	}

	df.defense += 2
	at.Attack(df, d)

	if hp := df.GetHP(); hp != expectedHP {
		t.Fatalf("HP do defensor deveria ser %d; é %d", expectedHP, hp)
	}
}

// defend
func TestDefendReturnsDefenseRate(t *testing.T) {
	d := dice.NewLoaded(6)
	c := NewCreature(100, 10, 10)

	if def := c.Defend(d); def != 10 {
		t.Fatalf("defesa deveria ser 10; é %d", def)
	}

	d = dice.New(6)
	def := c.Defend(d)

	for i := uint(0); i <= 10; i += 2 {
		if def == i {
			return
		}
	}

	t.Fatalf("defesa deveria ser um de [0, 2, 4, 6, 8, 10]; é %d", def)
}
