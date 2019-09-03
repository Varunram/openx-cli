package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	// "strings"

	erpc "github.com/Varunram/essentials/rpc"
	scan "github.com/Varunram/essentials/scan"
	opensolar "github.com/YaleOpenLab/opensolar/core"
	wallet "github.com/YaleOpenLab/openx/chains/xlm/wallet"
	rpc "github.com/YaleOpenLab/openx/rpc"
)

func HttpsGetRequest(url string) ([]byte, error) {
	// Create a CA certificate pool and add cert.pem to it
	var dummy []byte
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return dummy, errors.Wrap(err, "did not create new GET request")
	}
	req.Header.Set("Origin", "localhost")
	res, err := client.Do(req)
	if err != nil {
		return dummy, errors.Wrap(err, "did not make request")
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// Login logs on to the platform
func Login(username string, pwhash string) (string, error) {
	var wString string

	// first get the accessToken with the help of a post req
	postData := url.Values{}
	postData.Set("username", username)
	postData.Set("pwhash", pwhash)

	data, err := erpc.PostForm(ApiUrl + "/token", postData)
	if err != nil {
		return wString, errors.Wrap(err, "failed to send a post request")
	}

	var t rpc.GenAccessTokenReturn
	err = json.Unmarshal(data, &t)
	if err != nil {
		log.Println(string(data))
		return wString, errors.Wrap(err, "could not unmarshal json")
	}

	Token = t.Token
	if len(Token) != 32 {
		return wString, errors.Wrap(err, "could not generate token")
	}
	log.Println("TOKEN=", Token)
	data, err = erpc.GetRequest(ApiUrl + "/user/validate?" + "username=" + username + "&token=" + Token)
	if err != nil {
		return wString, errors.Wrap(err, "validate request failed")
	}
	var x rpc.ValidateParams
	err = json.Unmarshal(data, &x)
	if err != nil {
		log.Println(string(data))
		return wString, errors.Wrap(err, "could not unmarshal json")
	}
	switch x.Role {
	case "Investor":
		wString = "Investor"
		data, err = erpc.GetRequest(ApiUrl + "/investor/validate?" + "username=" + username + "&pwhash=" + pwhash)
		if err != nil {
			return wString, errors.Wrap(err, "could not call ivnestor validate function")
		}
		var inv opensolar.Investor
		err = json.Unmarshal(data, &inv)
		if err != nil {
			return wString, errors.Wrap(err, "could not unmarshal json")
		}
		log.Println(string(data))
		LocalInvestor = inv
		ColorOutput("ENTER YOUR SEEDPWD: ", CyanColor)
		LocalSeedPwd, err = scan.ScanRawPassword()
		if err != nil {
			return wString, errors.Wrap(err, "could not scan raw password")
		}
		log.Println("LocalInvestor: ", LocalInvestor)
		log.Println("U: ", LocalInvestor.U)

		LocalSeed, err = wallet.DecryptSeed(LocalInvestor.U.StellarWallet.EncryptedSeed, LocalSeedPwd)
		if err != nil {
			return wString, errors.Wrap(err, "could not decrypt seed")
		}
	case "Recipient":
		wString = "Recipient"
		data, err = erpc.GetRequest(ApiUrl + "/recipient/validate?" + "username=" + username + "&pwhash=" + pwhash)
		if err != nil {
			return wString, errors.Wrap(err, "could not call recipient validate endpoint")
		}
		var recp opensolar.Recipient
		err = json.Unmarshal(data, &recp)
		if err != nil {
			return wString, errors.Wrap(err, "could not unmarshal json")
		}
		LocalRecipient = recp
		ColorOutput("ENTER YOUR SEEDPWD: ", CyanColor)
		LocalSeedPwd, err = scan.ScanRawPassword()
		if err != nil {
			return wString, errors.Wrap(err, "could not scan raw password")
		}
		LocalSeed, err = wallet.DecryptSeed(LocalRecipient.U.StellarWallet.EncryptedSeed, LocalSeedPwd)
		if err != nil {
			return wString, errors.Wrap(err, "could not decrypt seed")
		}
	case "Entity":
		log.Println("ENTITY?")
		data, err = erpc.GetRequest(ApiUrl + "/entity/validate?" + "username=" + username + "&pwhash=" + pwhash)
		if err != nil {
			return wString, errors.Wrap(err, "could not call entity validate endpoint")
		}
		var entity opensolar.Entity
		err = json.Unmarshal(data, &entity)
		if err != nil {
			return wString, errors.Wrap(err, "could not unmarshal json")
		}
		if entity.Contractor {
			LocalContractor = entity
			wString = "Contractor"
		} else if entity.Originator {
			LocalOriginator = entity
			wString = "Originator"
		} else {
			return wString, errors.New("Not a contractor")
		}
		ColorOutput("ENTER YOUR SEEDPWD: ", CyanColor)
		LocalSeedPwd, err = scan.ScanRawPassword()
		if err != nil {
			return wString, errors.Wrap(err, "could not scan raw password")
		}
		if entity.Contractor {
			LocalSeed, err = wallet.DecryptSeed(LocalContractor.U.StellarWallet.EncryptedSeed, LocalSeedPwd)
			if err != nil {
				return wString, errors.Wrap(err, "could not decrypt seed")
			}
		} else if entity.Originator {
			LocalSeed, err = wallet.DecryptSeed(LocalOriginator.U.StellarWallet.EncryptedSeed, LocalSeedPwd)
			if err != nil {
				return wString, errors.Wrap(err, "could not decrypt seed")
			}
		}
	}
	ColorOutput("AUTHENTICATED USER, YOUR ROLE IS: "+wString, GreenColor)
	return wString, nil
}
