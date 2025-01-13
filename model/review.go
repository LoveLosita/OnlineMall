package model

type AddReview struct {
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	ParentID  *int   `json:"parent_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}

type ShowReview struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	Replies   []ShowReview
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ReplyToReview struct {
	ReplyToID int    `json:"review_id"`
	Reply     string `json:"reply"`
}
