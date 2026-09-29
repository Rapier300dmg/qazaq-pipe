package models

import (
	"context"
	"pipe-company/backend/internal/database"
	"time"
)

type Lead struct {
	ID             int64     `json:"id"`
	ClientName     string    `json:"client_name"`
	CompanyName    string    `json:"company_name"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	DiameterMM     int       `json:"diameter_mm"`
	SDR            float64   `json:"sdr"`
	QuantityMeters int       `json:"quantity_meters"`
	Comment        string    `json:"comment"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

func CreateLead(lead *Lead) error {
	query := `
		INSERT INTO leads (client_name, company_name, phone, email, diameter_mm, sdr, quantity_meters, comment)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, status, created_at`

	return database.DB.QueryRow(
		context.Background(),
		query,
		lead.ClientName, lead.CompanyName, lead.Phone, lead.Email,
		lead.DiameterMM, lead.SDR, lead.QuantityMeters, lead.Comment,
	).Scan(&lead.ID, &lead.Status, &lead.CreatedAt)
}