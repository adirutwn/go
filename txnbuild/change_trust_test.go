package txnbuild

import (
	"github.com/stellar/go/network"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeTrust_BuildXDR(t *testing.T) {
	t.Run("change trust source account is defined", func(t *testing.T) {
		kp0 := newKeypair0()
		txSourceAccount := horizon.Account{
			HistoryAccount: horizon.HistoryAccount{
				AccountID: kp0.Address(),
			},
			Sequence: "1",
		}

		kp1 := newKeypair1()

		changeTrustOp := ChangeTrust{
			Line: CreditAsset{
				Code:   "ABC",
				Issuer: kp0.Address(),
			},
			Limit: "100",
			SourceAccount: &horizon.Account{
				HistoryAccount: horizon.HistoryAccount{
					AccountID: kp1.Address(),
				},
			},
		}

		tx := Transaction{
			SourceAccount: &txSourceAccount,
			Operations:    []Operation{&changeTrustOp},
			Timebounds:    NewInfiniteTimeout(),
			Network:       network.TestNetworkPassphrase,
		}
		received := buildSignEncode(t, tx, kp0, kp1)

		// ref: https://www.stellar.org/laboratory/#xdr-viewer?input=AAAAAODcbeFyXKxmUWK1L6znNbKKIkPkHRJNbLktcKPqLnLFAAAAZAAAAAAAAAACAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAEAAAAAJcrx2g%2FHbs%2FohF5CVFG7B5JJSJR%2BOqDKzDGK7dKHZH4AAAAGAAAAAUFCQwAAAAAA4Nxt4XJcrGZRYrUvrOc1sooiQ%2BQdEk1suS1wo%2BoucsUAAAAAO5rKAAAAAAAAAAAC6i5yxQAAAEDXbTvlrszCXSWWcAqTRguDu85r0DQTdb3v9G5WyJScjmQk0CGUGTpBV%2BcDntgXIL24MIHa%2FOUCmVkpLOdTHGgO0odkfgAAAEDCof5Mzix5bBMArg%2BiZ5fpUohvUvXrBW%2BYH74p7qOFeMlf9SElUKY86ZsL6iwD7wZ9pIwECcxq6YSKxeI6ckUG&type=TransactionEnvelope&network=test
		expectedXdr := "AAAAAODcbeFyXKxmUWK1L6znNbKKIkPkHRJNbLktcKPqLnLFAAAAZAAAAAAAAAACAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAEAAAAAJcrx2g/Hbs/ohF5CVFG7B5JJSJR+OqDKzDGK7dKHZH4AAAAGAAAAAUFCQwAAAAAA4Nxt4XJcrGZRYrUvrOc1sooiQ+QdEk1suS1wo+oucsUAAAAAO5rKAAAAAAAAAAAC6i5yxQAAAEDXbTvlrszCXSWWcAqTRguDu85r0DQTdb3v9G5WyJScjmQk0CGUGTpBV+cDntgXIL24MIHa/OUCmVkpLOdTHGgO0odkfgAAAEDCof5Mzix5bBMArg+iZ5fpUohvUvXrBW+YH74p7qOFeMlf9SElUKY86ZsL6iwD7wZ9pIwECcxq6YSKxeI6ckUG"
		assert.Equal(t, expectedXdr, received)
	})

	t.Run("change trust source account is empty", func(t *testing.T) {
		kp0 := newKeypair0()
		txSourceAccount := horizon.Account{
			HistoryAccount: horizon.HistoryAccount{
				AccountID: kp0.Address(),
			},
			Sequence: "1",
		}

		changeTrustOp := ChangeTrust{
			Line: CreditAsset{
				Code:   "ABC",
				Issuer: kp0.Address(),
			},
			Limit: "100",
		}

		tx := Transaction{
			SourceAccount: &txSourceAccount,
			Operations:    []Operation{&changeTrustOp},
			Timebounds:    NewInfiniteTimeout(),
			Network:       network.TestNetworkPassphrase,
		}
		received := buildSignEncode(t, tx, kp0)

		// ref: AAAAAODcbeFyXKxmUWK1L6znNbKKIkPkHRJNbLktcKPqLnLFAAAAZAAAAAAAAAACAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAGAAAAAUFCQwAAAAAA4Nxt4XJcrGZRYrUvrOc1sooiQ
		expectedXdr := "AAAAAODcbeFyXKxmUWK1L6znNbKKIkPkHRJNbLktcKPqLnLFAAAAZAAAAAAAAAACAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAGAAAAAUFCQwAAAAAA4Nxt4XJcrGZRYrUvrOc1sooiQ+QdEk1suS1wo+oucsUAAAAAO5rKAAAAAAAAAAAB6i5yxQAAAEC4Ukr2r6E1MdZoAqR8MA/5mSCbUWlsaaUOY6QZ7UB6EA3TsesoTGz2g20tepWo09knomr+EJFH+sgeFLrVE9MJ"
		assert.Equal(t, expectedXdr, received)
	})
}
