package main

import (
	"encoding/json"
	"fmt"
)

var jsonValueShort json.RawMessage = json.RawMessage(`
	{
		"type": "test",
		"data": {
			"id": "1"
		}
	}
`)

type Struct01 struct {
	Type string       `json:"type"`
	Data Struct03Data `json:"data"`
}

func Unmarshal01(n int) {
	var obj Struct01
	for i := 0; i < n; i++ {
		json.Unmarshal(jsonValueShort, &obj)
	}
}

type Struct02 struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func Unmarshal02(n int) {
	var obj Struct02
	for i := 0; i < n; i++ {
		json.Unmarshal(jsonValueShort, &obj)
	}
}

type Struct03 struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type Struct03Data struct {
	ID string `json:"id"`
}

func Unmarshal03(n int) {
	var obj Struct03
	var data Struct03Data
	for i := 0; i < n; i++ {
		json.Unmarshal(jsonValueShort, &obj)

		json.Unmarshal(obj.Data, &data)
	}
}

var jsonValueLong json.RawMessage = json.RawMessage(`
	{
		"type": "test",
		"data": {
			"id1": "1",
			"id2": "2",
			"id3": "3",
			"id4": "4",
			"id5": "5",
			"id6": "6",
			"id7": "7",
			"id8": "8",
			"id9": "9",
			"id10": "10",
			"id11": "11",
			"id12": "12",
			"id13": "13",
			"id14": "14",
			"id15": "15",
			"id16": "16",
			"id17": "17",
			"id18": "18",
			"id19": "19",
			"id20": "20",
			"id21": "21",
			"id22": "22",
			"id23": "23",
			"id24": "24",
			"id25": "25",
			"id26": "26",
			"id27": "27",
			"id28": "28",
			"id29": "29",
			"id30": "30",
			"id31": "31",
			"id32": "32",
			"id33": "33",
			"id34": "34",
			"id35": "35",
			"id36": "36",
			"id37": "37",
			"id38": "38",
			"id39": "39",
			"id40": "40",
			"id41": "41",
			"id42": "42",
			"id43": "43",
			"id44": "44",
			"id45": "45",
			"id46": "46",
			"id47": "47",
			"id48": "48",
			"id49": "49",
			"id50": "50",
			"id51": "51",
			"id52": "52",
			"id53": "53",
			"id54": "54",
			"id55": "55",
			"id56": "56",
			"id57": "57",
			"id58": "58",
			"id59": "59",
			"id60": "60",
			"id61": "61",
			"id62": "62",
			"id63": "63",
			"id64": "64",
			"id65": "65",
			"id66": "66",
			"id67": "67",
			"id68": "68",
			"id69": "69",
			"id70": "70",
			"id71": "71",
			"id72": "72",
			"id73": "73",
			"id74": "74",
			"id75": "75",
			"id76": "76",
			"id77": "77",
			"id78": "78",
			"id79": "79",
			"id80": "80",
			"id81": "81",
			"id82": "82",
			"id83": "83",
			"id84": "84",
			"id85": "85",
			"id86": "86",
			"id87": "87",
			"id88": "88",
			"id89": "89",
			"id90": "90",
			"id91": "91",
			"id92": "92",
			"id93": "93",
			"id94": "94",
			"id95": "95",
			"id96": "96",
			"id97": "97",
			"id98": "98",
			"id99": "99",
			"id100": "100"
		}
}
`)

type Struct04 struct {
	Type string       `json:"type"`
	Data Struct06Data `json:"data"`
}

func Unmarshal04(n int) {
	var obj Struct04
	for i := 0; i < n; i++ {
		json.Unmarshal(jsonValueLong, &obj)
	}
}

type Struct05 struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func Unmarshal05(n int) {
	var obj Struct05
	for i := 0; i < n; i++ {
		json.Unmarshal(jsonValueLong, &obj)
	}
}

