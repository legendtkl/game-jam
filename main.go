package main

import (
	"math/rand"
	"runtime/debug"

	"github.com/legendtkl/game-jam/types"
	"github.com/google/uuid"

	"github.com/loomnetwork/go-loom/plugin"
	contract "github.com/loomnetwork/go-loom/plugin/contractpb"
	"github.com/pkg/errors"
	//"strconv"
	"fmt"
	"time"
)

const (
	FogsProportion = 0.05
)

func main() {
	rand.Seed(99)
	plugin.Serve(Contract)
}

type Game struct {
}

func (e *Game) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "game",
		Version: "0.0.1",
	}, nil
}

func (e *Game) Init(ctx contract.Context, req *plugin.Request) error {
	return nil
}

func (e *Game) UserKey(user types.User) []byte {
	return []byte("user_" + user.Username)
}

func (e *Game) MapKey(gameMap types.GameMap) []byte {
	return []byte("map_" + gameMap.MapId)
}

func (e *Game) MapListKey() []byte {
	return []byte("maplist")
}

func (e *Game) ChallengeKey(challenge types.Challenge) []byte {
	return []byte("challenge_" + challenge.ChanllengeId)
}

func (e* Game) CreateUser(ctx contract.Context, tx *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	var user types.User
	var response types.CreateUserResponse

	if len(tx.Username) <= 0 {
		response.Code = 1000
		response.Message = "invalid username"
		return &response, errors.New(response.Message)
	}
	user.Username = tx.Username
	if ctx.Has(e.UserKey(user)) {
		response.Code = 1001
		response.Message = "username already exist"
		return &response, errors.New(response.Message)
	}
	user.Balance = 100
	ctx.Set(e.UserKey(user), &user)
	response.User = &user
	return &response, nil
}

func (e *Game) CreateMap(ctx contract.Context, tx *types.CreateGameMapRequest) (*types.CreateGameMapResponse, error) {
	var gameMap types.GameMap
	var response types.CreateGameMapResponse

	for i := 0; i < 3; i = i+1 {
		gameMap.MapId = uuid.New().String()
		if !ctx.Has(e.MapKey(gameMap)) {
			break
		}
	}
	if ctx.Has(e.MapKey(gameMap)) {
		response.Code = 2000
		response.Message = "failed to generate uuid"
		return &response, errors.New(response.Message)
	}
	var creator types.User
	creator.Username = tx.Creator.Username
	ctx.Get(e.UserKey(creator), &creator)

	gameMap.Creator = &creator
	gameMap.Graph = tx.Graph
	gameMap.Fee = tx.Fee
	gameMap.Reward = tx.Reward
	gameMap.State = 0
	var cstZone = time.FixedZone("CST", 8*3600)
	gameMap.CreateDateTime = time.Now().In(cstZone).Format("2006-01-02 15:04:05")
	for _, pixel := range tx.Graph {
		rd := rand.Intn(len(tx.Graph) * int(1.0/FogsProportion))
		if rd < len(tx.Graph) {
			gameMap.Fogs = append(gameMap.Fogs, pixel.Point)
		}
	}
	response.Map = &gameMap

	var maplist types.MapList
	ctx.Get(e.MapListKey(), &maplist)
	maplist.Maps = append(maplist.Maps, &gameMap)
	ctx.Set(e.MapListKey(), &maplist)

	return &response, nil
}

func (e *Game) ListGameMap(ctx contract.Context, tx *types.ListGameMapRequest) (*types.ListGameMapResponse, error) {
	response := types.ListGameMapResponse{}
	var maplist types.MapList

	ctx.Get(e.MapListKey(), &maplist)
	response.Maps = maplist.Maps

	return &response, nil
}

func (e *Game) Challenge(ctx contract.Context, tx *types.ChallengeRequest) (*types.ChallengeResponse, error) {
	var challenge types.Challenge
	var response types.ChallengeResponse

	for i := 0; i < 3; i = i+1 {
		challenge.ChanllengeId = uuid.New().String()
		if !ctx.Has(e.ChallengeKey(challenge)) {
			break
		}
	}
	if ctx.Has(e.ChallengeKey(challenge)) {
		response.Code = 3000
		response.Message = "failed to generate uuid"
		return &response, errors.New(response.Message)
	}

	challenge.MapId = tx.MapId
	challenge.Player = tx.Player
	ctx.Set(e.ChallengeKey(challenge), &challenge)

	var gameMap types.GameMap
	gameMap.MapId = tx.MapId
	ctx.Get(e.MapKey(gameMap), &gameMap)

	var player types.User
	player.Username = tx.Player.Username
	ctx.Get(e.UserKey(player), &player)

	if (player.Balance < gameMap.Fee ) {
		response.Code = 3001
		response.Message = "Not Enough Coin. You Poor Man!"
		return &response, errors.New(response.Message)
	}

	player.Balance -= gameMap.Fee
	ctx.Set(e.UserKey(player), &player)

	response.Code = 0
	response.ChanllengeId = challenge.ChanllengeId
	response.Message = "OK"

	return &response, nil
}
             //UploadChallenge
func (e *Game) UploadChallenge(ctx contract.Context, tx *types.ChallengeResultRequest) (*types.ChallengeResultResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
			ctx.Logger().Log(err)
			ctx.Logger().Log(string(debug.Stack()))
		}
	}()
	var challenge types.Challenge
	challenge.ChanllengeId = tx.ChanllengeId
	ctx.Get(e.ChallengeKey(challenge), &challenge)
	ctx.Logger().Log("challenge:", challenge)

	var gameMap types.GameMap
	gameMap.MapId = challenge.MapId
	ctx.Get(e.MapKey(gameMap), &gameMap)
	ctx.Logger().Log("gamemap:", gameMap)

	var player types.User
	player.Username = tx.Player.Username
	ctx.Get(e.UserKey(player), &player)
	ctx.Logger().Log("player:", player)

	var response types.ChallengeResultResponse
	if gameMap.State == 2 {
		response.Code = 3002
		response.Message = "The Map has been Challenged Successful. You CHEAT ??"
		return &response, errors.New(response.Message)
	}
	ctx.Logger().Log("line 199")

	if (tx.Result == true) {
		//改 map 状态
		gameMap.State = 2
		gameMap.Solver = &player
		ctx.Set(e.MapKey(gameMap), &gameMap)

		//加钱
		player.Balance += gameMap.Reward
		ctx.Set(e.UserKey(player), &player)
	}
	ctx.Logger().Log("line 211")
	response.Code = 0
	response.Player = &player
	response.Message = "OK"
	return &response, nil
}

var Contract plugin.Contract = contract.MakePluginContract(&Game{})