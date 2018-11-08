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
	"strings"
	"sync"

	"github.com/hexoul/go-coinsuper/types"
	"github.com/hexoul/go-coinsuper/util"
)

// Interface for APIs
type Interface interface {
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
		} else if arg[0] == "-accesskey" {
			accessKey = arg[1]
		} else if arg[0] == "-secret" {
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
	params := "accesskey=" + s.accessKey + "&secretkey=" + s.secretKey
	hasher := md5.New()
	hasher.Write([]byte(params))
	sign := hex.EncodeToString(hasher.Sum(nil))

	common := &types.RequestCommon{
		AccessKey: s.accessKey,
		Sign:      sign,
	}
	return &types.Request{
		Common: common,
		Data:   &types.Options{},
	}
}

func (s *Client) getResponse(url string, req *types.Request) (*types.Response, []byte, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("reqBody %s\n", string(reqBody))
	httpReq, err := http.NewRequest("POST", url, bytes.NewBufferString(string(reqBody)))
	if err != nil {
		return nil, nil, err
	}
	body, err := util.DoReq(httpReq)
	if err != nil {
		return nil, nil, err
	}
	fmt.Print(string(body))
	return nil, body, nil
}

// UserAssetInfo obtains your own personal asset information
func (s *Client) UserAssetInfo(options *types.Options) (*types.Response, error) {
	url := fmt.Sprintf("%s/asset/userAssetInfo", baseURL)

	_, _, err := s.getResponse(url, s.parseOptions(options))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
