package models

type StartGameParam struct {
	TokenId  string   `json:"tokenId"`
	GameId   string   `json:"gameId"`
	Types    string   `json:"types"`
	Level    string   `json:"level"`
	Distance string   `json:"distance"`
	HorseIds []string `json:"horseIds"`
}

type CancelGameParam struct {
	TokenId string   `json:"tokenId"`
	GameIds []string `json:"gameIds"`
}

type EndGameParam struct {
	TokenId string   `json:"tokenId"`
	GameId  string   `json:"gameId"`
	Rank    []string `json:"rank"`
	Score   []string `json:"score"`
	Comp    []string `json:"comp"`
}

type MintEquipParam struct {
	To         string `json:"to"`
	EquipTypes string `json:"equip_types"`
	EquipStyle string `json:"equip_style"`
}

type Rank struct {
	HorseIds    []string `json:"horseIds"`
	TotalScores []string `json:"totalScores"`
	Accounts    []string `json:"accounts"`
	Rank        []string `json:"rank"`
}

type InitGrade struct {
	HorseIds []string `json:"horseIds"`
	Comp     []string `json:"comp"`
}

type VerifySignParam struct {
	Message   string `json:"message"`
	Address   string `json:"address"`
	Signature string `json:"signature"`
}
