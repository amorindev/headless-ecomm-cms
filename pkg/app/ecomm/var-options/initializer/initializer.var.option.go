package initializer

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/domain"
	varOptP "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/port"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/port"
)

type VariationOptionItz struct {
	VariationOptionRepo varOptP.VariationOptionRepo
	VariationRepo       port.VariationRepo
}

func NewVariationOptionItz(variationOptionRepo varOptP.VariationOptionRepo, variationRepo port.VariationRepo) *VariationOptionItz {
	return &VariationOptionItz{
		VariationOptionRepo: variationOptionRepo,
		VariationRepo:       variationRepo,
	}
}

func (voi *VariationOptionItz) SeedEssential(ctx context.Context) error {
	vOptionsColor := []*domain.VariationOption{
		domain.NewVarOpt("Black", "#000000"),
		domain.NewVarOpt("White", "#FFFFFF"),
		domain.NewVarOpt("Blue", "#0B19D9"),
		domain.NewVarOpt("Grey", "#666666"),
		domain.NewVarOpt("Brown", "#9C5300"),
		domain.NewVarOpt("Pink", "#9C5300"),
	}

	vOptionsSize := []*domain.VariationOption{
		domain.NewVarOpt("6", ""),
		domain.NewVarOpt("7", ""),
		domain.NewVarOpt("36", ""),
		domain.NewVarOpt("37", ""),
		domain.NewVarOpt("38", ""),
		domain.NewVarOpt("39", ""),
		domain.NewVarOpt("40", ""),
		domain.NewVarOpt("41", ""),
		domain.NewVarOpt("42", ""),
		domain.NewVarOpt("X", ""),
		domain.NewVarOpt("XS", ""),
		domain.NewVarOpt("S", ""),
		domain.NewVarOpt("M", ""),
		domain.NewVarOpt("L", ""),
		domain.NewVarOpt("XL", ""),
	}

	// * get the reference
	colorVariation, err := voi.VariationRepo.FindByName(context.Background(), "Color")
	if err != nil {
		return err
	}

	sizeVariation, err := voi.VariationRepo.FindByName(context.Background(), "Size")
	if err != nil {
		return err
	}

	for _, vOptionColor := range vOptionsColor {
		vOptionColor.VariationID = colorVariation.ID
		exists, err := voi.VariationOptionRepo.Exists(ctx, vOptionColor.Name)
		if err != nil {
			return err
		}
		if !exists {
			if err := voi.VariationOptionRepo.Insert(ctx, vOptionColor); err != nil {
				return err
			}
		}
	}

	for _, vOptionSize := range vOptionsSize {
		vOptionSize.VariationID = sizeVariation.ID
		exists, err := voi.VariationOptionRepo.Exists(ctx, vOptionSize.Name)
		if err != nil {
			return err
		}
		if !exists {
			if err := voi.VariationOptionRepo.Insert(ctx, vOptionSize); err != nil {
				return err
			}
		}
	}
	return nil
}
