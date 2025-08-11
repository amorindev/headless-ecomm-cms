package helpers

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
)

func GenerateItemSKU(p *domain.Product, item *domain.ProductItem) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Prefijo categoría
	categoryPart := "CAT"
	if p.CategoryName != "" {
		categoryPart = strings.ToUpper(p.CategoryName[:min(3, len(p.CategoryName))])
	}

	// Prefijo nombre producto
	namePart := strings.ToUpper(strings.ReplaceAll(p.Name, " ", ""))[:min(3, len(p.Name))]

	// Variaciones (ej. RED, L)
	var optionPart string
	if len(item.Options) > 0 {
		for _, opt := range item.Options {
			optionPart += strings.ToUpper(opt.VarOptName[:min(2, len(opt.VarOptName))])
		}
	}

	// Número aleatorio
	randomPart := fmt.Sprintf("%04d", r.Intn(10000))

	return fmt.Sprintf("%s-%s-%s-%s", categoryPart, namePart, optionPart, randomPart)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

