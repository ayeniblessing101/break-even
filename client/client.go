package main

import (
	"log"

	"github.com/ayeniblessing101/calculate-break-even/breakeven"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}

	defer conn.Close()

	c := breakeven.NewBreakEvenServiceClient(conn)

	response, err := c.CalculateBreakEven(context.Background(), &breakeven.Request{E: &breakeven.Estimation{DownPayment: 8000.0, MortgageInterestRate: 2.960, PropertyTax: 825, PropertyTransferTax: 825, Term: 48}, Rent: 399776.0, PriceOfPotentialHouse: 200000.00, LoanTerm: 30})

	if err != nil {
		log.Fatalf("An error occured: %s", err)
	}
	log.Printf("Response from: %s", response.Result)
}