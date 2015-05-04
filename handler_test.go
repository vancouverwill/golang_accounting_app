package main

import (
	"github.com/vancouverwill/accountingApp/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	t.Log("TestIndex")
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	controllers.Index(w, req)

	if w.Code != 200 {
		t.Error("index() did not work as expected. the status was not 200")
	}

	t.Log("status:", w.Code, "body:", w.Body.String())
}

/**
*
* test simple 404 error page to make sure it returns the correct status if a user is lost
*
**/
func Test404Page(t *testing.T) {
	resp := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://localhost:8080/sdfsfd", nil)
	if err != nil {
		t.Fatal(err)
	}

	http.DefaultServeMux.ServeHTTP(resp, req)

	if resp.Code != 404 {
		t.Error("404 error is not being returned!")
	}
}

/**
*
* test the balance with a valid AccountAccountHolderOrCompany and relatedToId
*
**/
func TestBalancesIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/balances/?AccountAccountHolderOrCompany=Account&relatedToId=9", nil)
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	controllers.BalancesIndex(w, req)

	if w.Code != 200 && w.Code != 202 {
		t.Error("BalancesIndex() did not work as expected. the status was not ", http.StatusOK, ", it was ", w.Code)
	}

	t.Log("status:", w.Code, "body:", w.Body.String())
}

/**
*
* test non valid entry of "country" for AccountAccountHolderOrCompany paramater of GET
*
**/
func TestBalancesValidatesAccountAccountHolderOrCompany(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/balances/?AccountAccountHolderOrCompany=Country&relatedToId=234", nil)
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	controllers.BalancesIndex(w, req)

	if w.Code != 400 {
		t.Error("TransactionsIndex() did not work as expected. the status was not 400, it was ", w.Code)
	}
}

/**
*
* test GET all transactions without filter
*
**/
func TestTransactionsIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/transactions/", nil)
	if err != nil {
		t.Error("index() did not work as expected.")
		t.Log(err)
	}

	w := httptest.NewRecorder()
	controllers.TransactionsIndex(w, req)

	if w.Code != 200 && w.Code != 202 {
		t.Error("TransactionsIndex() did not work as expected. the status was not ", http.StatusOK, ", it was ", w.Code)
	}

	t.Log("status:", w.Code, "body:", w.Body.String())
}
