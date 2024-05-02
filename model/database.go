package model

import "time"

type User struct {
	ID       string `json:"ID"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	PlayerId string `json:"playerId"`
	Disabled bool   `json:"disabled"`
}

type Player struct {
	ID       string `json:"ID"`
	Name     string `json:"name"`
	Disabled bool   `json:"disabled"`
}

type WeekResult struct {
	Year        int       `json:"year"`
	Week        int       `json:"week"`
	Day         time.Time `json:"day"`
	PlayerID    string    `json:"playerID"`
	BattlesDone int       `json:"battlesDone"`
}

type Permission struct {
	PlayerID string `json:"playerID"`
	StartDay time.Time `json:"startDay"`
	EndDay time.Time `json:"endDay"`
}

type RaceConfig struct {
	StartTime time.Time `json:"startTime"`
	EndTime time.Time `json:"endtime"`
}