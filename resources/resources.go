package resources

var PlayerInfoIds = map[string]int32{
	// Chests
	"CurrentPosition": 2,
	"SuperMagicalChest": 12,
	"LegendaryChest": 15,
	"MagicalChest": -1,
	"EpicChest": 22,
	"GiantChest": -1,

	// Shop Offers
	"LegendaryOffer": 17,
	"EpicOffer": 18,
	"ArenaOffer": 19,

	// Player Info
	"Gold": 1,
	"CurrentDay": 16, // usefull to calculate when will be the offer
	"RecordTrophies": 6,
	"3CrownWins": 7,
	"UnlockedCards": 8,
	"FavouriteCard": 9,
	"Donations": 10,
	"ChallengeMaxWins": 20,
	"ChallengeCardsWon": 21,
}