package packets

import (
	"encoding/binary"
	"bytes"

	"github.com/segura2010/cr-go/utils"
	"github.com/segura2010/cr-go/packets/components"
	"github.com/segura2010/cr-go/resources"
)

type ClientVisitHome struct {
	Hi int32
	Lo int32
}

func NewClientVisitHomeFromBytes(buff []byte) (ClientVisitHome){
	o := ClientVisitHome{}

	buf := bytes.NewReader(buff)

	binary.Read(buf, binary.BigEndian, &o.Hi)
	binary.Read(buf, binary.BigEndian, &o.Lo)

	return o
}

func (o *ClientVisitHome) Bytes() ([]byte){
	// It creates the message bytes ready to be sent
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, o.Hi)
	binary.Write(buf, binary.BigEndian, o.Lo)

	return buf.Bytes()
}

// It represents the ServerVisitHome response for the ClientVisitHome request
type ServerVisitHome struct {
	// actual player's deck
	Deck [8]components.Card
	// player identifier
	Hi int32
	Lo int32
	// information about the seasons (it is optional, depending on the player trophies)
	Seasons []components.Season
	Username string
	Trophies int32
	ChestCycle components.ChestCycle
	ShopOffers components.ShopOffers
	Gold int32
	FavouriteCard int32
	Stats components.Stats
	Gems int32
	Experience int32
	Level int32
	HasClan bool
	Clan components.Clan
	// total played games
	Games int32
	TournamentGames int32
	Wins int32
	Losses int32
	CurrentStreak int32
}