type Struct06 struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type Struct06Data struct {
	ID1   string `json:"id1"`
	ID2   string `json:"id2"`
	ID3   string `json:"id3"`
	ID4   string `json:"id4"`
	ID5   string `json:"id5"`
	ID6   string `json:"id6"`
	ID7   string `json:"id7"`
	ID8   string `json:"id8"`
	ID9   string `json:"id9"`
	ID10  string `json:"id10"`
	ID11  string `json:"id11"`
	ID12  string `json:"id12"`
	ID13  string `json:"id13"`
	ID14  string `json:"id14"`
	ID15  string `json:"id15"`
	ID16  string `json:"id16"`
	ID17  string `json:"id17"`
	ID18  string `json:"id18"`
	ID19  string `json:"id19"`
	ID20  string `json:"id20"`
	ID21  string `json:"id21"`
	ID22  string `json:"id22"`
	ID23  string `json:"id23"`
	ID24  string `json:"id24"`
	ID25  string `json:"id25"`
	ID26  string `json:"id26"`
	ID27  string `json:"id27"`
	ID28  string `json:"id28"`
	ID29  string `json:"id29"`
	ID30  string `json:"id30"`
	ID31  string `json:"id31"`
	ID32  string `json:"id32"`
	ID33  string `json:"id33"`
	ID34  string `json:"id34"`
	ID35  string `json:"id35"`
	ID36  string `json:"id36"`
	ID37  string `json:"id37"`
	ID38  string `json:"id38"`
	ID39  string `json:"id39"`
	ID40  string `json:"id40"`
	ID41  string `json:"id41"`
	ID42  string `json:"id42"`
	ID43  string `json:"id43"`
	ID44  string `json:"id44"`
	ID45  string `json:"id45"`
	ID46  string `json:"id46"`
	ID47  string `json:"id47"`
	ID48  string `json:"id48"`
	ID49  string `json:"id49"`
	ID50  string `json:"id50"`
	ID51  string `json:"id51"`
	ID52  string `json:"id52"`
	ID53  string `json:"id53"`
	ID54  string `json:"id54"`
	ID55  string `json:"id55"`
	ID56  string `json:"id56"`
	ID57  string `json:"id57"`
	ID58  string `json:"id58"`
	ID59  string `json:"id59"`
	ID60  string `json:"id60"`
	ID61  string `json:"id61"`
	ID62  string `json:"id62"`
	ID63  string `json:"id63"`
	ID64  string `json:"id64"`
	ID65  string `json:"id65"`
	ID66  string `json:"id66"`
	ID67  string `json:"id67"`
	ID68  string `json:"id68"`
	ID69  string `json:"id69"`
	ID70  string `json:"id70"`
	ID71  string `json:"id71"`
	ID72  string `json:"id72"`
	ID73  string `json:"id73"`
	ID74  string `json:"id74"`
	ID75  string `json:"id75"`
	ID76  string `json:"id76"`
	ID77  string `json:"id77"`
	ID78  string `json:"id78"`
	ID79  string `json:"id79"`
	ID80  string `json:"id80"`
	ID81  string `json:"id81"`
	ID82  string `json:"id82"`
	ID83  string `json:"id83"`
	ID84  string `json:"id84"`
	ID85  string `json:"id85"`
	ID86  string `json:"id86"`
	ID87  string `json:"id87"`
	ID88  string `json:"id88"`
	ID89  string `json:"id89"`
	ID90  string `json:"id90"`
	ID91  string `json:"id91"`
	ID92  string `json:"id92"`
	ID93  string `json:"id93"`
	ID94  string `json:"id94"`
	ID95  string `json:"id95"`
	ID96  string `json:"id96"`
	ID97  string `json:"id97"`
	ID98  string `json:"id98"`
	ID99  string `json:"id99"`
	ID100 string `json:"id100"`
}

func Unmarshal06(n int) {
	var obj Struct06
	var data Struct06Data
	for i := 0; i < n; i++ {
		json.Unmarshal(jsonValueLong, &obj)

		json.Unmarshal(obj.Data, &data)
		fmt.Print(data)
	}
}
