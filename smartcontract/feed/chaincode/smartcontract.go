package chaincode

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Feed struct {
	FeedID	uint	`json:"feedid"`
	Name		string	`json:"name"`
	Dna			uint	`json:"dna"`
}

const dnaDigits uint = 14

// Define key names for options
const feedsCountKey = "feedsCount"

func (s *SmartContract) Initialize(ctx contractapi.TransactionContextInterface) (bool, error) {

	// Check minter authorization - this sample assumes Org1 is the central banker with privilege to intitialize contract
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return false, fmt.Errorf("failed to get MSPID: %v", err)
	}
	// only Org2MSP members can call
	if clientMSPID != "Org2MSP" {
		return false, fmt.Errorf("client is not authorized to initialize contract")
	}

	// Check contract options are not already set, client is not authorized to change them once intitialized
	feedsCount, err := ctx.GetStub().GetState(feedsCountKey)
	
	if err != nil {
		return false, fmt.Errorf("failed to get feedscount: %v", err)
	}
	if feedsCount != nil {
		return false, fmt.Errorf("contract options are already set, client is not authorized to change them")
	}

	// Initialize FeedsCountKey to zero(0)
	err = ctx.GetStub().PutState(feedsCountKey, []byte(strconv.Itoa(0)))
	if err != nil {
		return false, fmt.Errorf("failed to set token count: %v", err)
	}

	return true, nil
}

// create a new feed API
func (s *SmartContract) CreateRandomFeed(ctx contractapi.TransactionContextInterface, name string) error{

	dna:=s._generateRandomDna(name)
	err := s._createFeeds(ctx, name, dna)
	if err != nil {
		return fmt.Errorf("failed to createFeeds: %v", err)
	}
	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) _createFeeds(ctx contractapi.TransactionContextInterface, name string, dna uint) error {

	//  make feedid
	feedsCountBytes, err := s._getState(ctx, feedsCountKey)	
	if err != nil {
		return fmt.Errorf("failed to get feedsCount: %v", err)
	}
	feedIdINT,_ := strconv.Atoi(string(feedsCountBytes))
	feedId := uint(feedIdINT)
	
	// overwriting original fungus with new fungus
	feed := Feed{
		FeedID:  feedId,
		Name:      name,
		Dna:       dna,
	}
	feedJSON, err := json.Marshal(feed)
	if err != nil {
		return fmt.Errorf("failed to marshal feed: %v", err)
	}
	
	// PutState feedId
	err = ctx.GetStub().PutState(strconv.Itoa(int(feedId)), feedJSON)
	if err != nil {
		return fmt.Errorf("failed to put feed state: %v", err)
	}

	// PutState feedsCount++
	feedId += 1
	err = ctx.GetStub().PutState(feedsCountKey, []byte(strconv.Itoa(int(feedId))))
	if err != nil {
		return fmt.Errorf("failed to put feedsCount state: %v", err)
	}

	return nil
}


// generate random dna func
func (S *SmartContract) _generateRandomDna(name string) uint {
	nowTime := time.Now()
	unixTime := nowTime.Unix()
	data := strconv.Itoa(int(unixTime)) + name
	hash := sha256.New()
	hash.Write([]byte(data))
	dnaHash := uint(binary.BigEndian.Uint64(hash.Sum(nil)))

	// make 14digits dna
	dna := dnaHash % uint(math.Pow(10, float64(dnaDigits)))
	dna = dna - dna%100

	return dna
}

func (S *SmartContract) GetFeed(ctx contractapi.TransactionContextInterface, feedId uint) (*Feed, error) {
	feedBytes, err := S._getState(ctx, strconv.Itoa(int(feedId)))
	if err != nil {
		return nil, fmt.Errorf("failed to get feedinfo: %v", err)
	}

	var feed Feed
	err = json.Unmarshal(feedBytes, &feed)
	if err != nil {
		return nil, fmt.Errorf("failed to Unmarshal feed: %v", err)
	}

	return &feed, nil
}

// todo  delFeed(id)