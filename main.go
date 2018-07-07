package main

import (
	"math/rand"

	"github.com/legendtkl/game-jam/types"
	"github.com/google/uuid"

	"github.com/loomnetwork/go-loom/plugin"
	contract "github.com/loomnetwork/go-loom/plugin/contractpb"
	"github.com/pkg/errors"
	//"strconv"
)

const (
	FogsProportion = 0.2
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
	gameMap.Creator = tx.Creator
	gameMap.Graph = tx.Graph
	gameMap.Fee = tx.Fee
	gameMap.Reward = tx.Reward
	gameMap.State = 0
	for _, pixel := range tx.Graph {
		rd := rand.Intn(len(tx.Graph) * int(1.0/FogsProportion))
		if rd < len(tx.Graph) {
			gameMap.Fogs = append(gameMap.Fogs, pixel.Point)
		}
	}
	response.Map = &gameMap
	return &response, nil
}

func (e *Game) ListGameMap(ctx contract.Context, tx *types.ListGameMapRequest) (*types.ListGameMapResponse, error) {
	response := types.ListGameMapResponse{}
	//var maps []types.GameMap

	ctx.Get(e.MapListKey(), &response)
	return &response, nil
}

func (e *Game) Challenge(ctx contract.Context, tx *types.ChallengeRequest) (*types.ChallengeResponse, error) {
	response := types.ChallengeResponse{}

	return &response, nil
}
var Contract plugin.Contract = contract.MakePluginContract(&Game{})