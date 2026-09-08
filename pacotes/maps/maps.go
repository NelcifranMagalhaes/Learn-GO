package main

import "fmt"

func main() {
	user := map[string]string{
		"name":     "Midoria",
		"lastname": "Izuku",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	game := map[string]map[string]string{
		"basic_infos": {
			"name":      "Resident Evil",
			"developer": "Capcom",
		},
		"prices": {
			"sony":     "250.25",
			"nintendo": "500",
		},
	}
	fmt.Println(game)
	delete(game, "prices")
	fmt.Println(game)
	game["popularity"] = map[string]string{
		"vote": "10",
	}
	fmt.Println(game)
}
