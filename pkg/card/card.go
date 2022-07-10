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
	Name() string
	ManaCost() int
	TypeOfCard() typeOfCard
	SetName(name string)
	SetManaCost(manaCost int)
	SetTypeOfCard(typeOfCard typeOfCard)
}

type BaseOfCard struct {
	name       string     //name of card
	manaCost   int        // mana cost of card
	typeOfCard typeOfCard // what type card is
}

func (card BaseOfCard) Name() string {
	return card.name
}

func (card BaseOfCard) ManaCost() string {
	return card.name
}

func (card BaseOfCard) TypeOfCard() string {
	return card.name
}

func (card *BaseOfCard) SetName(name string) {
	card.name = name
}

func (card *BaseOfCard) SetManaCost(manaCost int) {
	card.manaCost = manaCost
}

func (card *BaseOfCard) SetTypeOfCard(typeOfCard typeOfCard) {
	card.typeOfCard = typeOfCard
}
