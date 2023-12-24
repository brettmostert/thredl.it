package ui

import "net/http"

// "github.com/google/uuid"
// "github.com/brettmostert/fnple/ledger/accounts"
// "github.com/brettmostert/fnple/ledger/transactions"
// "github.com/google/uuid"

func (s *ui) routes() {
	sr := s.router.PathPrefix("/ui").Subrouter()
	// subRoute =s.router.PathPrefix("/ui/")
	// s.router.HandleFunc("/accounts", s.handleAccount())
	// s.router.HandleFunc("/transactions", s.handleTransaction())
	// s.router.HandleFunc("/info", s.handleInfo())
	sr.HandleFunc("/info", s.handleInfo())
}

func (s *ui) handleInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ui\n"))
	}
}

// func (s *api) handleAccount() http.HandlerFunc {
// 	// thing := prepareThing()
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		acc := accounts.NewAccount{
// 			CoalationId: uuid.New().String(),
// 			Description: "123",
// 		}
// 		newAcc := accounts.CreateAccount(s.ctx, acc)
// 		b, _ := json.MarshalIndent(newAcc, "", "  ")
// 		print(string(b))
// 	}
// }

// func (s *api) handleTransaction() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var tx transactions.NewTransaction

// 		// breakout into utility func
// 		dec := json.NewDecoder(r.Body)
// 		dec.DisallowUnknownFields()
// 		dec.Decode(&tx)

// 		newTx := transactions.CreateTransaction(s.ctx, tx)
// 		_, _ = json.MarshalIndent(newTx, "", "  ")
// 	}
// }
