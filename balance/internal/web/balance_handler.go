package web

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/marcohnp/fullcycle_eda/internal/usecase/get_balance"
	"net/http"
)

type WebBalanceHandler struct {
	GetBalanceUsecase get_balance.GetBalanceUsecase
}

func NewWebBalanceHandler(getBalanceUsecase get_balance.GetBalanceUsecase) *WebBalanceHandler {
	return &WebBalanceHandler{
		GetBalanceUsecase: getBalanceUsecase,
	}
}

func (h *WebBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	var input get_balance.GetBalanceInputDto
	accountId := chi.URLParam(r, "id")

	if accountId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("account id is required"))
		fmt.Println("account id is required")
		return
	}

	input.AccountId = accountId

	output, err := h.GetBalanceUsecase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}
}
