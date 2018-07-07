package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/auth"
	"github.com/loomnetwork/go-loom/client"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ed25519"
	"github.com/legendtkl/game-jam/types"
)

var writeURI = fmt.Sprintf("http://%s:%d/rpc", "localhost", 46658)
var readURI = fmt.Sprintf("http://%s:%d/query", "localhost", 46658)

func getPrivKey(privKeyFile string) ([]byte, error) {
	return ioutil.ReadFile(privKeyFile)
}

func main() {
	var privFile string

	rpcClient := client.NewDAppChainRPCClient("default", writeURI, readURI)

	contractAddr, err := loom.LocalAddressFromHexString("0xe288d6eec7150D6a22FDE33F0AA2d81E06591C4d")
	if err != nil {
		log.Fatalf("Cannot generate contract address: %v", err)
	}
	contract := client.NewContract(rpcClient, contractAddr)

	var userName string
	createUser := &cobra.Command{
		Use:   "create-user",
		Short: "Create User",
		RunE: func(cmd *cobra.Command, args []string) error {
			privKey, err := getPrivKey(privFile)
			if err != nil {
				log.Fatal(err)
			}
			var result types.CreateUserResponse
			params := &types.CreateUserRequest{
				Username:	userName,
			}
			signer := auth.NewEd25519Signer(privKey)
			resp, err := contract.Call("CreateUser", params, signer, &result)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(resp)
			fmt.Println(result)

			return nil
		},
	}
	createUser.Flags().StringVarP(&privFile, "key", "k", "", "private key file")
	createUser.Flags().StringVarP(&userName, "username", "u", "", "user name")


	var creator string
	var graph string
	var reward int64
	var fee	int64
	//CreateGameMap
	createMapCmd := &cobra.Command{
		Use:   "create-map",
		Short: "send a transaction",
		RunE: func(cmd *cobra.Command, args []string) error {
			privKey, err := getPrivKey(privFile)
			if err != nil {
				log.Fatal(err)
			}
			var result types.CreateGameMapResponse
			params := &types.CreateGameMapRequest{
				Creator: &types.User{creator, 1000},
				Graph: 	nil,
				Reward: reward,
				Fee:	fee,
			}
			signer := auth.NewEd25519Signer(privKey)
			resp, err := contract.Call("CreateMap", params, signer, &result)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(resp)
			fmt.Println(result)

			return nil
		},
	}
	createMapCmd.Flags().StringVarP(&privFile, "key", "k", "", "private key file")
	createMapCmd.Flags().StringVarP(&creator, "creator", "c", "", "private key file")
	createMapCmd.Flags().StringVarP(&graph, "graph", "g", "", "private key file")
	createMapCmd.Flags().Int64VarP(&reward, "reward", "r", 0, "reward")
	createMapCmd.Flags().Int64VarP(&fee, "fee", "f", 0, "fee")


	listGameMapCmd := &cobra.Command{
		Use:   "list-game-map",
		Short: "list game map",
		RunE: func(cmd *cobra.Command, args []string) error {
			privKey, err := getPrivKey(privFile)
			if err != nil {
				log.Fatal(err)
			}
			signer := auth.NewEd25519Signer(privKey)
			var result types.ListGameMapResponse
			params := &types.ListGameMapRequest{}
			resp, err := contract.Call("ListGameMap", params, signer, &result)
			fmt.Println(resp)
			fmt.Println(result)
			return nil
		},
	}
	listGameMapCmd.Flags().StringVarP(&privFile, "key", "k", "", "private key file")

	var mapId string
	challengeCmd := &cobra.Command{
		Use:   "challenge",
		Short: "challenge the game",
		RunE: func(cmd *cobra.Command, args []string) error {
			privKey, err := getPrivKey(privFile)
			if err != nil {
				log.Fatal(err)
			}
			signer := auth.NewEd25519Signer(privKey)
			var result types.ChallengeResponse
			params := &types.ChallengeRequest{
				Player: &types.User{Username: userName},
				MapId:  mapId,
			}
			resp, err := contract.Call("Challenge", params, signer, &result)
			fmt.Printf("response: %v", resp)
			fmt.Println(result)
			return nil
		},
	}
	challengeCmd.Flags().StringVarP(&privFile, "key", "k", "", "private key file")
	challengeCmd.Flags().StringVarP(&userName, "username", "u", "loom", "user")
	challengeCmd.Flags().StringVarP(&mapId, "mapid", "m", "loom", "user")

	var challengeId string
	var result bool
	uploadChallengeCmd := &cobra.Command{
		Use:   "upload-challenge",
		Short: "upload challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			privKey, err := getPrivKey(privFile)
			if err != nil {
				log.Fatal(err)
			}
			signer := auth.NewEd25519Signer(privKey)
			var result1 types.ChallengeResultResponse
			params := &types.ChallengeResultRequest{
				Player: &types.User{Username: userName},
				ChanllengeId: challengeId,
				Result: result,
			}
			resp, err := contract.Call("UploadChallenge", params, signer, &result1)
			fmt.Println(resp)
			fmt.Println(result1)
			return nil
		},
	}
	uploadChallengeCmd.Flags().StringVarP(&privFile, "key", "k", "", "private key file")
	uploadChallengeCmd.Flags().StringVarP(&userName, "username", "u", "", "private key file")
	uploadChallengeCmd.Flags().BoolVarP(&result, "result", "r", false, "")
	uploadChallengeCmd.Flags().StringVarP(&challengeId, "challengeid", "c", "", "")

	keygenCmd := &cobra.Command{
		Use:   "genkey",
		Short: "generate a public and private key pair",
		RunE: func(cmd *cobra.Command, args []string) error {

			_, priv, err := ed25519.GenerateKey(nil)
			if err != nil {
				log.Fatalf("Error generating key pair: %v", err)
			}
			if err := ioutil.WriteFile(privFile, priv, 0664); err != nil {
				log.Fatalf("Unable to write private key: %v", err)
			}
			return nil
		},
	}
	keygenCmd.Flags().StringVarP(&privFile, "key", "k", "", "private key file")

	rootCmd := &cobra.Command{
		Use:   "loomdicecli",
		Short: "LoomDice cli tool",
	}
	rootCmd.AddCommand(keygenCmd)
	rootCmd.AddCommand(createUser)
	rootCmd.AddCommand(createMapCmd)
	rootCmd.AddCommand(listGameMapCmd)
	rootCmd.AddCommand(challengeCmd)
	rootCmd.AddCommand(uploadChallengeCmd)
	rootCmd.Execute()
}
