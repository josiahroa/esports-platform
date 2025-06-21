package constants

type MapName uint8

const (
	Sunset MapName = iota
	Lotus
	Pearl
	Fracture
	Breeze
	Icebox
	Bind
	Haven
	Split
	Ascent
)

var mapNames = map[MapName]string{
	Sunset:   "Sunset",
	Lotus:    "Lotus",
	Fracture: "Fracture",
	Breeze:   "Breeze",
	Icebox:   "Icebox",
	Bind:     "Bind",
	Haven:    "Haven",
	Split:    "Split",
	Ascent:   "Ascent",
}

func (m MapName) String() string {
	return mapNames[m]
}

type PlantSite uint8

const (
	PlantSiteA PlantSite = iota
	PlantSiteB
	PlantSiteC
)

var plantSites = map[PlantSite]string{
	PlantSiteA: "A",
	PlantSiteB: "B",
	PlantSiteC: "C",
}

func (p PlantSite) String() string {
	return plantSites[p]
}

type Map struct {
	Name       MapName
	PlantSites []PlantSite
}

var Maps = map[MapName]Map{
	Sunset: {
		Name:       Sunset,
		PlantSites: []PlantSite{PlantSiteA, PlantSiteB, PlantSiteC},
	},
	Lotus: {
		Name:       Lotus,
		PlantSites: []PlantSite{PlantSiteA, PlantSiteB, PlantSiteC},
	},
	Fracture: {
		Name:       Fracture,
		PlantSites: []PlantSite{PlantSiteA, PlantSiteB, PlantSiteC},
	},
}
