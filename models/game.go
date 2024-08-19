package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Game is the main game model.
type Game struct {
	gorm.Model
	GameState      string
	ActivityState  uint
	WhitePlayerFid uint `gorm:"index"`
	BlackPlayerFid uint `gorm:"index"`
}

// ActivityState
const (
	ACTIVE    = 0
	WHITE_WIN = 1
	BLACK_WIN = 2
	DRAW      = 3
)

type GameDoesNotExistError uint

func (e GameDoesNotExistError) Error() string {
	return fmt.Sprintf("Game with ID: %d does not exist", e)
}

func ActiveGames(fid string) (games []Game, err error) {
	result := DB.Where("activity_state = ?", ACTIVE).Where("white_player_fid = ? OR black_player_fid = ?", fid, fid).Find(&games)
	if err = result.Error; err != nil {
		return nil, err
	}

	return games, nil
}

// CreateGame creates a new game.
func CreateGame(game *Game) (err error) {
	if err = DB.Create(game).Error; err != nil {
		return err
	}

	return nil
}

// PatchGame patches an existing game.
func PatchGame(game *Game) (err error) {
	tx := DB.Where("id = ?", game.ID).Limit(1).Updates(game)

	fmt.Printf("Rows affected: %d", tx.RowsAffected)

	if err = tx.Error; err != nil {
		return err
	}

	if tx.RowsAffected == 0 {
		return GameDoesNotExistError(game.ID)
	}

	return nil
}
