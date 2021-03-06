// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"strings"

	moov "github.com/moov-io/go-client/client"

	"github.com/antihax/optional"
)

func createAccount(ctx context.Context, api *moov.APIClient, u *user, name, number string) (*moov.Account, error) {
	req := moov.CreateAccount{
		CustomerID: u.ID,
		Name:       name,
		Number:     number,
		Type:       "Savings",
		Balance:    1000 * 100, // $1,000
	}
	opts := &moov.CreateAccountOpts{}
	account, resp, err := api.AccountsApi.CreateAccount(ctx, u.ID, req, opts)
	if resp != nil {
		resp.Body.Close()
		if err := checkCORSHeaders(resp); err != nil {
			return nil, fmt.Errorf("create account: %v", err)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("problem creating account %q: %v", name, err)
	}
	return &account, nil
}

var (
	defaultRoutingNumber = "121042882"
)

func createMicroDepositAccount(ctx context.Context, api *moov.APIClient, u *user) (*moov.Account, error) {
	// The hardcoded values here need to match paygate's expectations for the micro-deposit origination account
	//
	// TODO(adam): We need to do something about micro-deposit buildup that slowly depletes this account balance.
	// It causes problems eventually as the balance runs too low for any further micro-deposits (apitest runs).
	opts := &moov.SearchAccountsOpts{
		Number:        optional.NewString("123"),
		RoutingNumber: optional.NewString(defaultRoutingNumber),
		Type_:         optional.NewString("Savings"),
	}
	accounts, resp, err := api.AccountsApi.SearchAccounts(ctx, u.ID, opts)
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
		if err := checkCORSHeaders(resp); err != nil {
			return nil, fmt.Errorf("create micro-deposit: %v", err)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("micro-deposits: SearchAccounts: userID=%s: %v", u.ID, err)
	}
	if len(accounts) > 0 {
		return &accounts[0], nil
	}
	return createAccount(ctx, api, u, "micro-deposit origination", "123")
}

// Verify accountID and Transaction exist of a given amount (used to double check transfers).
func checkTransactions(ctx context.Context, api *moov.APIClient, accountID string, u *user, amount string) error {
	opts := &moov.GetAccountTransactionsOpts{
		Limit: optional.NewFloat32(25),
	}
	transactions, resp, err := api.AccountsApi.GetAccountTransactions(ctx, accountID, u.ID, opts)
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
		if err := checkCORSHeaders(resp); err != nil {
			return fmt.Errorf("check transactions: %v", err)
		}
	}
	if err != nil {
		return fmt.Errorf("accounts: GetAccountTransactions: %v", err)
	}
	for i := range transactions {
		for j := range transactions[i].Lines {
			// match transaction against posted ones on the account
			line := transactions[i].Lines[j]
			if v := fmt.Sprintf("USD %.2f", float32(line.Amount)/100.0); line.AccountID == accountID && v == amount {
				return nil // Matched Transaction
			}
		}
	}
	return fmt.Errorf("accounts: unable to find %q transaction for account=%s", amount, accountID)
}

func getMicroDepositsTransactions(ctx context.Context, api *moov.APIClient, accountID string, u *user) ([]*moov.Transaction, error) {
	opts := &moov.GetAccountTransactionsOpts{
		Limit: optional.NewFloat32(25),
	}
	transactions, resp, err := api.AccountsApi.GetAccountTransactions(ctx, accountID, u.ID, opts)
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
		if err := checkCORSHeaders(resp); err != nil {
			return nil, fmt.Errorf("get micro-deposit transactions: %v", err)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("accounts: getMicroDeposits: %v", err)
	}
	var txs []*moov.Transaction
	for i := range transactions {
		if len(transactions[i].Lines) != 2 {
			continue
		}
		if *flagDebug {
			out := ""
			for j := range transactions[i].Lines {
				line := transactions[i].Lines[j]
				out += fmt.Sprintf("\n  accountID=%s purpose=%s amount=%v", line.AccountID, line.Purpose, line.Amount)
			}
		}
		for j := range transactions[i].Lines {
			line := transactions[i].Lines[j]
			if line.AccountID == accountID && strings.EqualFold(line.Purpose, "achcredit") && line.Amount < 100 {
				txs = append(txs, &transactions[i])
				break
			}
		}
	}
	if len(txs) == 0 {
		return nil, fmt.Errorf("unable to find micro-deposit transaction (found %d)", len(transactions))
	}
	return txs, nil
}
