package monsters

import "player"

type Monster struct {
	PV     int
	Name   string
	Speed  int
	Damage int
	Zone   string
	Drop   int
}

type Boss struct {
	PV           int
	Name         string
	Speed        int
	Damage       int
	Drop         player.Item
	Drop_chances float64
	Attacks      []Attacks
}

// Forest:
var Jagras Monster = Monster{40, "Jagras", 15, 20, "Forest", 10}
var Mernos Monster = Monster{40, "Mernos", 20, 20, "Forest", 10}
var Aptonoth Monster = Monster{40, "Aptonoth", 10, 10, "Forest", 10}
var Gastodon Monster = Monster{40, "Gastodon", 25, 10, "Forest", 10}
var Great_Jagras Boss = Boss{300, "Great Jagras", 3, 50, player.Jagras_scale, 0.75, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Roll", 1.0, "None"}}}

// Desert:
var Apceros Monster = Monster{40, "Apceros", 15, 20, "Desert", 10}
var Noios Monster = Monster{40, "Noios", 20, 20, "Desert", 10}
var Dellex Monster = Monster{40, "Delex", 25, 10, "Desert", 10}
var Barroth Boss = Boss{250, "Barroth", 5, 75, player.Barroth_scalp, 0.7, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Headbutt", 1.0, "None"}, Attacks{"Mud Armor", 0.0, "Armor"}}}

// Swamp:
var Shamos Monster = Monster{40, "Shamos", 15, 20, "Swamp", 10}
var Raphinos Monster = Monster{40, "Raphinos", 20, 20, "Swamp", 10}
var Girros Monster = Monster{40, "Girros", 15, 20, "Swamp", 10}
var Rathalos Boss = Boss{275, "Rathalos", 5, 100, player.Rathalos_cortex, 0.5, []Attacks{Attacks{"Poisonned Scratch", 0.8, "Poison"}, Attacks{"Fire Blast", 1.0, "Burn"}}}

//Ancient's Land:
var Barnos Monster = Monster{40, "Barnos", 20, 20, "Ancient Land", 10}
var Dodogama Monster = Monster{300, "Dodogama", 3, 50, "Ancient Land", 10}
var Uragaan Monster = Monster{150, "Uragaan", 5, 50, "Ancient Land", 10}
var Nergigante Boss = Boss{150, "Nergigante", 5, 75, player.Nergigante_talon, 0.75, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Allahu Akbar", 1.2, "None"}}}

//UnderWorld
var Xeno_Jiiva Boss = Boss{150, "Xeno' Jiiva", 5, 75, player.Xeno_jiiva_gem, 0.5, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Laser Beam", 1.5, "None"}}}

//Sacred Lands
var Alatreon Boss = Boss{2000, "Alatreon", 100, 100, player.Analtreon, 1.0, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Divine Blow", 5.0, "you dead"}}}
