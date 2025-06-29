package valorant

func NewMap(name MapName, plantSites []PlantSite) *Map {
	return &Map{
		Name:       name,
		PlantSites: plantSites,
	}
}

func CreateMaps() (maps []*Map) {
	maps = append(maps, NewMap(MapName_SUNSET, []PlantSite{
		PlantSite_PLANT_SITE_A,
		PlantSite_PLANT_SITE_B,
	}))
	maps = append(maps, NewMap(MapName_LOTUS, []PlantSite{
		PlantSite_PLANT_SITE_A,
		PlantSite_PLANT_SITE_B,
		PlantSite_PLANT_SITE_C,
	}))
	maps = append(maps, NewMap(MapName_PEARL, []PlantSite{
		PlantSite_PLANT_SITE_A,
		PlantSite_PLANT_SITE_B,
	}))
	maps = append(maps, NewMap(MapName_FRACTURE, []PlantSite{
		PlantSite_PLANT_SITE_A,
		PlantSite_PLANT_SITE_B,
	}))
	maps = append(maps, NewMap(MapName_BREEZE, []PlantSite{
		PlantSite_PLANT_SITE_A,
		PlantSite_PLANT_SITE_B,
	}))
	maps = append(maps, NewMap(MapName_ICEBOX, []PlantSite{
		PlantSite_PLANT_SITE_A,
		PlantSite_PLANT_SITE_B,
	}))
	maps = append(maps, NewMap(MapName_ASCENT, []PlantSite{
		PlantSite_PLANT_SITE_A,
		PlantSite_PLANT_SITE_B,
	}))
	maps = append(maps, NewMap(MapName_SPLIT, []PlantSite{
		PlantSite_PLANT_SITE_A,
		PlantSite_PLANT_SITE_B,
	}))
	maps = append(maps, NewMap(MapName_HAVEN, []PlantSite{
		PlantSite_PLANT_SITE_A,
		PlantSite_PLANT_SITE_B,
		PlantSite_PLANT_SITE_C,
	}))
	maps = append(maps, NewMap(MapName_BIND, []PlantSite{
		PlantSite_PLANT_SITE_A,
		PlantSite_PLANT_SITE_B,
	}))

	return maps
}
