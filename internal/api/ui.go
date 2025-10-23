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
		{Label: "List & Detail", Path: "/list-detail"},
		{Label: "Forms", Path: "/forms"},
		{Label: "Articles", Path: "/articles"},
	}
	c.JSON(http.StatusOK, nav)
}

// HomeHandler returns the home screen components
func HomeHandler(c *gin.Context) {
	comps := []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "Welcome Back!", "level": 1}},
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

// ListDetailHandler returns a list of items for list-and-detail screen
func ListDetailHandler(c *gin.Context) {
	items := []map[string]interface{}{
		{"id": "1", "title": "Item 1", "description": "Description for item 1"},
		{"id": "2", "title": "Item 2", "description": "Description for item 2"},
		{"id": "3", "title": "Item 3", "description": "Description for item 3"},
	}

	listItems := make([]Component, 0, len(items))
	for _, item := range items {
		listItems = append(listItems, Component{Type: "list_item", Props: item})
	}

	comps := []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "List & Detail", "level": 2}},
		{Type: "list", Props: map[string]interface{}{"items": listItems}},
	}
	c.JSON(http.StatusOK, comps)
}

// DetailHandler returns the detail for a specific item
func DetailHandler(c *gin.Context) {
	id := c.Param("id")
	// Mock data
	details := map[string]map[string]interface{}{
		"1": {"title": "Item 1", "fullDescription": "Full description for item 1. This is more detailed.", "image": "https://picsum.photos/300/200?random=1"},
		"2": {"title": "Item 2", "fullDescription": "Full description for item 2. This is more detailed.", "image": "https://picsum.photos/300/200?random=2"},
		"3": {"title": "Item 3", "fullDescription": "Full description for item 3. This is more detailed.", "image": "https://picsum.photos/300/200?random=3"},
	}

	if detail, exists := details[id]; exists {
		comps := []Component{
			{Type: "heading", Props: map[string]interface{}{"text": detail["title"].(string), "level": 2}},
			{Type: "image", Props: map[string]interface{}{"src": detail["image"].(string), "alt": detail["title"].(string)}},
			{Type: "paragraph", Props: map[string]interface{}{"text": detail["fullDescription"].(string)}},
			{Type: "button", Props: map[string]interface{}{"text": "Back to List", "actionId": "navigate_list-detail"}},
		}
		c.JSON(http.StatusOK, comps)
	} else {
		c.JSON(http.StatusNotFound, []Component{{Type: "paragraph", Props: map[string]interface{}{"text": "Item not found"}}})
	}
}

// FormsHandler returns a form screen
func FormsHandler(c *gin.Context) {
	comps := []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "Contact Form", "level": 2}},
		{Type: "form_input", Props: map[string]interface{}{"label": "Name", "name": "name", "type": "text"}},
		{Type: "form_input", Props: map[string]interface{}{"label": "Email", "name": "email", "type": "email"}},
		{Type: "form_input", Props: map[string]interface{}{"label": "Message", "name": "message", "type": "textarea"}},
		{Type: "form_button", Props: map[string]interface{}{"text": "Submit", "actionId": "form_submit"}},
	}
	c.JSON(http.StatusOK, comps)
}

// ArticlesHandler returns an articles screen
func ArticlesHandler(c *gin.Context) {
	comps := []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "Latest Articles", "level": 2}},
		{Type: "article", Props: map[string]interface{}{
			"title":   "Article 1",
			"content": "This is the content of article 1. It can be long and detailed.",
			"author":  "Author 1",
			"date":    "2023-10-01",
		}},
		{Type: "article", Props: map[string]interface{}{
			"title":   "Article 2",
			"content": "This is the content of article 2. More text here.",
			"author":  "Author 2",
			"date":    "2023-10-02",
		}},
	}
	c.JSON(http.StatusOK, comps)
}
