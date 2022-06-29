/*

As a package developer, you can think of it this way: everything with a capital letter is your
public API, while everything with a lower-case letter is for internal use only.
Creating Package

Additionally, we now have comments above each of our exported types. This isn’t just a nice thing, 
this is a rule enforced by most Go static analyzers, and required if we end up using 
a documentation generation tool.

*/
package npcs

//Power describes the attack and defense power of an NPC
type Power struct {
		Attack 	int 
		Defense int 
}

// Location describes where in the virtual world an NPC exists
type Location struct {
	X float64
	Y float64
	Z float64
}

// NonPlayerCharacter represents metadata for an in-game creature
type NonPlayerCharacter struct {
	Name	string
	Speed	int
	HP 		int
	Power 	Power
	Loc 	Location
}







