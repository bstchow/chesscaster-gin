package controllers

import (
	"chesscaster-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameCreationRequest struct {
	GameState      string `json:"game_state"`
	ActivityState  uint   `json:"activity_state"`
	WhitePlayerFid uint   `json:"white_player_fid"`
	BlackPlayerFid uint   `json:"black_player_fid"`
}

type GameStatusUpdateRequest struct {
	GameState     string `json:"game_state"`
	ActivityState uint   `json:"activity_state"`
}

// ActiveGamesOfUser gets all existing games for the given user
func ActiveGamesOfUser(c *gin.Context) {
	fid := c.Query("fid")

	games, games_err := models.ActiveGames(fid)
	if games_err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else {
		c.JSON(http.StatusOK, games)
	}
}

// CreateGame creates a new game.
func CreateGame(c *gin.Context) {
	req := GameCreationRequest{}

	if bind_err := c.BindJSON(&req); bind_err != nil {
		c.AbortWithError(http.StatusBadRequest, bind_err)
		return
	}

	game := models.Game{
		GameState:      req.GameState,
		ActivityState:  req.ActivityState,
		WhitePlayerFid: req.WhitePlayerFid,
		BlackPlayerFid: req.BlackPlayerFid,
	}

	err := models.CreateGame(&game)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		c.JSON(http.StatusCreated, game)
	}
}

// PatchGame patches a new game by ID.
func PatchGame(c *gin.Context) {
	var req GameStatusUpdateRequest

	if bind_err := c.BindJSON(&req); bind_err != nil {
		c.AbortWithError(http.StatusBadRequest, bind_err)
		return
	}

	id32, parseError := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id32)

	if parseError != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	game := models.Game{
		Model:         gorm.Model{ID: id},
		GameState:     req.GameState,
		ActivityState: req.ActivityState,
	}

	patch_err := models.PatchGame(&game)

	if patch_err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		c.JSON(http.StatusOK, game)
	}
}
