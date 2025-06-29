package valorant

func NewWeapon(name WeaponName, weaponType WeaponType, cost uint32) *Weapon {
	return &Weapon{
		Name: name,
		Type: weaponType,
		Cost: cost,
	}
}

func CreateWeapons() (weapons []*Weapon) {
	weapons = append(weapons, NewWeapon(WeaponName_CLASSIC, WeaponType_RIFLE, 0))
	weapons = append(weapons, NewWeapon(WeaponName_SHORTY, WeaponType_SIDEARM, 200))
	weapons = append(weapons, NewWeapon(WeaponName_FRENZY, WeaponType_SIDEARM, 400))
	weapons = append(weapons, NewWeapon(WeaponName_GHOST, WeaponType_SIDEARM, 500))
	weapons = append(weapons, NewWeapon(WeaponName_SHERIFF, WeaponType_SIDEARM, 800))
	weapons = append(weapons, NewWeapon(WeaponName_STINGER, WeaponType_SMG, 1000))
	weapons = append(weapons, NewWeapon(WeaponName_SPECTRE, WeaponType_SMG, 1600))
	weapons = append(weapons, NewWeapon(WeaponName_BUCKY, WeaponType_SHOTGUN, 850))
	weapons = append(weapons, NewWeapon(WeaponName_JUDGE, WeaponType_SHOTGUN, 1600))
	weapons = append(weapons, NewWeapon(WeaponName_BULLDOG, WeaponType_RIFLE, 2100))
	weapons = append(weapons, NewWeapon(WeaponName_GUARDIAN, WeaponType_RIFLE, 2700))
	weapons = append(weapons, NewWeapon(WeaponName_PHANTOM, WeaponType_RIFLE, 2900))
	weapons = append(weapons, NewWeapon(WeaponName_VANDAL, WeaponType_RIFLE, 2900))
	weapons = append(weapons, NewWeapon(WeaponName_MARSHAL, WeaponType_SNIPER_RIFLE, 1100))
	weapons = append(weapons, NewWeapon(WeaponName_OUTLAW, WeaponType_SNIPER_RIFLE, 2400))
	weapons = append(weapons, NewWeapon(WeaponName_OPERATOR, WeaponType_SNIPER_RIFLE, 4700))
	weapons = append(weapons, NewWeapon(WeaponName_ARES, WeaponType_MACHINE_GUN, 1600))
	weapons = append(weapons, NewWeapon(WeaponName_ODIN, WeaponType_MACHINE_GUN, 3200))
	weapons = append(weapons, NewWeapon(WeaponName_KNIFE, WeaponType_MELEE, 0))

	return weapons
}
