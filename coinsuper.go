// Package coinsuper is an API Client for Coinsuper
package coinsuper

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hexoul/go-coinsuper/types"
	"github.com/hexoul/go-coinsuper/util"
)

// Interface for APIs
type Interface interface {
	UserAssetInfo(options *types.Options) (*types.UserInfo, error)
	TradeHistory(options *types.Options) (*types.TransactionInfoList, error)
}

// Client for Coinsuper API
type Client struct {
	accessKey string
	secretKey string
}

var (
	instance  *Client
	once      sync.Once
	accessKey string
	secretKey string
)

const (
	baseURL = "https://api.coinsuper.com/api/v1"
)

func init() {
	for _, val := range os.Args {
		arg := strings.Split(val, "=")
		if len(arg) < 2 {
			continue
		} else if arg[0] == "-coinsuper:accesskey" {
			accessKey = arg[1]
		} else if arg[0] == "-coinsuper:secretkey" {
			secretKey = arg[1]
		}
	}
}

// GetInstance returns singleton
func GetInstance() *Client {
	once.Do(func() {
		if accessKey == "" || secretKey == "" {
			panic("KEYS FOR BOTH ACCESS AND SECRET REQUIRED")
		}
		instance = &Client{
			accessKey: accessKey,
			secretKey: secretKey,
		}
	})
	return instance
}

// GetInstanceWithKey returns singleton
func GetInstanceWithKey(accessKey, secretKey string) *Client {
	once.Do(func() {
		if accessKey == "" || secretKey == "" {
			panic("KEYS FOR BOTH ACCESS AND SECRET REQUIRED")
		}
		instance = &Client{
			accessKey: accessKey,
			secretKey: secretKey,
		}
	})
	return instance
}

func (s *Client) parseOptions(options *types.Options) *types.Request {
	// Make params
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	params := []string{}
	if options == nil {
		options = &types.Options{}
	}
	options.AccessKey = s.accessKey
	options.SecretKey = s.secretKey
	options.Timestamp = timestamp
	if bOption, err := json.Marshal(options); err == nil {
		mOption := new(map[string]interface{})
		if err := json.Unmarshal(bOption, &mOption); err == nil {
			for _, k := range types.OptionOrder {
				if (*mOption)[k] != nil {
					params = append(params, fmt.Sprintf("%s=%v", k, (*mOption)[k]))
				}
			}
		}
	}
	options.AccessKey = ""
	options.SecretKey = ""
	options.Timestamp = ""

	// Sign
	hasher := md5.New()
	hasher.Write([]byte(strings.Join(params, "&")))
	sign := hex.EncodeToString(hasher.Sum(nil))

	// Make request
	common := &types.RequestCommon{
		AccessKey: s.accessKey,
		Sign:      sign,
		Timestamp: timestamp,
	}
	return &types.Request{
		Common: common,
		Data:   options,
	}
}

func (s *Client) getResponse(url string, req *types.Request) (*types.Response, []byte, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBufferString(string(reqBody)))
	if err != nil {
		return nil, nil, err
	}
	body, err := util.DoReq(httpReq)
	if err != nil {
		return nil, nil, err
	}
	resp := new(types.Response)
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, err
	}
	if resp.Code != "1000" {
		return nil, nil, fmt.Errorf("[%s] %s", resp.Code, resp.Msg)
	}
	return resp, body, nil
}

// UserAssetInfo obtains your own personal asset information
//   arg: -
//   src: https://api.coinsuper.com/api/v1/asset/userAssetInfo
//   doc: https://github.com/coinsuperapi/API_docs_en/wiki#11-personal-asset-information
func (s *Client) UserAssetInfo(options *types.Options) (*types.UserInfo, error) {
	url := fmt.Sprintf("%s/asset/userAssetInfo", baseURL)

	resp, _, err := s.getResponse(url, s.parseOptions(options))
	if err != nil {
		return nil, err
	}

	var result = new(types.UserInfo)
	b, err := json.Marshal(resp.Data.Result)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, result); err != nil {
		return nil, err
	}
	return result, nil
}

// TradeHistory returns personal trade history
//   arg: utcStart(mandatory), symbol, utcEnd, startDealNo, size
//   src: https://api.coinsuper.com/api/v1/order/tradeHistory
//   doc: https://github.com/coinsuperapi/API_docs_en/wiki#29-query-personal-trade-history
func (s *Client) TradeHistory(options *types.Options) (*types.TransactionInfoList, error) {
	url := fmt.Sprintf("%s/order/tradeHistory", baseURL)

	resp, _, err := s.getResponse(url, s.parseOptions(options))
	if err != nil {
		return nil, err
	}

	var result = new(types.TransactionInfoList)
	b, err := json.Marshal(resp.Data.Result)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, result); err != nil {
		return nil, err
	}
	return result, nil
}
