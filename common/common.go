package common

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	etherCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type HttpRequestError struct {
	StatusCode int
	Body       []byte
}

func (a HttpRequestError) Error() string {
	return fmt.Sprintf("status: %d, body: %s", a.StatusCode, string(a.Body))
}

func MakeGetRequest(url string, headers map[string]string, query map[string]string, timeout time.Duration, response interface{}) error {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body := res.Body
	defer body.Close()

	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return HttpRequestError{
			StatusCode: res.StatusCode,
			Body:       bytes,
		}
	}
	return json.Unmarshal(bytes, response)
}

func GetAccountFromEnv() (etherCommon.Address, *ecdsa.PrivateKey, error) {
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return etherCommon.Address{}, nil, fmt.Errorf("failed to parse private key %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return etherCommon.Address{}, nil, errors.New("failed to get public Key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return fromAddress, privateKey, nil
}

func WaitForTransactionReceipt(ctx context.Context, ethClient *ethclient.Client, txHash etherCommon.Hash, timeCoolDownWaitForTx, timeout time.Duration) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		receipt, err := ethClient.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, fmt.Errorf("error fetching receipt: %v", err)
		}

		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("transaction not mined within %v", timeout)
		case <-ticker.C:
			time.Sleep(timeCoolDownWaitForTx)
			continue
		}
	}
}
