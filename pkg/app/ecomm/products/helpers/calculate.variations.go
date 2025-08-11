package helpers

import (
	"sort"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
)

func CalculateVariations(p *domain.Product) {
	variationMap := make(map[string]map[string]struct{})

	for _, item := range p.ProductItems {
		for _, opt := range item.Options {
			if _, exists := variationMap[opt.Name]; !exists {
				variationMap[opt.Name] = make(map[string]struct{})
			}
			variationMap[opt.Name][opt.VarOptName] = struct{}{}
		}
	}

	var variations []*domain.Variation
	for name, valuesSet := range variationMap {
		var values []string
		for val := range valuesSet {
			values = append(values, val)
		}
		sort.Strings(values) 
		variations = append(variations, &domain.Variation{
			Name:   name,
			Values: values,
		})
	}

	p.Variations = variations
}
