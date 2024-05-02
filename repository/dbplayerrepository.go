package repository

import (
	"clashapiv2-api/model"
	"clashapiv2-api/service/database"
	"context"

	"github.com/google/uuid"
)

func CreateDbPlayer(ctx context.Context, player model.Player) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	sql := `INSERT INTO public.players (id, name, disabled)
			VALUES($1, $2, true)`
	ID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, sql, ID, player.Name)
	if err != nil {			
		return err
	}
	return nil
}

//func GetDbPlayers()
//func GetDbPlayerById()
func GetDbPlayerByName(ctx context.Context, name string)(model.Player, error){
	var player model.Player
	db, err := database.Connect()
	if err != nil {
		return player, err
	}
	defer db.Close()
	sql := `SELECT id, name, disabled FROM public.players WHERE name = $1`
	row := db.QueryRowContext(ctx, sql, name)
	err = row.Scan(&player.ID, &player.Name, &player.Disabled)
	if err != nil {
		return player, err
	}
	return player, nil
}
//func UpdateDbPlayer()
//func DeleteDbPlayer()