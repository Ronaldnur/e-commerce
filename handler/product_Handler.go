package handler

import (
	"encoding/json"
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	ProductService service.ProductService
}

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{
		ProductService: productService,
	}
}

func (ph *productHandler) MakeProduct(ctx *gin.Context) {

	var newProduct dto.NewProductRequest

	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	user := ctx.MustGet("userData").(entity.User)

	result, err := ph.ProductService.ProductCreate(&newProduct, user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (ph *productHandler) GetAllData(ctx *gin.Context) {
	// Panggil metode untuk mengambil semua data produk
	products, err := ph.ProductService.GetProductData()
	if err != nil {
		// Kembalikan kesalahan jika terjadi
		ctx.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	// Kirimkan data produk sebagai JSON
	ctx.JSON(http.StatusOK, products)
}

func (ph *productHandler) GetOneProduct(ctx *gin.Context) {
	productId := ctx.Param("productId")

	result, err := ph.ProductService.GetOneProductData(productId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (ph *productHandler) UpdateProductById(ctx *gin.Context) {
	productId := ctx.Param("productId")

	var newUpdateProduct dto.NewProductRequest

	if err := ctx.ShouldBindJSON(&newUpdateProduct); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := ph.ProductService.UpdateProductById(productId, &newUpdateProduct)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)

}

func (ph *productHandler) DeleteProduct(ctx *gin.Context) {

	productId := ctx.Param("productId")

	result, err := ph.ProductService.DeleteProduct(productId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)

}

func (ph *productHandler) UpdateStock(ctx *gin.Context) {
	productId := ctx.Param("productId")

	var updateStock dto.UpdateStockRequest

	if err := ctx.ShouldBindJSON(&updateStock); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := ph.ProductService.UpdateProductStock(productId, updateStock)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)

}

// func (ph *productHandler) HttpMakeProduct(w http.ResponseWriter, r *http.Request) {
// 	// Set content type for response
// 	w.Header().Set("Content-Type", "application/json")

// 	// Decode the JSON request body directly into the DTO
// 	var newProduct dto.NewProductRequest
// 	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
// 		// Handle error for invalid request body
// 		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
// 		http.Error(w, errBindJson.Message(), errBindJson.Status())
// 		return
// 	}

// 	// Call the service to create the product
// 	result, err := ph.ProductService.ProductCreate(&newProduct)
// 	if err != nil {
// 		// Handle error dari service
// 		w.WriteHeader(err.Status()) // Status code dari error
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"statusCode": err.Status(),
// 			"message":    err.Message(),
// 			"error":      err.Error(),
// 		})
// 		return
// 	}
// 	// Send successful response as JSON
// 	w.WriteHeader(result.StatusCode)
// 	if err := json.NewEncoder(w).Encode(result); err != nil {
// 		// Handle error if response encoding fails
// 		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
// 		return
// 	}

// }

func (ph *productHandler) HttpGetOneProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("productId")

	result, err := ph.ProductService.GetOneProductData(productId)

	if err != nil {
		// Handle error dari service
		w.WriteHeader(err.Status()) // Status code dari error
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": err.Status(),
			"message":    err.Message(),
			"error":      err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)

}

func (ph *productHandler) HttpGetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := ph.ProductService.GetProductData()

	if err != nil {
		// Handle error dari service
		w.WriteHeader(err.Status()) // Status code dari error
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": err.Status(),
			"message":    err.Message(),
			"error":      err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Failed to encode products", http.StatusInternalServerError)
	}
}

func (ph *productHandler) HttpDeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	productId := r.URL.Query().Get("productId")

	result, err := ph.ProductService.DeleteProduct(productId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)

}

func (ph *productHandler) HttpUpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	productId := r.URL.Query().Get("productId")

	// Decode the JSON request body directly into the DTO
	var newProduct dto.NewProductRequest
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		// Handle error for invalid request body
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		http.Error(w, errBindJson.Message(), errBindJson.Status())
		return
	}

	result, err := ph.ProductService.UpdateProductById(productId, &newProduct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)

	if err != nil {
		// Handle any errors from the service
		http.Error(w, err.Message(), err.Status())
		return
	}

	// Send successful response as JSON
	w.WriteHeader(result.StatusCode)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		// Handle error if response encoding fails
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}
