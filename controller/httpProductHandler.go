package controller

import (
	"fmt"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/payload"
	"github.com/Eng21072546/API_maketing/useCase"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type HttpProductHandler struct {
	productUseCase useCase.ProductUseCase
}

func NewHttpProductHandler(ProductUseCase useCase.ProductUseCase) *HttpProductHandler {
	return &HttpProductHandler{productUseCase: ProductUseCase}
}

func (h *HttpProductHandler) GetAllProducts(c *fiber.Ctx) error {
	fmt.Println("productUseCase")
	products, err := h.productUseCase.GetAllProduct(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Server Error"})
	}
	return c.JSON(fiber.Map{"Products": products})
}
func (h *HttpProductHandler) GetProductById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id int
	_, err := fmt.Sscan(idStr, &id) // Convert string ID to int
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid Request"})
	}

	product, err := h.productUseCase.GetProduct(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Product Not Found"})
	}
	return c.JSON(fiber.Map{"Product": product})
}
func (h *HttpProductHandler) CreateProduct(c *fiber.Ctx) error {
	var create payload.ProductCreate
	if err := c.BodyParser(&create); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid Request"})
	}
	var product = &entity.Product{
		Name:  create.Name,
		Price: create.Price,
		Stock: create.Stock,
	}
	product, err := h.productUseCase.CreateProduct(c.Context(), product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Product": product})
}

func (h *HttpProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var update payload.ProductUpdate
	updateQuery := make(bson.M)
	idStr := c.Params("id")
	var id int
	_, err := fmt.Sscan(idStr, &id) // Convert string ID to int
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid Request"})
	}
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid Request"})
	}
	var productUpdate = &entity.ProductUpdate{Name: update.Name, Price: update.Price, Stock: update.Stock}
	if productUpdate.Name != nil {
		updateQuery["name"] = update.Name
	}
	if productUpdate.Price != nil {
		updateQuery["price"] = update.Price
	}
	if productUpdate.Stock != nil {
		updateQuery["stock"] = update.Stock
	}

	result, err := h.productUseCase.UpdateProduct(c.Context(), id, updateQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Product": result})
}

func (h HttpProductHandler) DeleteProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	var id int
	_, err := fmt.Sscan(idStr, &id) // Convert string ID to int
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid Request"})
	}
	result, err := h.productUseCase.DeleteProduct(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result": result})
}
