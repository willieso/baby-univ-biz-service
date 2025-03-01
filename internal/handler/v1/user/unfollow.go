package user

import (
	"github.com/willieso/baby-univ-biz-service/internal/ecode"
	"github.com/willieso/baby-univ-biz-service/internal/service"

	"github.com/gin-gonic/gin"

	"github.com/willieso/baby-univ-biz-service/pkg/errcode"
	"github.com/willieso/baby-univ-biz-service/pkg/log"
)

// Unfollow 取消关注
// @Summary 通过用户id取消关注用户
// @Description Get an user by user id
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param user_id body string true "用户id"
// @Success 200 {object} model.UserInfo "用户信息"
// @Router /users/unfollow [post]
func Unfollow(c *gin.Context) {
	var req FollowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Warnf("follow bind param err: %v", err)
		response.Error(c, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	// Get the user by the `user_id` from the database.
	_, err := service.Svc.Users().GetUserByID(c, req.UserID)
	if err != nil {
		response.Error(c, ecode.ErrUserNotFound.WithDetails(err.Error()))
		return
	}

	userID := service.GetUserID(c)
	// 不能关注自己
	if userID == req.UserID {
		response.Error(c, ecode.ErrUserNotFound)
		return
	}

	// 检查是否已经关注过
	isFollowed := service.Svc.Relations().IsFollowing(c.Request.Context(), userID, req.UserID)
	if !isFollowed {
		response.Error(c, errcode.Success)
		return
	}

	if isFollowed {
		// 取消关注
		err = service.Svc.Relations().Unfollow(c.Request.Context(), userID, req.UserID)
		if err != nil {
			log.Warnf("[follow] cancel user follow err: %v", err)
			response.Error(c, errcode.ErrInternalServer.WithDetails(err.Error()))
			return
		}
	}

	response.Success(c, nil)
}
