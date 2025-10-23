package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Component represents a generic UI component
type Component struct {
	Type  string                 `json:"type"`
	Props map[string]interface{} `json:"props"`
}

// NavItem represents a navigation entry
type NavItem struct {
	Label string `json:"label"`
	Path  string `json:"path"`
}

// NavHandler returns the navigation structure
func NavHandler(c *gin.Context) {
	nav := []NavItem{
		{Label: "Home", Path: "/home"},
		{Label: "Products", Path: "/products"},
		{Label: "Profile", Path: "/profile"},
	}
	c.JSON(http.StatusOK, nav)
}

// HomeHandler returns the home screen components
func HomeHandler(c *gin.Context) {
	comps := []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "Welcome Back!2", "level": 1}},
		{Type: "paragraph", Props: map[string]interface{}{"text": "Explore new features and products."}},
		{Type: "image", Props: map[string]interface{}{"src": "https://picsum.photos/600/200", "alt": "Banner"}},
		{Type: "button", Props: map[string]interface{}{"text": "View Products", "actionId": "navigate_products"}},
	}
	c.JSON(http.StatusOK, comps)
}

// ProductsHandler returns a product list
func ProductsHandler(c *gin.Context) {
	products := []map[string]interface{}{
		{"name": "Widget Pro", "price": "$99", "imageSrc": "https://picsum.photos/150/150?random=1"},
		{"name": "Gadget Plus", "price": "$149", "imageSrc": "https://picsum.photos/150/150?random=2"},
		{"name": "Thingamajig", "price": "$49", "imageSrc": "https://picsum.photos/150/150?random=3"},
	}

	productCards := make([]Component, 0, len(products))
	for _, p := range products {
		productCards = append(productCards, Component{Type: "product_card", Props: p})
	}

	comps := []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "Featured Products", "level": 2}},
		{Type: "product_list", Props: map[string]interface{}{"items": productCards}},
		{Type: "button", Props: map[string]interface{}{"text": "Go to Home", "actionId": "navigate_home"}},
	}
	c.JSON(http.StatusOK, comps)
}

// ProfileHandler returns a profile screen
func ProfileHandler(c *gin.Context) {
	comps := []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "User Profile", "level": 3}},
		{Type: "stat_item", Props: map[string]interface{}{"label": "Email", "value": "user@example.com"}},
		{Type: "divider", Props: map[string]interface{}{}},
		{Type: "stat_item", Props: map[string]interface{}{"label": "Joined", "value": "Jan 2023"}},
		{Type: "button", Props: map[string]interface{}{"text": "Logout", "actionId": "user_logout"}},
	}
	c.JSON(http.StatusOK, comps)
}
