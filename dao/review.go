package dao

import (
	"OnlineMall/model"
	"OnlineMall/respond"
)

func AddReviewOrReply(review model.AddReview) error {
	query := "INSERT INTO reviews(user_id,product_id,parent_id,rating,comment,is_anonymous) VALUES(?,?,?,?,?,?)"
	_, err := Db.Exec(query, review.UserID, review.ProductID, review.ParentID, review.Rating, review.Comment, review.ISAnonymous)
	if err != nil {
		return err
	}
	return nil
}

func IfUserHasReviewedThisProduct(userID int, productID int) (bool, error) {
	query := "SELECT id FROM reviews WHERE user_id=? AND product_id=?"
	rows, err := Db.Query(query, userID, productID)
	if err != nil {
		return false, err
	}
	if rows.Next() { //如果有这个评论
		return true, nil
	} else { //如果没有这个评论
		return false, nil
	}
}

func GetProductRatings(productID int) ([]int, error) {
	query := "SELECT rating FROM reviews WHERE product_id=?"
	rows, err := Db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	var ratings []int
	for rows.Next() { //遍历所有获取的评分
		var rating int
		err = rows.Scan(&rating)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}
	return ratings, nil
}

func IfReviewExists(parentID int) (bool, error) {
	query := "SELECT id FROM reviews WHERE id=?"
	rows, err := Db.Query(query, parentID)
	if err != nil {
		return false, err
	}
	if rows.Next() { //如果有这个评论
		return true, nil
	} else { //如果没有这个评论
		return false, nil
	}
}

func GetAProductReviews(productID int) ([]model.ShowReview, error) {
	//1.获取直接属于该商品的评论
	query := "SELECT * FROM reviews WHERE product_id=?"
	rows, err := Db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	//2.获取回复上面评论的评论
	query = "SELECT * FROM reviews WHERE parent_id IN (SELECT id FROM reviews WHERE product_id=?)"
	rows2, err := Db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	//3.将两组评论放入切片中
	var reviews []model.ShowReview
	for rows.Next() { //遍历所有获取的评论
		var review model.ShowReview
		err = rows.Scan(&review.ID, &review.UserID, &review.ProductID, &review.ParentID, &review.Rating,
			&review.Comment, &review.ISAnonymous, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	for rows2.Next() { //遍历所有获取的评论
		var review model.ShowReview
		err = rows2.Scan(&review.ID, &review.UserID, &review.ProductID, &review.ParentID, &review.Rating,
			&review.Comment, &review.ISAnonymous, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func SearchForProductReviews(productID int, keyword string) ([]model.ShowReview, error) {
	//1.获取直接属于该商品的评论
	query := "SELECT * FROM reviews WHERE product_id=? AND comment LIKE ?"
	rows, err := Db.Query(query, productID, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	//2.获取回复上面评论的评论或者有相同关键词的回复
	query = "SELECT * FROM reviews WHERE parent_id IN (SELECT id FROM reviews WHERE product_id=? AND comment LIKE ?)"
	rows2, err := Db.Query(query, productID, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	//3.获取同样含有关键词的回复
	query = "SELECT * FROM reviews WHERE product_id IS NOT NULL AND comment LIKE ?"
	rows3, err := Db.Query(query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	//4.将3组评论放入切片中
	var reviews []model.ShowReview
	for rows.Next() { //遍历所有获取的评论
		var review model.ShowReview
		err = rows.Scan(&review.ID, &review.UserID, &review.ProductID, &review.ParentID, &review.Rating,
			&review.Comment, &review.ISAnonymous, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	for rows2.Next() { //遍历所有获取的评论
		var review model.ShowReview
		err = rows2.Scan(&review.ID, &review.UserID, &review.ProductID, &review.ParentID, &review.Rating,
			&review.Comment, &review.ISAnonymous, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	for rows3.Next() { //遍历所有获取的评论
		var review model.ShowReview
		err = rows3.Scan(&review.ID, &review.UserID, &review.ProductID, &review.ParentID, &review.Rating,
			&review.Comment, &review.ISAnonymous, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func DeleteReview(reviewID int) error {
	//1.先删除回复（如果有的话）
	query := "DELETE FROM reviews WHERE parent_id=?"
	_, err := Db.Exec(query, reviewID)
	if err != nil {
		return err
	}
	//2.再删除主评论
	query = "DELETE FROM reviews WHERE id=?"
	_, err = Db.Exec(query, reviewID)
	if err != nil {
		return err
	}
	return nil
}

func GetProductIDByReviewID(reviewID int) (int, error) {
	query := "SELECT product_id FROM reviews WHERE id=?"
	rows, err := Db.Query(query, reviewID)
	if err != nil {
		return 0, err
	}
	var productID int
	for rows.Next() {
		err = rows.Scan(&productID)
		if err != nil {
			return 0, err
		}
	}
	if productID == 0 {
		return 0, respond.ErrReviewNotExists
	}
	return productID, nil
}
