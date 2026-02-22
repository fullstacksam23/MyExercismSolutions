package dndcharacter

import (
    "math/rand"
    "math"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score)-10)/2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	n1 := rand.Intn(6) + 1
	n2 := rand.Intn(6) + 1
	n3 := rand.Intn(6) + 1
	n4 := rand.Intn(6) + 1

	sum := n1 + n2 + n3 + n4

	minVal := math.Min(
		math.Min(float64(n1), float64(n2)),
		math.Min(float64(n3), float64(n4)),
	)

	return sum - int(minVal)
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	Strength := Ability()
    Dexterity := Ability()
    Constitution := Ability()
    Intelligence := Ability()
    Wisdom := Ability()
    Charisma := Ability()
    Hitpoints := 10+ Modifier(Constitution)
	return Character{
		Strength:     Strength,
		Dexterity:    Dexterity,
		Constitution: Constitution,
		Intelligence: Intelligence,
		Wisdom:       Wisdom,
		Charisma:     Charisma,
		Hitpoints:    Hitpoints,
	}
}
