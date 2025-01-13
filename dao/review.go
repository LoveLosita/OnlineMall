package dao

import "OnlineMall/model"

func RateAndReviewProduct(review model.AddReview) error {
	query := "INSERT INTO reviews(user_id,product_id,parent_id,rating,comment) VALUES(?,?,?,?,?)"
	_, err := Db.Exec(query, review.UserID, review.ProductID, review.ParentID, review.Rating, review.Comment)
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
