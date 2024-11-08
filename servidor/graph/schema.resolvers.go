package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kissmarkrivas/sumago/graph/generated"
	"github.com/kissmarkrivas/sumago/graph/model"
)

// CreateSuma is the resolver for the createSuma field.
func (r *mutationResolver) CreateSuma(ctx context.Context, input model.SumInput) (*model.Suma, error) {
	sumatorias := model.Suma{
		Nombre: input.Nombre,
		Sum1:   input.Sum1,
		Sum2:   input.Sum2,
		Result: input.Result,
	}
	err := r.DB.Create(&sumatorias).Error
	if err != nil {
		return nil, err
	}
	return &sumatorias, nil
}

// UpdateSuma is the resolver for the updateSuma field.
func (r *mutationResolver) UpdateSuma(ctx context.Context, sumaID int, input model.SumInput) (*model.Suma, error) {
	updateSuma := model.Suma{
		ID:     sumaID,
		Nombre: input.Nombre,
		Sum1:   input.Sum1,
		Sum2:   input.Sum2,
		Result: input.Result,
	}
	err := r.DB.Save(&updateSuma).Error
	if err != nil {
		return nil, err
	}
	return &updateSuma, nil
}

// DeleteSuma is the resolver for the deleteSuma field.
func (r *mutationResolver) DeleteSuma(ctx context.Context, sumaID int) (bool, error) {
	r.DB.Where("suma_id = ?", sumaID).Delete(&model.Suma{})
	r.DB.Where("id = ?", sumaID).Delete(&model.Suma{})
	return true, nil
}

// Sumas is the resolver for the sumas field.
func (r *queryResolver) Sumas(ctx context.Context) ([]*model.Suma, error) {
	var sumas []*model.Suma
	// err := r.DB.Preload("Items").Find(&sumas).Error
	err := r.DB.Set("gorm:auto_preload", true).Find(&sumas).Error
	if err != nil {
		return nil, err
	}
	return sumas, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
