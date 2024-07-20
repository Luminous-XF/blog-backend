package database

import (
    model "blog-backend/app/model/post"
    "blog-backend/global"
    "github.com/google/uuid"
)

// GetPostById 通过帖子 id 查询帖子
func GetPostById(id uint) (post *model.Post, err error) {
    err = global.GDB.First(&post, id).Error
    return post, err
}

func GetPostByUUID(uuid uuid.UUID) (post *model.Post, err error) {
    err = global.GDB.Where("uuid = ?", uuid).First(&post).Error
    return post, err
}

// func GetPostCount() (count int64, err error) {
// 	err = global.GDB.Model(&model.Post{}).Count(&count).Error
// 	return count, err
// }

// GetPostList 按时间由近到远获取帖子列表
func GetPostList(offset, limit int) (postList []*model.Post, err error) {
    err = global.GDB.Order("create_time DESC").Limit(limit).Offset(offset).Find(&postList).Error
    return postList, err
}

// GetPostByAuthorId 根据 user_id 查询该用户发布的帖子列表
// func GetPostByAuthorId(userId int64) (postList []*model.Post, err error) {
// 	err = global.GDB.Where("user_id = ?", userId).Find(&postList).Error
// 	return postList, err
// }

// UpdatePost 修改帖子标题和内容
func UpdatePost(post *model.Post) (newPost *model.Post, err error) {
    err = global.GDB.Where("id = ?", post.ID).
        First(&newPost).Updates(model.Post{
        Title:   post.Title,
        Content: post.Content,
    }).Error
    return newPost, err
}
