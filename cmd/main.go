package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/ibldzn/fincloud-base/internal/fincloud"
	"github.com/ibldzn/fincloud-base/internal/fincloudapi"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file: %v\n", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	fincloudApi, err := fincloudapi.NewClient(
		fincloudapi.WithBaseURL(os.Getenv("FINCLOUD_API_BASE_URL")),
		fincloudapi.WithSecretKey(os.Getenv("FINCLOUD_API_SECRET_KEY")),
	)
	if err != nil {
		panic(err)
	}
	_ = fincloudApi

	fincloudClient, err := fincloud.NewClient(
		fincloud.Credentials{
			Username:   os.Getenv("FINCLOUD_CLIENT_USERNAME"),
			Password:   os.Getenv("FINCLOUD_CLIENT_PASSWORD"),
			LocationID: os.Getenv("FINCLOUD_CLIENT_LOCATION_ID"),
			RoleID:     os.Getenv("FINCLOUD_CLIENT_ROLE_ID"),
		},
		fincloud.WithBaseURL(os.Getenv("FINCLOUD_CLIENT_BASE_URL")),
	)
	if err != nil {
		panic(err)
	}

	if err := fincloudClient.Login(ctx); err != nil {
		panic(err)
	}

	data, err := ReadCSV("data_utomo.csv")
	if err != nil {
		panic(err)
	}

	result, err := os.Create("result.csv")
	if err != nil {
		panic(err)
	}
	defer result.Close()

	writer := csv.NewWriter(result)
	defer writer.Flush()

	writer.Comma = ';'

	writer.Write([]string{
		"journalId",
		"transactionId",
		"trxReference",
		"accountNumber",
		"alternateNumber",
		"principalPaid",
		"interestPaid",
		"penaltyPaid",
		"interestWaive",
		"penaltyWaive",
		"description",
	})

	for _, row := range data {
		nominal, err := strconv.Atoi(row["nominal"])
		if err != nil {
			fmt.Printf("error parsing nominal for account %s: %v\n", row["rekening"], err)
			continue
		}

		payload := fincloudapi.LoanTerminationRequest{
			TrxReference:   "abcd",
			AccountNumber:  row["rekening"],
			AltNumber:      "",
			PrincipalPaid:  int64(nominal),
			InterestPaid:   0,
			PenaltyPaid:    0,
			PrincipalWaive: 0,
			InterestWaive:  0,
			Description:    "Pengalihan Hak Tagih (Cessie) Kredit Pensiunan",
			BranchCode:     row["cabang"],
		}

		resp, err := fincloudApi.TerminateLoan(ctx, payload)
		if err != nil {
			fmt.Printf("error processing loan termination for account %s: %v\n", row["rekening"], err)
			writer.Write([]string{
				"", "", "", row["rekening"], "", "", "", "", "", "", fmt.Sprintf("error: %v", err),
			})
			continue
		}

		writer.Write([]string{
			resp.JournalID,
			resp.TransactionID,
			resp.TrxReference,
			resp.AccountNumber,
			resp.AlternateNumber,
			resp.PrincipalPaid,
			resp.InterestPaid,
			resp.PenaltyPaid,
			resp.InterestWaive,
			resp.PenaltyWaive,
			resp.Description,
		})
	}

	fmt.Println("Processing completed. Results written to result.csv")
}

func ReadCSV(path string) ([]map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.TrimLeadingSpace = true

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("csv kosong")
	}

	headers := rows[0]
	// remove BOM character from the first header if present
	if len(headers) > 0 {
		headers[0] = strings.TrimPrefix(headers[0], "\uFEFF")
	}
	var result []map[string]string

	for _, row := range rows[1:] {
		data := make(map[string]string)

		for i, header := range headers {
			header = strings.TrimSpace(header)

			if i < len(row) {
				data[header] = strings.TrimSpace(row[i])
			} else {
				data[header] = ""
			}
		}

		result = append(result, data)
	}

	return result, nil
}
