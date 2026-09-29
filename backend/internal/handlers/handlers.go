package handlers

import (
	"encoding/json"
	"net/http"
	"pipe-company/backend/internal/models"
	"strconv"
)

// GET /api/v1/products?pe_grade=PE-100&diameter=110
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	diameter, _ := strconv.Atoi(query.Get("diameter"))
	sdr, _ := strconv.ParseFloat(query.Get("sdr"), 64)

	filter := models.ProductFilter{
		PEGrade:  query.Get("pe_grade"),
		Diameter: diameter,
		SDR:      sdr,
	}

	products, err := models.GetFilteredProducts(filter)
	if err != nil {
		http.Error(w, "Ошибка получения данных из БД", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// GET /api/v1/contacts
func GetContactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contacts := map[string]string{
		"company":    "ТОО QAZAQ PIPE",
		"phone":      "+7 (701) 000-00-00",
		"whatsapp":   "https://wa.me/77010000000",
		"email":      "sales@qazaqpipe.kz",
		"address":    "Республика Казахстан, г. Актау, Промышленная зона",
		"work_hours": "Пн - Пт: с 09:00 до 18:00",
	}

	json.NewEncoder(w).Encode(contacts)
}