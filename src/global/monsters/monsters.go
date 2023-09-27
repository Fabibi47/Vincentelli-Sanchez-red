package monsters

import "player"

type Monster struct {
	PV     int
	Name   string
	Speed  int
	Damage int
	Zone   string
	Drop   int
	Exp    int
}

type Boss struct {
	PV           int
	Name         string
	Speed        int
	Damage       int
	Drop         player.Item
	Drop_chances float64
	Attacks      []Attacks
	Affliction   string
}

// Forest:
var Jagras Monster = Monster{50, "Jagras", 10, 5, "Forest", 25, 10}
var Mernos Monster = Monster{25, "Mernos", 5, 5, "Forest", 5, 5}
var Aptonoth Monster = Monster{100, "Aptonoth", 10, 3, "Forest", 10, 5}
var Kestodon Monster = Monster{75, "Kestodon", 50, 1, "Forest", 15, 7}
var Great_Jagras Boss = Boss{500, "Great Jagras", 10, 25, player.Jagras_scale, 0.90, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Roll", 1.0, "None"}}, "none"}

// Desert:
var Apceros Monster = Monster{200, "Apceros", 5, 10, "Desert", 40, 10}
var Noios Monster = Monster{100, "Noios", 10, 10, "Desert", 25, 10}
var Dellex Monster = Monster{50, "Delex", 35, 15, "Desert", 50, 15}
var Barroth Boss = Boss{250, "Barroth", 5, 75, player.Barroth_scalp, 0.7, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Headbutt", 1.0, "None"}, Attacks{"Mud Armor", 0.0, "Armor"}}, "none"}

// Swamp:
var Shamos Monster = Monster{250, "Shamos", 25, 40, "Swamp", 100, 25}
var Raphinos Monster = Monster{150, "Raphinos", 50, 15, "Swamp", 75, 20}
var Girros Monster = Monster{300, "Girros", 30, 30, "Swamp", 200, 35}
var Rathalos Boss = Boss{275, "Rathalos", 5, 100, player.Rathalos_cortex, 0.5, []Attacks{Attacks{"Poisonned Scratch", 0.8, "Poison"}, Attacks{"Fire Blast", 1.0, "Burn"}}, "none"}

// Ancient's Land:
var Barnos Monster = Monster{250, "Barnos", 40, 50, "Ancient Land", 300, 50}
var Dodogama Monster = Monster{500, "Dodogama", 25, 100, "Ancient Land", 500, 75}
var Uragaan Monster = Monster{500, "Uragaan", 50, 50, "Ancient Land", 500, 75}
var Nergigante Boss = Boss{150, "Nergigante", 5, 75, player.Nergigante_talon, 0.75, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Allahu Akbar", 1.2, "None"}}, "none"}

// UnderWorld
var Xeno_Jiiva Boss = Boss{150, "Xeno' Jiiva", 5, 75, player.Xeno_jiiva_gem, 0.5, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Laser Beam", 1.5, "None"}}, "None"}

// Sacred Lands
var Alatreon Boss = Boss{2000, "Alatreon", 100, 100, player.Analtreon, 1.0, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Divine Blow", 5.0, "you dead"}}, "None"}
