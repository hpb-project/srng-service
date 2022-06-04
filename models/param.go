package models

import (
	"encoding/json"
)

/**
用于接收事件
*/
type Param struct {
	Token_Id      int       `json:"Token_Id"`
	Energy        int       `json:"energy"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Ids           []string  `json:"ids"`
	Types         []int     `json:"types"`
	Styles        []int     `json:"styles"`
	Prices        []float64 `json:"price"`
	GeneGrip      []string  `json:"geneGrip"`
	GeneSpeedUp   []string  `json:"geneSpeedUp"`
	GeneEndurance []string  `json:"geneEndurance"`
	GeneSpeed     []string  `json:"geneSpeed"`
	GeneSwerve    []string  `json:"geneSwerve"`
	GeneControl   []string  `json:"geneControl"`
}

type ParamDate struct {
	Year  uint `json:"year"`
	Month uint `json:"month"`
	Week  uint `json:"week"`
}

type Constant struct {
	BuyTaxRate               string `json:"buyTaxRate"`
	ModifyNameTimes          string `json:"modifyNameTimes"`
	TakeBackFromMarketCD     string `json:"takeBackFromMarketCD"`
	PutMarketCD              string `json:"putMarketCD"`
	GrowUpTime               string `json:"growUpTime"`
	TakeBackFromStudFarmCD   string `json:"takeBackFromStudFarmCD"`
	PutStudFarmCD            string `json:"putStudFarmCD"`
	MareBreedCD              string `json:"mareBreedCD"`
	StudBreedCD              string `json:"studBreedCD"`
	BreedCostType1           string `json:"breedCostType1"`
	BreedCostValue1          string `json:"breedCostValue1"`
	BreedCostType2           string `json:"breedCostType2"`
	BreedCostValue2          string `json:"breedCostValue2"`
	GeneRate                 string `json:"geneRate"`
	GenerationRate           string `json:"generationRate"`
	HighGeneWeight           string `json:"highGeneWeight"`
	MediuM_GeneWeight        string `json:"mediuM_GeneWeight"`
	LowGeneWeight            string `json:"lowGeneWeight"`
	GenerationWeightMax      string `json:"generationWeightMax"`
	GenerationWeightDecrease string `json:"generationWeightDecrease"`
	GenerationWeightMin      string `json:"generationWeightMin"`
	BreedTaxRate             string `json:"breedTaxRate"`
	TrainingValueMin         string `json:"trainingValueMin"`
	TrainingValueMax         string `json:"trainingValueMax"`
	TrainingAddValue         string `json:"trainingAddValue"`
	TrainingCD               string `json:"trainingCD"`
	TrainingReduceCD         string `json:"trainingReduceCD"`
	TrainingReduceValue      string `json:"trainingReduceValue"`
	TrainingCostType         string `json:"trainingCostType"`
	TrainingCostValue        string `json:"trainingCostValue"`
	RacecourseDayLimit       string `json:"racecourseDayLimit"`
	ApplyRaceCD              string `json:"applyRaceCD"`
	CancelRaceCD             string `json:"cancelRaceCD"`
	RaceCostType             string `json:"raceCostType"`
	RaceCostValue            string `json:"raceCostValue"`
	OwnerGainRate            string `json:"ownerGainRate"`
	EnergyCost               string `json:"energyCost"`
	EnergyRecoverCD          string `json:"energyRecoverCD"`
	RacecourseCloseCD        string `json:"racecourseCloseCD"`

	WeekRankCount      string           `json:"weekRankCount"`
	MonthRankCount     string           `json:"monthRankCount"`
	YearRankCount      string           `json:"yearRankCount"`
	MaxEnergy          string           `json:"maxEnergy"`
	InitScore          string           `json:"initScore"`
	InitGrade          string           `json:"initGrade"`
	MaxApplyCount      string           `json:"maxApplyCount"`
	TrackNumber        string           `json:"trackNumber"`
	BreCoefficient     string           `json:"breCoefficient"`
	MinDiscountOfHorse string           `json:"minDiscountOfHorse"`
	MaxRewardOfHorse   string           `json:"maxRewardOfHorse"`
	MortAmount         string           `json:"mortAmount"`
	MortToken          string           `json:"mortToken"`
	MaxSpacing         string           `json:"maxSpacing"`
	DistApplyFee       string           `json:"distApplyFee"`
	AwardToken         string           `json:"awardToken"`
	AwardAmount        string           `json:"awardAmount"`
	ExtraAwardToken    string           `json:"extraAwardToken"`
	ExtraAwardAmount   string           `json:"extraAwardAmount"`
	MaxScore           string           `json:"maxScore"`
	FeeRateOfEquip     string           `json:"feeRateOfEquip"`
	MinDiscountOfEquip string           `json:"minDiscountOfEquip"`
	MaxRewardOfEquip   string           `json:"maxRewardOfEquip"`
	MinSpacing         string           `json:"minSpacing"`
	Resting            string           `json:"resting"`
	OnSale             string           `json:"onSale"`
	Breeding           string           `json:"breeding"`
	Signing            string           `json:"signing"`
	InGame             string           `json:"inGame"`
	OnHold             string           `json:"onHold"`
	OnSell             string           `json:"onSell"`
	Equipped           string           `json:"equipped"`
	ApplyGameFee       map[string]int64 `json:"apply_game_fee"` // {"1200":5, "2400": 10,  "3600": 15 }
	AllAwardAmount     map[string]int64 `json:"allAwardAmount"`
}

type MintArena struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Time        string `json:"time" bson:"time"`
	Types       string `json:"types" bson:"types"`
	Mort_Amount string `json:"mort_amount" bson:"mort_amount"`
	Name        string `json:"name" bson:"name"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}

