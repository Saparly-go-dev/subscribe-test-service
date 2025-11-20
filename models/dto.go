package models

type SubscriptionCreateRequest struct {
	ServiceName string `json:"service_name" example:"Yandex Plus" binding:"required"`
	Price       int    `json:"price" example:"400" binding:"required,gte=0"`
	UserID      string `json:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba" binding:"required,uuid4"`
	StartDate   string `json:"start_date" example:"07-2025" binding:"required"` // MM-YYYY
	EndDate     string `json:"end_date" example:"09-2025"`                      // optional
}

type SubscriptionUpdateRequest struct {
	ServiceName *string `json:"service_name,omitempty"`
	Price       *int    `json:"price,omitempty"`
	StartDate   *string `json:"start_date,omitempty"`
	EndDate     *string `json:"end_date,omitempty"`
}

type SubscriptionResponse struct {
	ID          string `json:"id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"` // MM-YYYY
	EndDate     string `json:"end_date,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type SummaryResponse struct {
	TotalPrice int `json:"total_price" example:"1600"`
}
