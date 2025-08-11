package logger

import (
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
)

func PrintProducts(products []*domain.Product) {

	for i, product := range products {
		fmt.Printf("---------- Product %d ------------------\n", i+1)
		fmt.Printf("Product name %v\n", product.Name)
		fmt.Printf("Product desc %v\n", product.Description)
		fmt.Printf("Product category name %v\n", product.CategoryName)
		fmt.Printf("Product filepath %v\n", product.FilePath)
		fmt.Printf("Product file length %v\n", len(product.File))
		fmt.Printf("Product content-type %v\n", product.ContentType)

		for j, pItem := range product.ProductItems {
			fmt.Printf("---------- Product Item %d ------------------\n", j+1)
			fmt.Printf("Status %v\n", pItem.Status)
			fmt.Printf("Qty in stock %v\n", pItem.QtyInStock)
			fmt.Printf("Discount %v\n", pItem.Discount)
			fmt.Printf("Price %v\n", pItem.Price)
			fmt.Printf("FilePath %v\n", pItem.FilePath)
			fmt.Printf("ProductItem file length %v\n", len(pItem.File))

			fmt.Printf("Content-Type %v\n", pItem.ContentType)
			for k, option := range pItem.Options {
				fmt.Printf("---------- Option %d ------------------\n", k+1)
				fmt.Printf("Name %v\n", option.Name)
				fmt.Printf("varOptionName %v\n", option.VarOptName)
				fmt.Printf("varOptionValue %v\n", option.VarOptValue)
			}
		}
	}
}
