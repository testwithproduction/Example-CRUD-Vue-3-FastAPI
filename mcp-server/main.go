package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"mcp-server/database"
	"mcp-server/models"
)

func main() {
	database.InitDB()

	s := server.NewMCPServer("Product MCP Server", "1.0.0")

	s.AddTool(mcp.NewTool("list_products",
		mcp.WithDescription("List all products"),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var products []models.Product
		result := database.DB.Find(&products)
		if result.Error != nil {
			return mcp.NewToolResultError(fmt.Sprintf("DB error: %v", result.Error)), nil
		}
		b, _ := json.MarshalIndent(products, "", "  ")
		return mcp.NewToolResultText(string(b)), nil
	})

	s.AddTool(mcp.NewTool("get_product",
		mcp.WithDescription("Get a product by ID"),
		mcp.WithString("id", mcp.Required(), mcp.Description("Product ID")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("id")
		if err != nil {
			return mcp.NewToolResultError("Missing or invalid 'id' argument"), nil
		}
		var product models.Product
		result := database.DB.First(&product, id)
		if result.Error != nil {
			return mcp.NewToolResultError("Product not found"), nil
		}
		b, _ := json.MarshalIndent(product, "", "  ")
		return mcp.NewToolResultText(string(b)), nil
	})

	s.AddTool(mcp.NewTool("create_product",
		mcp.WithDescription("Create a new product"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Product name")),
		mcp.WithNumber("price", mcp.Required(), mcp.Description("Product price")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := req.RequireString("name")
		if err != nil {
			return mcp.NewToolResultError("Missing or invalid 'name' argument"), nil
		}
		price, err := req.RequireFloat("price")
		if err != nil {
			return mcp.NewToolResultError("Missing or invalid 'price' argument"), nil
		}
		product := models.Product{Name: name, Price: price}
		result := database.DB.Create(&product)
		if result.Error != nil {
			return mcp.NewToolResultError(fmt.Sprintf("DB error: %v", result.Error)), nil
		}
		b, _ := json.MarshalIndent(product, "", "  ")
		return mcp.NewToolResultText(string(b)), nil
	})

	s.AddTool(mcp.NewTool("update_product",
		mcp.WithDescription("Update an existing product"),
		mcp.WithString("id", mcp.Required(), mcp.Description("Product ID")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Product name")),
		mcp.WithNumber("price", mcp.Required(), mcp.Description("Product price")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("id")
		if err != nil {
			return mcp.NewToolResultError("Missing or invalid 'id' argument"), nil
		}
		name, err := req.RequireString("name")
		if err != nil {
			return mcp.NewToolResultError("Missing or invalid 'name' argument"), nil
		}
		price, err := req.RequireFloat("price")
		if err != nil {
			return mcp.NewToolResultError("Missing or invalid 'price' argument"), nil
		}
		var product models.Product
		result := database.DB.First(&product, id)
		if result.Error != nil {
			return mcp.NewToolResultError("Product not found"), nil
		}
		product.Name = name
		product.Price = price
		result = database.DB.Save(&product)
		if result.Error != nil {
			return mcp.NewToolResultError(fmt.Sprintf("DB error: %v", result.Error)), nil
		}
		b, _ := json.MarshalIndent(product, "", "  ")
		return mcp.NewToolResultText(string(b)), nil
	})

	s.AddTool(mcp.NewTool("delete_product",
		mcp.WithDescription("Delete a product by ID"),
		mcp.WithString("id", mcp.Required(), mcp.Description("Product ID")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("id")
		if err != nil {
			return mcp.NewToolResultError("Missing or invalid 'id' argument"), nil
		}
		result := database.DB.Delete(&models.Product{}, id)
		if result.Error != nil {
			return mcp.NewToolResultError(fmt.Sprintf("DB error: %v", result.Error)), nil
		}
		if result.RowsAffected == 0 {
			return mcp.NewToolResultError("Product not found"), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("Product %s deleted", id)), nil
	})

	port := os.Getenv("MCP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting MCP-Go server on :%s", port)

	sseServer := server.NewSSEServer(s)
	if err := sseServer.Start(":" + port); err != nil {
		log.Fatal(err)
	}
} 