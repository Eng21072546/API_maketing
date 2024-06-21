package controller

import (
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/payload"
	"github.com/Eng21072546/API_maketing/useCase"
	"github.com/gofiber/fiber/v2"
)

type HttpProductHandler struct {
	productUseCase useCase.ProductUseCase
}

func NewHttpProductHandler(ProductUseCase useCase.ProductUseCase) *HttpProductHandler {
	return &HttpProductHandler{productUseCase: ProductUseCase}
}

func (h *HttpProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.productUseCase.GetAllProduct()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Server Error"})
	}
	return c.JSON(fiber.Map{"Products": products})
}
func (h *HttpProductHandler) GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.productUseCase.GetProduct(id)
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
	product, err := h.productUseCase.CreateProduct(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Product": product})
}

func (h *HttpProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var update payload.ProductUpdate
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid Request"})
	}
	var product = &entity.ProductUpdate{Name: update.Name, Price: update.Price, Stock: update.Stock}
	id := c.Params("id")
	result, err := h.productUseCase.UpdateProduct(id, product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Product": result})
}

func (h HttpProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.productUseCase.DeleteProduct(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result": result})
}
