/*
 * Rosetta
 *
 * <h2>Backstory</h2> Writing reliable blockchain integrations is complicated and time-consuming. The process requires careful analysis of the unique aspects of each blockchain and extensive communication with its developers to understand the best strategies to deploy nodes, recognize deposits, broadcast transactions, etc. Even a minor misunderstanding can lead to downtime, or even worse, incorrect fund attribution. Not to mention, this integration must be continuously modified and tested each time a blockchain team releases new software.  Instead of spending time working on their blockchain, project developers spend countless hours answering similar support questions for each team integrating their blockchain. With their questions answered, each integrating team then writes similar code to interface with the blockchain instead of spending their engineering resources adding support for more blockchain projects or working on unique products and applications.  <h2>Standard for Blockchain Interaction</h2> Rosetta is a new project from Coinbase to standardize the process of deploying and interacting with blockchains. With an explicit specification to adhere to, all parties involved in blockchain development can spend less time figuring out how to integrate with each other and more time working on the novel advances that will push the blockchain ecosystem forward. In practice, this means that any blockchain project that implements the requirements outlined in this specification will enable exchanges, block explorers, and wallets to integrate with much less communication overhead and network-specific work.  <h5>© 2020 Coinbase</h5>
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

// A MempoolApiController binds http requests to an api service and writes the service results to the http response
type MempoolApiController struct {
	service MempoolApiServicer
}

// NewMempoolApiController creates a default api controller
func NewMempoolApiController(s MempoolApiServicer) Router {
	return &MempoolApiController{service: s}
}

// Routes returns all of the api route for the MempoolApiController
func (c *MempoolApiController) Routes() Routes {
	return Routes{
		{
			"Mempool",
			strings.ToUpper("Post"),
			"/mempool",
			c.Mempool,
		},
		{
			"MempoolTransaction",
			strings.ToUpper("Post"),
			"/mempool/transaction",
			c.MempoolTransaction,
		},
	}
}

// Mempool - Get All Mempool Transactions
func (c *MempoolApiController) Mempool(w http.ResponseWriter, r *http.Request) {
	mempoolRequest := &MempoolRequest{}
	if err := json.NewDecoder(r.Body).Decode(&mempoolRequest); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.Mempool(r.Context(), *mempoolRequest)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// MempoolTransaction - Get a Mempool Transaction
func (c *MempoolApiController) MempoolTransaction(w http.ResponseWriter, r *http.Request) {
	mempoolTransactionRequest := &MempoolTransactionRequest{}
	if err := json.NewDecoder(r.Body).Decode(&mempoolTransactionRequest); err != nil {
		w.WriteHeader(500)
		return
	}

	result, err := c.service.MempoolTransaction(r.Context(), *mempoolTransactionRequest)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}
