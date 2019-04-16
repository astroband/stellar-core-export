package es

import (
	"time"

	"github.com/stellar/go/xdr"
)

// BalanceSource represents type of balance record
type BalanceSource string

const (
	// BalanceSourceFee marks record came from tx_fee_meta
	BalanceSourceFee BalanceSource = "Fee"

	// BalanceSourceMeta marks record came from tx_meta
	BalanceSourceMeta BalanceSource = "Meta"
)

// Balance represents balance log entry
type Balance struct {
	AccountID string        `json:"account_id"`
	Balance   int           `json:"balance"`
	CreatedAt time.Time     `json:"created_at"`
	Source    BalanceSource `json:"source"`
	Asset     Asset         `json:"asset"`
}

// NewBalanceFromAccountEntry creates Balance from AccountEntry
func NewBalanceFromAccountEntry(a xdr.AccountEntry) *Balance {
	return &Balance{
		AccountID: a.AccountId.Address(),
		Balance:   int(a.Balance),
		Source:    BalanceSourceMeta,
		CreatedAt: time.Now(),
		Asset:     *NewNativeAsset(),
	}
}

// NewBalanceFromTrustLineEntry creates Balance from TrustLineEntry
func NewBalanceFromTrustLineEntry(t xdr.TrustLineEntry) *Balance {
	return &Balance{
		AccountID: t.AccountId.Address(),
		Balance:   int(t.Balance),
		Source:    BalanceSourceMeta,
		CreatedAt: time.Now(),
		Asset:     *NewAsset(&t.Asset),
	}
}

// ExtractBalances returns balances extracted from metas
func ExtractBalances(c []xdr.LedgerEntryChange) []*Balance {
	var prev = make(map[string]int)
	var balances []*Balance

	for _, change := range c {
		switch t := change.Type; t {
		case xdr.LedgerEntryChangeTypeLedgerEntryState:
			state := change.MustState().Data

			switch x := state.Type; x {
			case xdr.LedgerEntryTypeAccount:
				account := state.MustAccount()
				prev[account.AccountId.Address()] = int(account.Balance)
			case xdr.LedgerEntryTypeTrustline:
				line := state.MustTrustLine()
				prev[line.AccountId.Address()] = int(line.Balance)
			}

		case xdr.LedgerEntryChangeTypeLedgerEntryCreated:
			created := change.MustCreated().Data

			switch x := created.Type; x {
			case xdr.LedgerEntryTypeAccount:
				balances = append(balances, NewBalanceFromAccountEntry(created.MustAccount()))
			case xdr.LedgerEntryTypeTrustline:
				balances = append(balances, NewBalanceFromTrustLineEntry(created.MustTrustLine()))
			}

			// case xdr.LedgerEntryChangeTypeLedgerEntryUpdated:
			// 	updated := change.MustUpdated().Data

			// 	switch x := updated.Type; x {
			// 	case xdr.LedgerEntryTypeAccount:

			// 		//updated.MusAccountId.Address()
			// 		//balances = append(balances, NewBalanceFromAccountEntry(created.MustAccount()))
			// 	case xdr.LedgerEntryTypeTrustline:
			// 		//balances = append(balances, NewBalanceFromTrustLineEntry(created.MustTrustLine()))
			// 	}

			// 	//case xdr.LedgerEntryChangeTypeLedgerEntryRemoved:

		}
	}

	return balances
}

// DocID balance es document id
func (b *Balance) DocID() *string {
	return nil
}

// IndexName balances index name
func (b *Balance) IndexName() string {
	return balanceIndexName
}
