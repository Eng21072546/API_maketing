package controller

import (
	"fmt"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/logger"
	"github.com/Eng21072546/API_maketing/payload"
	"github.com/Eng21072546/API_maketing/useCase/interface"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"os"
)

type HttpProductHandler struct {
	productUseCase _interface.ProductUseCase
}

func NewHttpProductHandler(ProductUseCase _interface.ProductUseCase) *HttpProductHandler {
	return &HttpProductHandler{productUseCase: ProductUseCase}
}

func (h *HttpProductHandler) newProductHandlerLog() (*zap.Logger, func()) {

	logFactory := logger.NewLoggerFactory("./logger/logs")
	logger, loggerClose, err := logFactory.NewLogger()
	if err != nil {
		fmt.Printf("%v", err.Error())
		os.Exit(1)
	}
	//defer loggerClose()
	return logger, loggerClose

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
	log, logClose := h.newProductHandlerLog()
	defer logClose()

	idStr := c.Params("id")
	log.Info("GetProductById", zap.String("ProductId", idStr))
	var id int
	_, err := fmt.Sscan(idStr, &id) // Convert string ID to int
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid Request"})
	}

	product, err := h.productUseCase.GetProduct(c.Context(), id)
	if err != nil {
		log.Error("Product not found", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Product Not Found"})
	}
	return c.JSON(fiber.Map{"Product": product})
}
func (h *HttpProductHandler) CreateProduct(c *fiber.Ctx) error {
	log, logClose := h.newProductHandlerLog()
	defer logClose()

	var create payload.ProductCreate
	if err := c.BodyParser(&create); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Invalid Request"})
	}
	var product = &entity.Product{
		Name:  create.Name,
		Price: create.Price,
		Stock: create.Stock,
	}
	product, err := h.productUseCase.CreateProduct(c.Context(), product)
	if err != nil {
		log.Error("Product cannot createProduct", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Server Error"})
	}
	log.Info("Created product", zap.String("ProductId", string(product.ID)))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Product": product})
}

func (h *HttpProductHandler) UpdateProduct(c *fiber.Ctx) error {
	log, logClose := h.newProductHandlerLog()
	defer logClose()

	var update payload.ProductUpdate
	idStr := c.Params("id")

	var id int
	_, err := fmt.Sscan(idStr, &id) // Convert string ID to int
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid Request"})
	}
	if err := c.BodyParser(&update); err != nil {
		log.Error("Invalid req", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid Request"})
	}
	var productUpdate = &entity.ProductUpdate{ID: id, Name: update.Name, Price: update.Price, Stock: update.Stock}

	result, err := h.productUseCase.UpdateProduct(c.Context(), productUpdate)
	if err != nil {
		log.Error("cannot updateProduct", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	log.Info("Update Product UseCase", zap.String("ProductId", idStr))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Product": result})
}

func (h HttpProductHandler) DeleteProduct(c *fiber.Ctx) error {
	log, logClose := h.newProductHandlerLog()
	defer logClose()

	idStr := c.Params("id")
	var id int
	_, err := fmt.Sscan(idStr, &id) // Convert string ID to int
	if err != nil {
		log.Error("Invalid Request", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Invalid Request"})
	}
	result, err := h.productUseCase.DeleteProduct(c.Context(), id)
	if (result.DeletedCount == 0) || err != nil {
		log.Error("cannot delete product", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	log.Info("Delete Product UseCase", zap.String("ProductId", idStr))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result": result})
}
