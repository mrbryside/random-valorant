package valorander

import "math/rand"

var (
	mapList = []MapDetail{
		{
			Name:     "Assent",
			ImageUrl: "https://static.wikia.nocookie.net/valorant/images/e/e7/Loading_Screen_Ascent.png/revision/latest/scale-to-width-down/1000?cb=20200607180020",
		},
		{
			Name:     "Bind",
			ImageUrl: "https://static.wikia.nocookie.net/valorant/images/2/23/Loading_Screen_Bind.png/revision/latest/scale-to-width-down/1000?cb=20200620202316",
		},
		{
			Name:     "Split",
			ImageUrl: "https://static.wikia.nocookie.net/valorant/images/d/d6/Loading_Screen_Split.png/revision/latest/scale-to-width-down/1000?cb=20230411161807",
		},
		{
			Name:     "Haven",
			ImageUrl: "https://static.wikia.nocookie.net/valorant/images/7/70/Loading_Screen_Haven.png/revision/latest/scale-to-width-down/1000?cb=20200620202335",
		},
		{
			Name:     "Icebox",
			ImageUrl: "https://static.wikia.nocookie.net/valorant/images/1/13/Loading_Screen_Icebox.png/revision/latest/scale-to-width-down/1000?cb=20201015084446",
		},
		{
			Name:     "Breeze",
			ImageUrl: "https://static.wikia.nocookie.net/valorant/images/1/10/Loading_Screen_Breeze.png/revision/latest/scale-to-width-down/1000?cb=20210427160616",
		},
		{
			Name:     "Fracture",
			ImageUrl: "https://static.wikia.nocookie.net/valorant/images/f/fc/Loading_Screen_Fracture.png/revision/latest/scale-to-width-down/1000?cb=20210908143656",
		},
		{
			Name:     "Breeze",
			ImageUrl: "https://static.wikia.nocookie.net/valorant/images/1/10/Loading_Screen_Breeze.png/revision/latest/scale-to-width-down/1000?cb=20210427160616",
		},
		{
			Name:     "Pearl",
			ImageUrl: "https://static.wikia.nocookie.net/valorant/images/a/af/Loading_Screen_Pearl.png/revision/latest/scale-to-width-down/1000?cb=20220622132842",
		},
		{
			Name:     "Lotus",
			ImageUrl: "https://static.wikia.nocookie.net/valorant/images/d/d0/Loading_Screen_Lotus.png/revision/latest/scale-to-width-down/1000?cb=20230106163526",
		},
		{
			Name:     "Sunset",
			ImageUrl: "https://static.wikia.nocookie.net/valorant/images/5/5c/Loading_Screen_Sunset.png/revision/latest/scale-to-width-down/1000?cb=20230829125442",
		},
	}
)

type MapDetail struct {
	Name     string
	ImageUrl string
}

func RandomMap() MapDetail {
	return mapList[rand.Intn(len(mapList))]
}
