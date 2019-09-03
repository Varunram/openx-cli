package main

import (
	"encoding/json"
	// "log"

	"github.com/stellar/go/protocols/horizon"

	erpc "github.com/Varunram/essentials/rpc"
	utils "github.com/Varunram/essentials/utils"
	database "github.com/YaleOpenLab/openx/database"

	solar "github.com/YaleOpenLab/opensolar/core"
)

func PingRpc() error {
	// make a curl request out to lcoalhost and get the ping response
	data, err := erpc.GetRequest(ApiUrl + "/ping")
	if err != nil {
		return err
	}
	var x erpc.StatusResponse
	// now data is in byte, we need the other structure now
	err = json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	// the result would be the status of the platform
	codeString, _ := utils.ToString(x.Code)
	ColorOutput("PLATFORM STATUS: "+codeString, GreenColor)
	return nil
}

func RetrieveProject(index int) ([]solar.Project, error) {
	// retrieve project at a particular stage
	var x []solar.Project
	indexString, _ := utils.ToString(index)
	data, err := erpc.GetRequest(ApiUrl + "/projects?index=" + indexString)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func GetBalances(username string) ([]horizon.Balance, error) {
	// get the balance from the balances API
	var x []horizon.Balance
	data, err := erpc.GetRequest(ApiUrl + "/user/balances?" + "username=" + username + "&token=" + Token)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func GetXLMBalance(username string) (int, error) {
	// get the balance from the balances API
	var x int
	data, err := erpc.GetRequest(ApiUrl + "/user/balance/xlm?" + "username=" + username + "&token=" + Token)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func GetAssetBalance(username string, asset string) (int, error) {
	// get the balance from the balances API
	var x int
	data, err := erpc.GetRequest(ApiUrl + "/user/balance/asset?" + "username=" + username + "&token=" + Token + "&asset=" + asset)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func GetStableCoin(username string, amount string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/stablecoin/get?" + "seedpwd=" + LocalSeedPwd + "&amount=" +
		amount + "&username=" + username + "&token=" + Token)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func GetIpfsHash(username string, hashString string) (string, error) {
	var x string
	data, err := erpc.GetRequest(ApiUrl + "/ipfs/hash?" + "string=" + hashString +
		"&username=" + username + "&token=" + Token)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func InvestInProject(projIndex string, amount string, username string, seedpwd string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/investor/invest?" + "username=" + username + "&token=" + Token +
		"&seedpwd=" + seedpwd + "&projIndex=" + projIndex + "&amount=" + amount)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func VoteTowardsProject(projIndex string, amount string, username string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/investor/vote?" + "username=" + username + "&token=" + Token +
		"&projIndex=" + projIndex + "&votes=" + amount)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func AuthKyc(userIndex string, username string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/user/kyc?" + "username=" + username + "&token=" + Token +
		"&userIndex=" + userIndex)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func Payback(projIndex string, seedpwd string, username string, assetName string,
	amount string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/recipient/payback?" + "username=" + username + "&token=" + Token +
		"&projIndex=" + projIndex + "&seedpwd=" + seedpwd + "&amount=" + amount + "&assetName=" + assetName +
		"&platformPublicKey=" + PlatformPublicKey)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func UnlockOpenSolar(username string, seedpwd string, projIndex string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	body := ApiUrl + "/recipient/unlock/opensolar?" + "username=" + username + "&token=" + Token +
		"&projIndex=" + projIndex + "&seedpwd=" + seedpwd

	data, err := erpc.GetRequest(body)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func FinalizeProject(username string, projIndex string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/recipient/finalize?" + "username=" + username + "&token=" + Token +
		"&projIndex=" + projIndex)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func OriginateProject(username string, projIndex string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/recipient/originate?" + "username=" + username + "&token=" + Token +
		"&projIndex=" + projIndex)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func GetStage1Contracts(username string) ([]solar.Project, error) {
	var x []solar.Project
	data, err := erpc.GetRequest(ApiUrl + "/entity/stage1?" + "username=" + username + "&token=" + Token)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func GetStage0Contracts(username string) ([]solar.Project, error) {
	var x []solar.Project
	data, err := erpc.GetRequest(ApiUrl + "/entity/stage0?" + "username=" + username + "&token=" + Token)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func GetStage2Contracts(username string) ([]solar.Project, error) {
	var x []solar.Project
	data, err := erpc.GetRequest(ApiUrl + "/entity/stage2?" + "username=" + username + "&token=" + Token)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func AddCollateral(username string, collateral string, amount string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/entity/addcollateral?" + "username=" + username + "&token=" + Token +
		"&collateral=" + collateral + "&amount=" + amount)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func CreateAssetInv(username string, assetName string, pubkey string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/investor/localasset?" + "username=" + username + "&token=" + Token +
		"&assetName=" + assetName)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func SendLocalAsset(username string, seedpwd string, assetName string,
	destination string, amount string) (string, error) {
	var x string

	data, err := erpc.GetRequest(ApiUrl + "/investor/sendlocalasset?" + "username=" + username + "&token=" + Token +
		"&assetName=" + assetName + "&destination=" + destination + "&amount=" + amount + "&seedpwd=" + seedpwd)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func SendXLM(username string, seedpwd string, destination string, amount string) (string, error) {
	var x string
	data, err := erpc.GetRequest(ApiUrl + "/user/sendxlm?" + "username=" + username + "&token=" + Token +
		"&destination=" + destination + "&amount=" + amount + "&seedpwd=" + seedpwd)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func NotKycView(username string) ([]database.User, error) {
	var x []database.User
	data, err := erpc.GetRequest(ApiUrl + "/user/notkycview?" + "username=" + username + "&token=" + Token)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func KycView(username string) ([]database.User, error) {
	var x []database.User
	data, err := erpc.GetRequest(ApiUrl + "/user/kycview?" + "username=" + username + "&token=" + Token)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func AskXLM(username string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/user/askxlm?" + "username=" + username + "&token=" + Token)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func TrustAsset(username string, assetName string, issuerPubkey string,
	limit string, seedpwd string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/user/trustasset?" + "username=" + username + "&token=" + Token +
		"&assetCode=" + assetName + "&assetIssuer=" + issuerPubkey + "&limit=" + limit + "&seedpwd=" + seedpwd)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func GetTrustLimit(username string, assetName string) (string, error) {
	var x string
	data, err := erpc.GetRequest(ApiUrl + "/recipient/trustlimit?" + "username=" + username + "&token=" +
		Token + "&assetName=" + assetName)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func InvestInOpzoneCBond(projIndex string, amount string, username string, seedpwd string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/constructionbond/invest?" + "username=" + username + "&token=" + Token +
		"&seedpwd=" + seedpwd + "&projIndex=" + projIndex + "&amount=" + amount)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func InvestInLivingUnitCoop(projIndex string, amount string, username string, seedpwd string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	data, err := erpc.GetRequest(ApiUrl + "/livingunitcoop/invest?" + "username=" + username + "&token=" + Token +
		"&seedpwd=" + seedpwd + "&projIndex=" + projIndex + "&amount=" + amount)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func UnlockCBond(username string, seedpwd string, projIndex string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	body := ApiUrl + "/recipient/unlock/opzones/cbond?" + "username=" + username + "&token=" + Token +
		"&projIndex=" + projIndex + "&seedpwd=" + seedpwd

	data, err := erpc.GetRequest(body)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func IncreaseTrustLimit(username string, seedpwd string, trust string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	body := ApiUrl + "/user/increasetrustlimit?" + "username=" + username + "&token=" + Token +
		"&seedpwd=" + seedpwd + "&trust=" + trust

	data, err := erpc.GetRequest(body)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func SendSharesEmail(username string, email1 string, email2 string, email3 string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	body := ApiUrl + "/user/sendrecovery?" + "username=" + username + "&token=" + Token +
		"&email1=" + email1 + "&email2=" + email2 + "&email3=" + email3

	data, err := erpc.GetRequest(body)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func SendNewSharesEmail(username string, seedpwd string, email1 string, email2 string, email3 string) (erpc.StatusResponse, error) {
	var x erpc.StatusResponse
	body := ApiUrl + "/user/newsecrets?" + "username=" + username + "&token=" + Token +
		"&seedpwd=" + seedpwd + "&email1=" + email1 + "&email2=" + email2 + "&email3=" + email3

	data, err := erpc.GetRequest(body)
	if err != nil {
		return x, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}
	return x, nil
}

func KillRpc(username string) {
	body := ApiUrl + "/admin/kill?" + "username=" + username + "&token=" + Token
	erpc.GetRequest(body)
	data, _ := erpc.GetRequest(ApiUrl + "/ping")
	var x erpc.StatusResponse
	// now data is in byte, we need the other structure now
	json.Unmarshal(data, &x)
	// the result would be the status of the platform
	if x.Code != 0 {
		ColorOutput("KILL COMMAND FAILED", RedColor)
	} else {
		ColorOutput("KILL COMMAND EXECUTED", GreenColor)
	}
}

func FreezeRpc(username string) {
	body := ApiUrl + "/admin/freeze?" + "username=" + username + "&token=" + Token
	erpc.GetRequest(body)
}

func GenKillCode(username string) (string, error) {
	body := ApiUrl + "/admin/gennuke?" + "username=" + username + "&token=" + Token
	data, err := erpc.GetRequest(body)
	if err != nil {
		return string(data), err
	}
	return string(data), nil
}
