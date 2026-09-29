package models

import (
	"context"
	"fmt"
	"pipe-company/backend/internal/database"
	"time"
)

type Product struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	PEGrade         string    `json:"pe_grade"`
	OuterDiameterMM int       `json:"outer_diameter_mm"`
	SDR             float64   `json:"sdr"`
	WallThicknessMM float64   `json:"wall_thickness_mm"`
	PressurePN      float64   `json:"pressure_pn"`
	PricePerMeter   float64   `json:"price_per_meter"`
	CreatedAt       time.Time `json:"created_at"`
}

type ProductFilter struct {
	PEGrade  string
	Diameter int
	SDR      float64
}

func GetFilteredProducts(filter ProductFilter) ([]Product, error) {
	query := "SELECT id, name, pe_grade, outer_diameter_mm, sdr, wall_thickness_mm, pressure_pn, price_per_meter, created_at FROM products WHERE 1=1"
	args := []interface{}{}
	argId := 1

	if filter.PEGrade != "" {
		query += fmt.Sprintf(" AND pe_grade = $%d", argId)
		args = append(args, filter.PEGrade)
		argId++
	}
	if filter.Diameter > 0 {
		query += fmt.Sprintf(" AND outer_diameter_mm = $%d", argId)
		args = append(args, filter.Diameter)
		argId++
	}
	if filter.SDR > 0 {
		query += fmt.Sprintf(" AND sdr = $%d", argId)
		args = append(args, filter.SDR)
		argId++
	}

	query += " ORDER BY outer_diameter_mm ASC"

	rows, err := database.DB.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]Product, 0)
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.PEGrade, &p.OuterDiameterMM, &p.SDR, &p.WallThicknessMM, &p.PressurePN, &p.PricePerMeter, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}