package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	first = []string{
		"Baby Oil", "Bad News", "Big Burps", "Bill 'Beenie-Weenie'",
		"Bob 'Stinkbug'", "Bowel Noises", "Boxelder", "Bud 'Lite'",
		"Butterbean", "Buttermilk", "Buttocks", "Chad", "Chesterfield",
		"Chewy", "Chigger", "Cinnabuns", "Cleet", "Cornbread", "Crab Meat",
		"Crapps", "Dark Skies", "Dennis Clawhammer", "Dicman", "Elphonso",
		"Fancypants", "Figgs", "Foncy", "Gootsy", "Greasy Jim", "Huckleberry",
		"Huggy", "Ignatious", "Jimbo", "Joe 'Pottin Soil'", "Johnny",
		"Lemongrass", "Lil Debil", "Longbranch", "Lunch Money", "Mergatroid",
		"Mr Peabody", "Oil-Can", "Oinks", "Old Scratch", "Ovaltine",
		"Pennywhistle", "Pitchfork Ben", "Potato Bug", "Pushmeet",
		"Rock Candy", "Schlomo", "Scratchensniff", "Scut",
		"Sid 'The Squirts'", "Skidmark", "Slaps", "Snakes", "Snoobs",
		"Snorki", "Soupcan Sam", "Spitzitout", "Squids", "Stinky",
		"Storyboard", "Sweet Tea", "TeeTee", "Wheezy Joe",
		"Winston Jazz Hands", "Worms",
	}

	last = []string{
		"Appleyard", "Bigmeat", "Bloominshine", "Boogerbottom",
		"Breedslovetrout", "Butterbaugh", "Clovenhoof", "Clutterbuck",
		"Cocktoasten", "Endicott", "Fewhairs", "Gooberdapple", "Goodensmith",
		"Goodpasture", "Guster", "Henderson", "Hooperbag", "Hoosenater",
		"Hootkins", "Jefferson", "Jenkins", "Jingley-Schmidt", "Johnson",
		"Kingfish", "Listenbee", "M'Bembo", "McFadden", "Moonshine", "Nettles",
		"Noseworthy", "Olivetti", "Outerbridge", "Overpeck", "Overturf",
		"Oxhandler", "Pealike", "Pennywhistle", "Peterson", "Pieplow",
		"Pinkerton", "Porkins", "Putney", "Quakenbush", "Rainwater",
		"Rosenthal", "Rubbins", "Sackrider", "Snuggleshine", "Splern",
		"Stevens", "Stroganoff", "Sugar-Gold", "Swackhamer", "Tippins",
		"Turnipseed", "Vinaigrette", "Walkingstick", "Wallbanger", "Weewax",
		"Weiners", "Whipkey", "Wigglesworth", "Wimplesnatch", "Winterkorn",
		"Woolysocks",
	}
)

func main() {
	var try_again string

	fmt.Print("Welcome to the Psych 'Sidekick Name Picker.'\n")
	fmt.Print("A name just like Sean would pick for Gus:\n\n")

	for {
		rand.Seed(time.Now().UnixNano())

		first_name := first[rand.Intn(len(first))]
		last_name := last[rand.Intn(len(last))]

		fmt.Print("\n\n")
		fmt.Print(first_name, last_name)
		fmt.Print("\n\n")
		fmt.Print("\n\nTry again? (Press Enter else n to quit)\n ")

		fmt.Scan(&try_again)
		if try_again == "n" {
			break
		}
	}
}
