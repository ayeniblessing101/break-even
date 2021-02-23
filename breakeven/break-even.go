package breakeven

import (
	"fmt"
	"log"
	"math"

	"golang.org/x/net/context"
)

// Server struct that defines the server structure
type Server struct {
}

type estimation struct {
	downPayment float64
	mortgageInterestRate float64
	propertyTax float64
	propertyTransferTax float64
	term int
}

func (e estimation) calculateMortgageNeeded(priceOfPotentialHouse float64) float64 {
	return priceOfPotentialHouse - e.downPayment
}

func (e estimation) calculateMortgagePayableMonthly(loanTerm int64, priceOfPotentialHouse float64) float64 {
	principal := e.calculateMortgageNeeded(priceOfPotentialHouse)
	monthlyInterestRate := (e.mortgageInterestRate)/12
	n := float64(loanTerm * 12)

	monthlyMortgagePayment := (principal * monthlyInterestRate * math.Pow((1+(monthlyInterestRate)), n)) /
		math.Pow((1+(monthlyInterestRate)), n) - 1

		return monthlyMortgagePayment
}

func (e estimation) calculateTotalCostToBuyForTheTerm(loanTerm int64, priceOfPotentialHouse float64) float64 {
	monthlyMortgagePayment := e.calculateMortgagePayableMonthly(loanTerm, priceOfPotentialHouse)
	termInMonths := e.term * 12

	monthlyPayment := monthlyMortgagePayment + e.propertyTax + e.propertyTransferTax

	return (monthlyPayment * float64(termInMonths)) + e.downPayment
}	

func (e estimation) calculateTotalCostToRentForTheTerm(rent float64) float64 {
	termInMonths  := e.term * 12
	securityDeposit := rent

	return (rent * float64(termInMonths)) + securityDeposit
}

// CalculateBreakEven Method defined on the server struct
func (s *Server) CalculateBreakEven(ctx context.Context, request *Request) (*Response, error) {

	log.Printf("CalculateBreakEven called")


	e := estimation{
		downPayment: request.E.DownPayment,
		mortgageInterestRate: request.E.MortgageInterestRate,
		propertyTax: request.E.PropertyTax,
		propertyTransferTax: request.E.PropertyTransferTax,
		term: int(request.E.Term),
	}

	totalCostTotalCostToBuy := e.calculateTotalCostToBuyForTheTerm(request.LoanTerm, request.PriceOfPotentialHouse)
	totalCostTotalCostToRent := e.calculateTotalCostToRentForTheTerm(request.Rent)

	if totalCostTotalCostToBuy > totalCostTotalCostToRent {
		return &Response{Result: fmt.Sprintln("For the term you plan to live in this address, It is better for you to Rent than Buy")}, nil
	}
	return &Response{Result: fmt.Sprintln("For the term you plan to live in this address, It is better for you to Buy than Rent")}, nil
}

func (s *Server) mustEmbedUnimplementedBreakEvenServiceServer() {}