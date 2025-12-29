package handlers

import (
	"net/http"
	"spec-builder/backend/internal/db"
	"spec-builder/backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm/clause"
)

func GetCabinets(c *gin.Context) {
	var cabinets []models.Cabinet
	db.DB.Distinct("department, feature, number, name, id").Find(&cabinets)
	c.JSON(200, cabinets)
}

func GenerateSpec(c *gin.Context) {
	var req struct {
		CabinetIDs []uint `json:"cabinet_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var equipment []models.Equipment
	db.DB.Where("cabinet_id IN ?", req.CabinetIDs).Find(&equipment)

	specMap := make(map[string]*models.SpecItem)
	for _, eq := range equipment {
		key := eq.ItemName + "|" + eq.Manufacturer + "|" + eq.Model
		if item, exists := specMap[key]; exists {
			item.TotalQty += eq.QtyToOrder
		} else {
			specMap[key] = &models.SpecItem{
				ItemName:     eq.ItemName,
				StandardName: eq.StandardName,
				Manufacturer: eq.Manufacturer,
				Model:        eq.Model,
				Article:      eq.Article,
				TotalQty:     eq.QtyToOrder,
				Specialist:   eq.Specialist,
				TMCType:      eq.TMCType,
			}
		}
	}

	spec := make([]*models.SpecItem, 0, len(specMap))
	for _, item := range specMap {
		spec = append(spec, item)
	}

	c.JSON(200, spec)
}

func ExportSpecExcel(c *gin.Context) {
	var req struct {
		CabinetIDs []uint `json:"cabinet_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var equipment []models.Equipment
	db.DB.Where("cabinet_id IN ?", req.CabinetIDs).Find(&equipment)

	specMap := make(map[string]*models.SpecItem)
	for _, eq := range equipment {
		key := eq.ItemName + "|" + eq.Manufacturer + "|" + eq.Model
		if item, exists := specMap[key]; exists {
			item.TotalQty += eq.QtyToOrder
		} else {
			specMap[key] = &models.SpecItem{
				ItemName:     eq.ItemName,
				StandardName: eq.StandardName,
				Manufacturer: eq.Manufacturer,
				Model:        eq.Model,
				Article:      eq.Article,
				TotalQty:     eq.QtyToOrder,
				Specialist:   eq.Specialist,
				TMCType:      eq.TMCType,
			}
		}
	}

	f := excelize.NewFile()
	sheet := "Спецификация"
	f.SetCellValue(sheet, "A1", "Наименование по Медси")
	f.SetCellValue(sheet, "B1", "Производитель")
	f.SetCellValue(sheet, "C1", "Модель")
	f.SetCellValue(sheet, "D1", "Артикул")
	f.SetCellValue(sheet, "E1", "Кол-во")
	f.SetCellValue(sheet, "F1", "Специалист")
	f.SetCellValue(sheet, "G1", "Вид ТМЦ")

	row := 2
	for _, item := range specMap {
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), item.ItemName)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), item.Manufacturer)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), item.Model)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), item.Article)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), item.TotalQty)
		f.SetCellValue(sheet, "F"+strconv.Itoa(row), item.Specialist)
		f.SetCellValue(sheet, "G"+strconv.Itoa(row), item.TMCType)
		row++
	}

	f.SetColWidth(sheet, "A", "G", 25)

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=specification.xlsx")
	f.Write(c.Writer)
}