type CloseArena struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Time        string `json:"time" bson:"time"`
	Mort_Amount string `json:"mort_amount" bson:"mort_amount"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type UptArena struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Game_Id     string `json:"game_id" bson:"game_id"`
	Time        string `json:"time" bson:"time"`
	Count       string `json:"count" bson:"count"`
	Total_Count string `json:"total_count" bson:"total_count"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type CreateGame struct {
	Account     string `json:"account" bson:"account"`
	Game_Id     string `json:"game_id" bson:"game_id"`
	Time        string `json:"time" bson:"time"`
	Level       string `json:"level" bson:"level"`
	Race_Type   string `json:"race_type" bson:"race_type"`
	Distance    string `json:"distance" bson:"distance"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type EndGame struct {
	Game_Id     string `json:"game_id" bson:"game_id"`
	Status      string `json:"status" bson:"status"`
	Time        string `json:"time" bson:"time"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type SetHorseIntegral struct {
	Horse_Id      string `json:"horse_id" bson:"horse_id"`
	Count         string `json:"count" bson:"count"`
	Time          string `json:"time" bson:"time"`
	IntegralYear  string `json:"integral_year" bson:"integral_year"`
	IntegralMonth string `json:"integral_month" bson:"integral_month"`
	IntegralWeek  string `json:"integral_week" bson:"integral_week"`
	Event_Index   uint   `json:"event_index" bson:"event_index"`
	Tx_Hash       string `json:"tx_hash" bson:"tx_hash"`
	Event_Name    string `json:"event_name" bson:"event_name"`
}
type SetHorseGradeSc struct {
	Horse_Id    string `json:"horse_id" bson:"horse_id"`
	Count       string `json:"count" bson:"count"`
	Mark        string `json:"mark" bson:"mark"`
	Time        string `json:"time" bson:"time"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type SetHorseGrade struct {
	Horse_Id    string `json:"horse_id" bson:"horse_id"`
	Grade       string `json:"grade" bson:"grade"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type EndGameOfHorse struct {
	Horse_Id    string `json:"horse_id" bson:"horse_id"`
	Status      string `json:"status" bson:"status"`
	Time        string `json:"time" bson:"time"`
	Arena_Id    string `json:"arena_id" bson:"arena_id"`
	Game_Id     string `json:"game_id" bson:"game_id"`
	Race_Type   string `json:"race_type" bson:"race_type"`
	Distance    string `json:"distance" bson:"distance"`
	Energy      string `json:"energy" bson:"energy"`
	Race_Count  string `json:"race_count" bson:"race_count"`
	Win_Count   string `json:"win_count" bson:"win_count"`
	Grade       string `json:"grade" bson:"grade"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type StartGame struct {
	Horse_Ids   []string `json:"horse_ids" bson:"horse_ids"`
	Status      string   `json:"status" bson:"status"`
	Time        string   `json:"time" bson:"time"`
	RaceId      string   `json:"race_id" bson:"race_id"`
	Event_Index uint     `json:"event_index" bson:"event_index"`
	Tx_Hash     string   `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string   `json:"event_name" bson:"event_name"`
}
type StartGameOfRace struct {
	Game_Id     string `json:"game_id" bson:"game_id"`
	Status      string `json:"status" bson:"status"`
	Time        string `json:"time" bson:"time"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type ApplyGame struct {
	Account     string `json:"account" bson:"account"`
	Horse_Id    string `json:"horse_id" bson:"horse_id"`
	Arena_Id    string `json:"arena_id" bson:"arena_id"`
	Race_Type   string `json:"race_type" bson:"race_type"`
	Distance    string `json:"distance" bson:"distance"`
	Time        string `json:"time" bson:"time"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type CancelApplyGame struct {
	Account     string `json:"account" bson:"account"`
	Horse_Id    string `json:"horse_id" bson:"horse_id"`
	Status      string `json:"status" bson:"status"`
	Time        string `json:"time" bson:"time"`
	Arena_Id    string `json:"arena_id" bson:"arena_id"`
	Game_Id     string `json:"game_id" bson:"game_id"`
	Race_Type   string `json:"race_type" bson:"race_type"`
	Distance    string `json:"distance" bson:"distance"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type CancelGame struct {
	Game_Id     []string `json:"game_id" bson:"game_id"`
	Status      string   `json:"status" bson:"status"`
	Time        string   `json:"time" bson:"time"`
	Event_Index uint     `json:"event_index" bson:"event_index"`
	Tx_Hash     string   `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string   `json:"event_name" bson:"event_name"`
}
type CancelGameOfHorse struct {
	Horse_Ids   []string `json:"horse_ids" bson:"horse_ids"`
	Status      string   `json:"status" bson:"status"`
	Time        string   `json:"time" bson:"time"`
	Event_Index uint     `json:"event_index" bson:"event_index"`
	Tx_Hash     string   `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string   `json:"event_name" bson:"event_name"`
}

type CancelGameOfArena struct {
	Account     string   `json:"account" bson:"account"`
	Token_Id    string   `json:"token_id" bson:"token_id"`
	Game_Id     []string `json:"game_id" bson:"game_id"`
	Count       string   `json:"count" bson:"count"`
	Total_Count string   `json:"total_count" bson:"total_count"`
	Time        string   `json:"time" bson:"time"`
	Event_Index uint     `json:"event_index" bson:"event_index"`
	Tx_Hash     string   `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string   `json:"event_name" bson:"event_name"`
}

type StopGame struct {
	Owner       string `json:"owner" bson:"owner"`
	Game_Id     string `json:"game_id" bson:"game_id"`
	Status      string `json:"status" bson:"status"`
	Time        string `json:"time" bson:"time"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type MintEquip struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Time        string `json:"time" bson:"time"`
	To          string `json:"to" bson:"to"`
	Types       string `json:"types" bson:"types"`
	Style       string `json:"style" bson:"style"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type Sell struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Kind        string `json:"kind" bson:"kind"`
	Coin        string `json:"coin" bson:"coin"`
	Price       string `json:"price" bson:"price"`
	Time        string `json:"time" bson:"time"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type CancelSell struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Status      string `json:"status" bson:"status"`
	Time        string `json:"time" bson:"time"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type Buy struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type UnloadEquip struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type UnloadEquipOfHorse struct {
	Account     string `json:"account" bson:"account"`
	Horse_Id    string `json:"horse_id" bson:"horse_id"`
	Types       string `json:"types" bson:"types"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type Breeding struct {
	Account         string `json:"account" bson:"account"`
	Token_Id        string `json:"token_id" bson:"token_id"`
	StallId         string `json:"stall_id" bson:"stall_id"`
	NewHorse_Id     string `json:"new_horse_id" bson:"new_horse_id"`
	Name            string `json:"name" bson:"name"`
	GenerationScore string `json:"generation_score" bson:"generation_score"`
	Gender          string `json:"gender" bson:"gender"`
	IntegralTime    string `json:"integral_time" bson:"integral_time"`
	EnergyTime      string `json:"energy_time" bson:"energy_time"`
	Status          string `json:"status" bson:"status"`
	Event_Index     uint   `json:"event_index" bson:"event_index"`
	Tx_Hash         string `json:"tx_hash" bson:"tx_hash"`
	Event_Name      string `json:"event_name" bson:"event_name"`
}
type Breeding1 struct {
	NewHorse_Id string `json:"new_horse_id" bson:"new_horse_id"`
	Time        string `json:"time" bson:"time"`
	Color       string `json:"color" bson:"color"`
	Gene        string `json:"gene" bson:"gene"`
	M_Gene      string `json:"m_gene" bson:"m_gene"`
	S_Gene      string `json:"s_gene" bson:"s_gene"`
	Tra_Value   string `json:"tra_value" bson:"tra_value"`
	Energy      string `json:"energy" bson:"energy"`
	Grade       string `json:"grade" bson:"grade"`
	Integral    string `json:"integral" bson:"integral"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type BreedingOfHorse struct {
	Horse_Id    string `json:"horse_id" bson:"horse_id"`
	Time        string `json:"time" bson:"time"`
	Count       string `json:"count" bson:"count"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type TrainingHorses struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Value       string `json:"value" bson:"value"`
	Time        string `json:"time" bson:"time"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type UptHorseName struct {
	Token_Id    string `json:"token_id" bson:"token_id"`
	Name        string `json:"name" bson:"name"`
	Count       string `json:"count" bson:"count"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type HorseDeco struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Types       string `json:"types" bson:"types"`
	Equip_Id    string `json:"equip_id" bson:"equip_id"`
	Time        string `json:"time" bson:"time"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type HorseDecoOfEquip struct {
	Equip_Id    string `json:"equip_id" bson:"equip_id"`
	Status      string `json:"status" bson:"status"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type SetHorseGene struct {
	Token_Id     string `json:"token_id" bson:"token_id"`
	Grip_Gene    string `json:"grip_gene" bson:"grip_gene"`
	Acc_Gene     string `json:"acc_gene" bson:"acc_gene"`
	End_Gene     string `json:"end_gene" bson:"end_gene"`
	Speed_Gene   string `json:"speed_gene" bson:"speed_gene"`
	Turn_To_Gene string `json:"turn_to_gene" bson:"turn_to_gene"`
	Control_Gene string `json:"control_gene" bson:"control_gene"`
	Event_Index  uint   `json:"event_index" bson:"event_index"`
	Tx_Hash      string `json:"tx_hash" bson:"tx_hash"`
	Event_Name   string `json:"event_name" bson:"event_name"`
}
type SellHorse struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Coin        string `json:"coin" bson:"coin"`
	Price       string `json:"price" bson:"price"`
	Time        string `json:"time" bson:"time"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type SireHorse struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Coin        string `json:"coin" bson:"coin"`
	Price       string `json:"price" bson:"price"`
	Time        string `json:"time" bson:"time"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type CancelSellHorse struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Time        string `json:"time" bson:"time"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type CancelSireHorse struct {
	Account     string `json:"account" bson:"account"`
	Token_Id    string `json:"token_id" bson:"token_id"`
	Time        string `json:"time" bson:"time"`
	Status      string `json:"status" bson:"status"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}
type BuyHorse struct {
	Account           string `json:"account" bson:"account"`
	Token_Id          string `json:"token_id" bson:"token_id"`
	Coin              string `json:"coin" bson:"coin"`
	Price             string `json:"price" bson:"price"`
	Status            string `json:"status" bson:"status"`
	Modify_Name_Times string `json:"modify_name_times" bson:"modify_name_times"`
	Event_Index       uint   `json:"event_index" bson:"event_index"`
	Tx_Hash           string `json:"tx_hash" bson:"tx_hash"`
	Event_Name        string `json:"event_name" bson:"event_name"`
}

type InitHorseGrade struct {
	Horse_Id    string `json:"horse_id" bson:"horse_id"`
	Grade       string `json:"grade" bson:"grade"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}

type ArenaTransfer struct {
	TxFrom      string `json:"from" bson:"from"`
	TxTo        string `json:"to" bson:"to"`
	ArenaId     string `json:"arena_id" bson:"arena_id"`
	Time        string `json:"time" bson:"time"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}

type EquipTransfer struct {
	TxFrom      string `json:"from" bson:"from"`
	TxTo        string `json:"to" bson:"to"`
	EquipId     string `json:"equip_id" bson:"equip_id"`
	Time        string `json:"time" bson:"time"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}

type HorseTransfer struct {
	TxFrom      string `json:"from" bson:"from"`
	TxTo        string `json:"to" bson:"to"`
	HorseId     string `json:"horse_id" bson:"horse_id"`
	Time        string `json:"time" bson:"time"`
	Event_Index uint   `json:"event_index" bson:"event_index"`
	Tx_Hash     string `json:"tx_hash" bson:"tx_hash"`
	Event_Name  string `json:"event_name" bson:"event_name"`
}

func Marshal(param interface{}) []byte {
	byteInfo, _ := json.Marshal(param)
	return byteInfo
}
