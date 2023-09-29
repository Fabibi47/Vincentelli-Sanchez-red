package monsters

import "red/global/player"

type Monster struct {
	PV         int
	Name       string
	Speed      int
	Damage     int
	Zone       string
	Drop       int
	Exp        int
	Affliction string
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
	Gold         int
	Exp          int
}

// Forest:
var Jagras Monster = Monster{50, "Jagras", 10, 5, "Forest", 25, 10, "None"}
var Mernos Monster = Monster{25, "Mernos", 5, 5, "Forest", 5, 5, "None"}
var Aptonoth Monster = Monster{100, "Aptonoth", 10, 3, "Forest", 10, 5, "None"}
var Kestodon Monster = Monster{75, "Kestodon", 50, 1, "Forest", 15, 7, "None"}
var Great_Jagras Boss = Boss{300, "Great Jagras", 25, 10, player.Jagras_scale, 0.90, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Roll", 1.0, "None"}}, "none", 100, 50}

// Desert:
var Apceros Monster = Monster{200, "Apceros", 5, 10, "Desert", 40, 10, "None"}
var Noios Monster = Monster{100, "Noios", 10, 10, "Desert", 25, 10, "None"}
var Dellex Monster = Monster{50, "Delex", 35, 15, "Desert", 50, 15, "None"}
var Barroth Boss = Boss{500, "Barroth", 40, 40, player.Barroth_scalp, 0.7, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Headbutt", 1.0, "None"}, Attacks{"Mud Armor", 0.0, "Armor"}}, "none", 200, 200}

// Swamp:
var Shamos Monster = Monster{250, "Shamos", 25, 40, "Swamp", 100, 25, "None"}
var Raphinos Monster = Monster{150, "Raphinos", 50, 15, "Swamp", 75, 20, "None"}
var Girros Monster = Monster{300, "Girros", 30, 30, "Swamp", 200, 35, "None"}
var Rathalos Boss = Boss{1000, "Rathalos", 75, 100, player.Rathalos_cortex, 0.5, []Attacks{Attacks{"Poisonned Scratch", 0.8, "Poison"}, Attacks{"Fire Blast", 1.0, "Burn"}}, "none", 300, 400}

// Ancient's Land:
var Barnos Monster = Monster{250, "Barnos", 40, 50, "Ancient Land", 300, 50, "None"}
var Dodogama Monster = Monster{500, "Dodogama", 25, 100, "Ancient Land", 500, 75, "None"}
var Uragaan Monster = Monster{500, "Uragaan", 50, 50, "Ancient Land", 500, 75, "None"}
var Nergigante Boss = Boss{2000, "Nergigante", 100, 200, player.Nergigante_talon, 0.75, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Allahu Akbar", 1.2, "None"}}, "none", 500, 1000}

// UnderWorld
var Xeno_Jiiva Boss = Boss{5000, "Xeno' Jiiva", 200, 200, player.Xeno_jiiva_gem, 0.5, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Laser Beam", 1.5, "None"}}, "None", 1000, 2500}

// Sacred Lands
var Alatreon Boss = Boss{10000, "Alatreon", 250, 500, player.Analtreon, 1.0, []Attacks{Attacks{"Scratch", 0.8, "None"}, Attacks{"Divine Blow", 5.0, "you dead"}}, "None", 10000, 10000}
