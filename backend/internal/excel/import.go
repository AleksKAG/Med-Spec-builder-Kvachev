package excel

import (
	"strconv"
	"strings"
	"spec-builder/backend/internal/db"
	"spec-builder/backend/internal/models"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func ImportFromXLSX(path string) error {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, sheet := range f.GetSheetList() {
		rows, err := f.GetRows(sheet)
		if err != nil || len(rows) < 2 {
			continue
		}

		for _, row := range rows[1:] {
			if len(row) < 20 {
				continue
			}

			// Пропускаем строки с пустым "Наименование кабинета"
			if strings.TrimSpace(row[5]) == "" {
				continue
			}

			cab := models.Cabinet{
				Department: cleanStr(row[2]),
				Feature:    cleanStr(row[3]),
				Number:     cleanStr(row[4]),
				Name:       cleanStr(row[5]),
			}

			result := db.DB.Where(cab).FirstOrCreate(&cab)
			if result.Error != nil {
				continue
			}

			qtyStr := cleanStr(row[9])
			qty := 0
			if qtyStr != "" && qtyStr != "расчетное кол-о" && qtyStr != "по требованию" {
				if q, err := strconv.Atoi(qtyStr); err == nil {
					qty = q
				}
			}

			eq := models.Equipment{
				CabinetID:         cab.ID,
				ItemName:          cleanStr(row[7]),
				StandardName:      cleanStr(row[8]),
				QtyToOrder:        qty,
				Manufacturer:      cleanStr(row[10]),
				Model:             cleanStr(row[11]),
				Article:           cleanStr(row[12]),
				Specialist:        cleanStr(row[30]),
				TMCType:           cleanStr(row[31]),
			}

			if eq.ItemName != "" {
				db.DB.Create(&eq)
			}
		}
	}
	return nil
}

func cleanStr(s string) string {
	return strings.TrimSpace(s)
}
