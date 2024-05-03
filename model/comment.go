package model

import "time"

type Commnet struct {
	ID          uint      `json:"comment_id" gorm:"primaryKey"`
	CommentedBy string    `json:"commented_by"`                                     // user name who puts comment
	CreatedAt   time.Time `json:"created_on"`                                       // comment created on
	Description string    `json:"description"`                                      // commnet description
	IssueID     uint      `json:"issue_id" gorm:"foreignKey:IssueID,referenced:ID"` // ID fo the issue for which addec comment
	Issue       Issue     `json:"-"`
} // @name model.Comment
