/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"context"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// AccountApiService is a service that implents the logic for the AccountApiServicer
// This service should implement the business logic for every endpoint for the AccountApi API.
// Include any external packages or services that will be required by this service.
type AccountApiService struct {
	celoClient *client.CeloClient
}

// NewAccountApiService creates a default api service
func NewAccountApiService(celoClient *client.CeloClient) AccountApiServicer {
	return &AccountApiService{
		celoClient: celoClient,
	}
}

// AccountBalance - Get an Account Balance
func (s *AccountApiService) AccountBalance(ctx context.Context, accountBalanceRequest AccountBalanceRequest) (interface{}, error) {

	err := ValidateNetworkId(&accountBalanceRequest.NetworkIdentifier, s.celoClient.Net, ctx)
	if err != nil {
		return BuildErrorResponse(1, err), nil
	}

	address := common.HexToAddress(accountBalanceRequest.AccountIdentifier.Address)

	latestHeader, err := s.celoClient.Eth.HeaderByNumber(ctx, nil) // nil == latest
	if err != nil {
		return BuildErrorResponse(2, err), nil
	}

	goldBalance, err := s.celoClient.Eth.BalanceAt(ctx, address, latestHeader.Number)
	if err != nil {
		return BuildErrorResponse(3, err), nil
	}

	registry, err := contract.NewRegistry(RegistrySmartContractAddress, s.celoClient.Eth)
	if err != nil {
		return BuildErrorResponse(4, err), nil
	}

	lockedGoldAddr, err := registry.GetAddressForString(&bind.CallOpts{
		BlockNumber: latestHeader.Number,
		Context:     ctx,
	}, LockedGoldRegistryId)
	if err != nil {
		return BuildErrorResponse(5, err), nil
	}

	lockedGold, err := contract.NewLockedGold(lockedGoldAddr, s.celoClient.Eth)
	if err != nil {
		return BuildErrorResponse(6, err), nil
	}

	lockedGoldBalance, err := lockedGold.GetAccountTotalLockedGold(&bind.CallOpts{
		BlockNumber: latestHeader.Number,
		Context:     ctx,
	}, address)
	if err != nil {
		return BuildErrorResponse(7, err), nil
	}

	// STABLETOKEN DISABLED FOR v1.0.0
	// stableTokenAddr, err := registry.GetAddressForString(&bind.CallOpts{
	// 	BlockNumber: latestHeader.Number,
	// 	Context:     ctx,
	// }, StableTokenRegistryId)
	// if err != nil {
	// 	return BuildErrorResponse(6, err), nil
	// }

	// stableToken, err := contract.NewStableToken(stableTokenAddr, s.celoClient.Eth)
	// if err != nil {
	// 	return BuildErrorResponse(7, err), nil
	// }

	// stableTokenBalance, err := stableToken.BalanceOf(&bind.CallOpts{
	// 	BlockNumber: latestHeader.Number,
	// 	Context:     ctx,
	// }, address)
	// if err != nil {
	// 	return BuildErrorResponse(8, err), nil
	// }

	response := AccountBalanceResponse{
		BlockIdentifier: *HeaderToBlockIdentifier(latestHeader),
		Balances: []Balance{
			Balance{
				AccountIdentifier: accountBalanceRequest.AccountIdentifier,
				Amounts: []Amount{
					Amount{
						Value:    goldBalance.String(),
						Currency: CeloGold,
					},
					// Amount{
					// 	Value:    stableTokenBalance.String(),
					// 	Currency: CeloDollar,
					// },
				},
			},
			Balance{
				AccountIdentifier: AccountIdentifier{
					Address: accountBalanceRequest.AccountIdentifier.Address,
					SubAccount: SubAccountIdentifier{
						SubAccount: LockedGoldRegistryId,
					},
				},
				Amounts: []Amount{
					Amount{
						Value:    lockedGoldBalance.String(),
						Currency: CeloGold,
					},
				},
			},
		},
	}
	return response, nil
}
