package services

// This package contains the UI service implementation which produces the
// server-driven UI component lists and navigation. It was extracted from the
// api package to follow a services-oriented project structure.

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

// UIService defines the behaviour to produce UI components/data.
// Implementations can be swapped for tests or alternative data sources.
type UIService interface {
	Nav() []NavItem
	Home() []Component
	Products() []Component
	Profile() []Component
	ListDetail() []Component
	Detail(id string) ([]Component, bool)
	Forms() []Component
	Articles() []Component
	Combined() []Component
}

// DefaultUIService is the in-repo implementation of UIService.
type DefaultUIService struct{}

// NewDefaultUIService constructs the default UI service implementation.
func NewDefaultUIService() UIService { return &DefaultUIService{} }

// Service methods (migrated from the api package)
func (s *DefaultUIService) Nav() []NavItem {
	return []NavItem{
		{Label: "Home", Path: "/home"},
		{Label: "Products", Path: "/products"},
		{Label: "Profile", Path: "/profile"},
		{Label: "List & Detail", Path: "/list-detail"},
		{Label: "Forms", Path: "/forms"},
		{Label: "Articles", Path: "/articles"},
		{Label: "Combined", Path: "/combined"},
	}
}

func (s *DefaultUIService) Home() []Component {
	return []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "Welcome Back!", "level": 1}},
		{Type: "paragraph", Props: map[string]interface{}{"text": "Explore new features and products."}},
		{Type: "image", Props: map[string]interface{}{"src": "https://picsum.photos/600/200", "alt": "Banner"}},
		{Type: "button", Props: map[string]interface{}{"text": "View Products", "actionId": "navigate_products"}},
	}
}

func (s *DefaultUIService) Products() []Component {
	products := []map[string]interface{}{
		{"name": "Widget Pro", "price": "$99", "imageSrc": "https://picsum.photos/150/150?random=1"},
		{"name": "Gadget Plus", "price": "$149", "imageSrc": "https://picsum.photos/150/150?random=2"},
		{"name": "Thingamajig", "price": "$49", "imageSrc": "https://picsum.photos/150/150?random=3"},
	}

	productCards := make([]Component, 0, len(products))
	for _, p := range products {
		productCards = append(productCards, Component{Type: "product_card", Props: p})
	}

	return []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "Featured Products", "level": 2}},
		{Type: "product_list", Props: map[string]interface{}{"items": productCards}},
		{Type: "button", Props: map[string]interface{}{"text": "Go to Home", "actionId": "navigate_home"}},
	}
}

func (s *DefaultUIService) Profile() []Component {
	return []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "User Profile", "level": 3}},
		{Type: "stat_item", Props: map[string]interface{}{"label": "Email", "value": "user@example.com"}},
		{Type: "divider", Props: map[string]interface{}{}},
		{Type: "stat_item", Props: map[string]interface{}{"label": "Joined", "value": "Jan 2023"}},
		{Type: "button", Props: map[string]interface{}{"text": "Logout", "actionId": "user_logout"}},
	}
}

func (s *DefaultUIService) ListDetail() []Component {
	items := []map[string]interface{}{
		{"id": "1", "title": "Item 1", "description": "Description for item 1"},
		{"id": "2", "title": "Item 2", "description": "Description for item 2"},
		{"id": "3", "title": "Item 3", "description": "Description for item 3"},
	}

	listItems := make([]Component, 0, len(items))
	for _, item := range items {
		listItems = append(listItems, Component{Type: "list_item", Props: item})
	}

	return []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "List & Detail", "level": 2}},
		{Type: "list", Props: map[string]interface{}{"items": listItems}},
	}
}

func (s *DefaultUIService) Detail(id string) ([]Component, bool) {
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
		return comps, true
	}

	return []Component{{Type: "paragraph", Props: map[string]interface{}{"text": "Item not found"}}}, false
}

func (s *DefaultUIService) Forms() []Component {
	return []Component{
		{Type: "heading", Props: map[string]interface{}{"text": "Contact Form", "level": 2}},
		{Type: "form_input", Props: map[string]interface{}{"label": "Name", "name": "name", "type": "text"}},
		{Type: "form_input", Props: map[string]interface{}{"label": "Email", "name": "email", "type": "email"}},
		{Type: "form_input", Props: map[string]interface{}{"label": "Message", "name": "message", "type": "textarea"}},
		{Type: "form_button", Props: map[string]interface{}{"text": "Submit", "actionId": "form_submit"}},
	}
}

func (s *DefaultUIService) Articles() []Component {
	return []Component{
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
}

// Combined returns a composed page pulling pieces from multiple screens.
func (s *DefaultUIService) Combined() []Component {
	// Take featured products, a small profile block, a list, forms and some articles
	products := s.Products()
	profile := s.Profile()
	list := s.ListDetail()
	forms := s.Forms()
	articles := s.Articles()

	// Assemble: heading, products, divider, profile, divider, list, divider, forms, divider, articles
	comps := []Component{{Type: "heading", Props: map[string]interface{}{"text": "Combined Overview", "level": 1}}}
	comps = append(comps, products...)
	comps = append(comps, Component{Type: "divider", Props: map[string]interface{}{}})
	comps = append(comps, profile...)
	comps = append(comps, Component{Type: "divider", Props: map[string]interface{}{}})
	comps = append(comps, list...)
	comps = append(comps, Component{Type: "divider", Props: map[string]interface{}{}})
	comps = append(comps, forms...)
	comps = append(comps, Component{Type: "divider", Props: map[string]interface{}{}})
	comps = append(comps, articles...)

	return comps
}
