package service

import (
	"clashapiv2-api/model"
	"clashapiv2-api/repository"
	"context"
)

//Comprovar l'existencia bidireccional entre clash i la bbdd


func SetDbPlayersFromClash(ctx context.Context)error{
	members, err := GetMembersFromClash()
	if err != nil {
		return err
	}
	for _, member := range members.Items {
		player,err  := repository.GetDbPlayerByName(ctx, member.Name)
		if err != nil {
			return err
		}
		if player.ID != "" {
			continue
		}
		player = model.Player{
			ID: "",
			Name: member.Name,
			Disabled: false,
		}
		err = repository.CreateDbPlayer(ctx, player)
		if err != nil {
			return err
		}
	}
	return nil
}