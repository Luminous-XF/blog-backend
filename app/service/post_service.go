package service

import (
    "blog-backend/app/common/error_code"
    "blog-backend/app/common/request"
    "blog-backend/app/common/response"
    "blog-backend/app/database"
)

func GetPostList(formData request.PageInfoRequest) (postResponseList []*response.PostResponse, code error_code.ErrorCode) {
    limit := formData.PageSize
    offset := formData.PageSize * (formData.Page - 1)
    postList, err := database.GetPostList(offset, limit)
    if err != nil {
        return nil, error_code.QueryPostListFailed
    }

    for _, post := range postList {
        postResponseList = append(postResponseList, &response.PostResponse{
            UUID:    post.UUID,
            Title:   post.Title,
            Excerpt: post.Excerpt,
            // Content:        post.Content, // 不用返回帖子内容, 帖子列表卡片信息中不需要
            Type:           post.Type,
            CommentCount:   post.CommentCount,
            Score:          post.Score,
            Status:         post.Status,
            Likes:          post.Likes,
            PageViews:      post.PageViews,
            HeaderImageUrl: post.HeaderImageUrl,
            CreateTime:     post.CreateTime,
            UpdateTime:     post.UpdateTime,
        })
    }

    return postResponseList, error_code.SUCCESS
}

func GetPostByUUID(formData request.GetByUUIDRequest) (postResponse *response.PostResponse, code error_code.ErrorCode) {
    uuid := formData.UUID

    post, err := database.GetPostByUUID(uuid)
    if err != nil {
        return nil, error_code.DatabaseError
    }

    postResponse = &response.PostResponse{
        UUID:           post.UUID,
        Title:          post.Title,
        Excerpt:        post.Excerpt,
        Content:        post.Content,
        Type:           post.Type,
        CommentCount:   post.CommentCount,
        Score:          post.Score,
        Status:         post.Status,
        Likes:          post.Likes,
        PageViews:      post.PageViews,
        HeaderImageUrl: post.HeaderImageUrl,
        CreateTime:     post.CreateTime,
        UpdateTime:     post.UpdateTime,
    }

    return postResponse, error_code.SUCCESS
}
