package mapper

import (
	"my-wallet-ntier-mongo/model"
	"my-wallet-ntier-mongo/response"
)

func TransactionTypeModelToResponse(transactionType model.TransactionType) response.TransactionTypeResponse {
	subCategories := []response.SubCategoryResponse{}
	for _, subCat := range transactionType.SubCategories {
		subCategories = append(subCategories, response.SubCategoryResponse{
			ID:   subCat.ID,
			Name: subCat.Name,
		})
	}

	return response.TransactionTypeResponse{
		ID:            transactionType.ID.Hex(),
		Name:          transactionType.Name,
		SubCategories: subCategories,
		Type:          transactionType.Type,
	}
}
