package constants

type WeaponType uint8

const (
	Sidearm WeaponType = iota
	SMG
	Shotgun
	Rifle
	SniperRifle
	MachineGun
	Melee
)

var weaponType = map[WeaponType]string{
	Sidearm:     "Sidearm",
	SMG:         "SMG",
	Shotgun:     "Shotgun",
	Rifle:       "Rifle",
	SniperRifle: "Sniper Rifle",
	MachineGun:  "Machine Gun",
	Melee:       "Melee",
}

func (ss WeaponType) String() string {
	return weaponType[ss]
}

type WeaponName uint8

const (
	Classic WeaponName = iota
	Shorty
	Frenzy
	Ghost
	Sheriff
	Stinger
	Spectre
	Bucky
	Judge
	Bulldog
	Guardian
	Phantom
	Vandal
	Marshal
	Outlaw
	Operator
	Ares
	Odin
	Knife
)

var weaponName = map[WeaponName]string{
	Classic:  "Classic",
	Shorty:   "Shorty",
	Frenzy:   "Frenzy",
	Ghost:    "Ghost",
	Sheriff:  "Sheriff",
	Stinger:  "Stinger",
	Spectre:  "Spectre",
	Bucky:    "Bucky",
	Judge:    "Judge",
	Bulldog:  "Bulldog",
	Guardian: "Guardian",
	Phantom:  "Phantom",
	Vandal:   "Vandal",
	Marshal:  "Marshal",
	Outlaw:   "Outlaw",
	Operator: "Operator",
	Ares:     "Ares",
	Odin:     "Odin",
	Knife:    "Knife",
}

func (ss WeaponName) String() string {
	return weaponName[ss]
}

type WallPenetration uint8

const (
	Low WallPenetration = iota
	Medium
	High
)

var wallPenetration = map[WallPenetration]string{
	Low:    "Low",
	Medium: "Medium",
	High:   "High",
}

func (ss WallPenetration) String() string {
	return wallPenetration[ss]
}

type Weapon struct {
	Name            WeaponName
	Type            WeaponType
	Cost            uint16
	Damage          uint16
	Range           uint16
	WallPenetration WallPenetration
}

var Weapons = map[WeaponName]Weapon{
	Classic:  {Name: Classic, Type: Sidearm, Cost: 0},
	Shorty:   {Name: Shorty, Type: Sidearm, Cost: 150},
	Frenzy:   {Name: Frenzy, Type: Sidearm, Cost: 450},
	Ghost:    {Name: Ghost, Type: Sidearm, Cost: 500},
	Sheriff:  {Name: Sheriff, Type: Sidearm, Cost: 800},
	Stinger:  {Name: Stinger, Type: SMG, Cost: 950},
	Spectre:  {Name: Spectre, Type: SMG, Cost: 1600},
	Bucky:    {Name: Bucky, Type: Shotgun, Cost: 850},
	Judge:    {Name: Judge, Type: Shotgun, Cost: 1850},
	Bulldog:  {Name: Bulldog, Type: Rifle, Cost: 2050},
	Guardian: {Name: Guardian, Type: Rifle, Cost: 2250},
	Phantom:  {Name: Phantom, Type: Rifle, Cost: 2900},
	Vandal:   {Name: Vandal, Type: Rifle, Cost: 2900},
	Marshal:  {Name: Marshal, Type: SniperRifle, Cost: 950},
	Outlaw:   {Name: Outlaw, Type: SniperRifle, Cost: 2400},
	Operator: {Name: Operator, Type: SniperRifle, Cost: 4700},
	Ares:     {Name: Ares, Type: MachineGun, Cost: 1600},
	Odin:     {Name: Odin, Type: MachineGun, Cost: 3200},
	Knife:    {Name: Knife, Type: Melee, Cost: 0},
}
