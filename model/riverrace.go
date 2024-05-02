package model

type Riverrace struct {
	State string `json:"state"`
	Clan  Clan   `json:"clan"`
}

type Clan struct {
	Tag          string         `json:"tag"`
	Name         string         `json:"name"`
	BadgeId      int            `json:"badgeId"`
	Fame         int            `json:"fame"`
	RepairPoints int            `json:"repairPoints"`
	Participants []Participants `json:"participants"`
	PeriodPoints int            `json:"periodPoints"`
	ClanScore    int            `json:"clanScore"`
}

type Participants struct {
	Tag            string `json:"tag"`
	Name           string `json:"name"`
	Fame           int    `json:"fame"`
	RepairPoints   int    `json:"repairPoints"`
	BoatAttacks    int    `json:"boatAttacks"`
	DecksUsed      int    `json:"decksUsed"`
	DecksUsedToday int    `json:"decksUsedToday"`
}