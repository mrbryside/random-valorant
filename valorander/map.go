package valorander

import "math/rand"

var (
	mapList = []string{"Ascent", "Bind", "Breeze", "Haven", "Icebox", "Split", "Fracture"}
)

func RandomMap() string {
	return mapList[rand.Intn(len(mapList))]
}