func NewServerVisitHomeFromBytes(buff []byte) (ServerVisitHome){
	o := ServerVisitHome{}

	var buf *bytes.Reader
	var tmp int32
	var btmp byte
	var isPresent byte

	buf = bytes.NewReader(buff)

	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	binary.Read(buf, binary.BigEndian, &btmp)
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)

	// read actual deck cards
	for i:=0;i<8;i++{
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Deck[i].Id)
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Deck[i].Level)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Deck[i].Count)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	}

	binary.Read(buf, binary.BigEndian, &o.Hi)
	binary.Read(buf, binary.BigEndian, &o.Lo)

	// HomeUnknownSeason optional (read it if present, continue if not)
	binary.Read(buf, binary.BigEndian, &isPresent)
	if isPresent > 0{
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	}

	// HomeSeason[]
	// read length, then read each season info
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	o.Seasons = make([]components.Season, tmp)
	for i:=0; i<int(tmp); i++{
		var ltmp int32
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Seasons[i].Id)
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Seasons[i].HigherTrophies)
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Seasons[i].Trophies)
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Seasons[i].RankingPosition)
		utils.ReadRrsInt32(buf, binary.BigEndian, &ltmp)
	}

	// more unknowns..
	for i:=0;i<7;i++{
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	}

	utils.ReadString(buf, binary.BigEndian, &o.Username)

	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)

	utils.ReadRrsInt32(buf, binary.BigEndian, &o.Trophies)

	for i:=0;i<15;i++{
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	}

	// chestCycle[] and other user info..
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	for i:=0;i<int(tmp);i++{
		var itype int32
		var id int32
		var value int32

		utils.ReadRrsInt32(buf, binary.BigEndian, &itype)
		//fmt.Printf("\ntype: %d", itype)
		utils.ReadRrsInt32(buf, binary.BigEndian, &id)
		//fmt.Printf("\nid: %d", id)
		utils.ReadRrsInt32(buf, binary.BigEndian, &value)
		//fmt.Printf("\nvalue: %d", value)

		// is chestcycle/shop offer/playergold/favouritecard
		if itype == 5{
			if id == resources.PlayerInfoIds["CurrentPosition"]{
				// chests
				o.ChestCycle.CurrentPosition = value
			}else if id == resources.PlayerInfoIds["EpicChest"]{
				o.ChestCycle.EpicPos = value
			}else if id == resources.PlayerInfoIds["SuperMagicalChest"]{
				o.ChestCycle.SuperMagicalPos = value
			}else if id == resources.PlayerInfoIds["LegendaryChest"]{
				o.ChestCycle.LegendaryPos = value
			}else if id == resources.PlayerInfoIds["CurrentDay"]{
				// shop offers
				o.ShopOffers.CurrentDay = value
			}else if id == resources.PlayerInfoIds["LegendaryOffer"]{
				o.ShopOffers.Legendary = value
			}else if id == resources.PlayerInfoIds["EpicOffer"]{
				o.ShopOffers.Epic = value
			}else if id == resources.PlayerInfoIds["ArenaOffer"]{
				o.ShopOffers.Arena = value
			}else if id == resources.PlayerInfoIds["Gold"]{
				o.Gold = value
			}else if id == resources.PlayerInfoIds["FavouriteCard"]{
				o.FavouriteCard = value
			}
		}
	}

	// decks4[],decks3[],decks2[] ??
	for d:=0;d<3;d++{
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		for i:=0; i<int(tmp); i++{
			var ltmp int32
			utils.ReadRrsInt32(buf, binary.BigEndian, &ltmp)
			utils.ReadRrsInt32(buf, binary.BigEndian, &ltmp)
			utils.ReadRrsInt32(buf, binary.BigEndian, &ltmp)
		}
	}

	// stats[]
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	for i:=0; i<int(tmp); i++{
		var itype int32
		var id int32
		var value int32
		utils.ReadRrsInt32(buf, binary.BigEndian, &itype)
		//fmt.Printf("\ntype: %d", itype)
		utils.ReadRrsInt32(buf, binary.BigEndian, &id)
		//fmt.Printf("\nid: %d", id)
		utils.ReadRrsInt32(buf, binary.BigEndian, &value)
		//fmt.Printf("\nvalue: %d", value)

		if itype == 5{
			if id == resources.PlayerInfoIds["FavouriteCard"]{
				o.FavouriteCard = value
			}else if id == resources.PlayerInfoIds["RecordTrophies"]{
				o.Stats.RecordTrophies = value
			}else if id == resources.PlayerInfoIds["3CrownWins"]{
				o.Stats.CrownWins3 = value
			}else if id == resources.PlayerInfoIds["Donations"]{
				o.Stats.Donations = value
			}else if id == resources.PlayerInfoIds["ChallengeMaxWins"]{
				o.Stats.ChallengeMaxWins = value
			}else if id == resources.PlayerInfoIds["ChallengeCardsWon"]{
				o.Stats.ChallengeCardsWon = value
			}else if id == resources.PlayerInfoIds["UnlockedCards"]{
				o.Stats.UnlockedCards = value
			}
		}
	}

	// cardsUsed[] (useless?)
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	for i:=0; i<int(tmp); i++{
		var itype int32
		var id int32
		var value int32
		utils.ReadRrsInt32(buf, binary.BigEndian, &itype)
		//fmt.Printf("\ntype: %d", itype)
		utils.ReadRrsInt32(buf, binary.BigEndian, &id)
		//fmt.Printf("\nid: %d", id)
		utils.ReadRrsInt32(buf, binary.BigEndian, &value)
		//fmt.Printf("\nvalue: %d", value)
	}

	// unknownProfileComponent (optional)
	binary.Read(buf, binary.BigEndian, &isPresent)
	if isPresent > 0{
		var b byte
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
		binary.Read(buf, binary.BigEndian, &b)
	}

	// more info..
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp) // unknown
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.Gems)
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp) // freegems
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.Experience)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.Level)

	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp) // unknown

	// clan info
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	o.HasClan = (tmp > 1)
	if tmp > 1{
		// player has clan
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Clan.Hi)
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Clan.Lo)
		utils.ReadString(buf, binary.BigEndian, &o.Clan.Name)
		utils.ReadRrsInt32(buf, binary.BigEndian, &o.Clan.Badge)
		binary.Read(buf, binary.BigEndian, &o.Clan.Role)
	}
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.Games)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.TournamentGames)
	
	utils.ReadRrsInt32(buf, binary.BigEndian, &tmp)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.Wins)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.Losses)
	utils.ReadRrsInt32(buf, binary.BigEndian, &o.CurrentStreak)

	return o
}

