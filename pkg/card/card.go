package card

type typeOfCard int

const (
	Combat typeOfCard = iota
	Defense
	Healing
	Spell
	Armor
	Weapon
)

type Card interface {
}

//----------------------------------------------------------------------------------------------------------------

type baseOfCard struct {
	name       string     //name of card
	manaCost   int        // mana cost of card
	typeOfCard typeOfCard // what type card is
}

func (card baseOfCard) Name() string {
	return card.name
}

func (card baseOfCard) ManaCost() int {
	return card.manaCost
}

func (card baseOfCard) TypeOfCard() typeOfCard {
	return card.typeOfCard
}

func (card *baseOfCard) SetName(name string) {
	card.name = name
}

func (card *baseOfCard) SetManaCost(manaCost int) {
	card.manaCost = manaCost
}

func (card *baseOfCard) SetTypeOfCard(typeOfCard typeOfCard) {
	card.typeOfCard = typeOfCard
}

//----------------------------------------------------------------------------------------------------------------

type combatCard struct {
	baseOfCard
	hp     int
	attack int
}

func (card combatCard) Hp() int {
	return card.hp
}

func (card combatCard) Attack() int {
	return card.attack
}

//----------------------------------------------------------------------------------------------------------------

type spellCard struct {
	baseOfCard
	value int
}

func (card spellCard) Value() int {
	return card.value
}

func (card *spellCard) SetValue(value int) {
	card.value = value
}

//----------------------------------------------------------------------------------------------------------------

type heroBuffCard struct {
	baseOfCard
	value int
}

func (card *heroBuffCard) Value() int {
	return card.value
}

func (card *heroBuffCard) SetValue(value int) {
	card.value = value
}
