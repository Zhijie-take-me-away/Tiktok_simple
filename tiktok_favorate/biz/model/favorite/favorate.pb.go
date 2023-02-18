// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: favorate.proto

package favorite

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	_ "tiktok_favorate/biz/model/api"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DouyinFavoriteActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      *string `protobuf:"bytes,1,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                                          // 用户鉴权token
	VideoId    *int64  `protobuf:"varint,2,req,name=video_id,json=videoId" json:"video_id,required" form:"video_id,required" query:"video_id,required"`                // 视频id
	ActionType *int32  `protobuf:"varint,3,req,name=action_type,json=actionType" json:"action_type,required" form:"action_type,required" query:"action_type,required"` // 1-点赞，2-取消点赞
}

func (x *DouyinFavoriteActionRequest) Reset() {
	*x = DouyinFavoriteActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteActionRequest) ProtoMessage() {}

func (x *DouyinFavoriteActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteActionRequest) Descriptor() ([]byte, []int) {
	return file_favorate_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinFavoriteActionRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *DouyinFavoriteActionRequest) GetVideoId() int64 {
	if x != nil && x.VideoId != nil {
		return *x.VideoId
	}
	return 0
}

func (x *DouyinFavoriteActionRequest) GetActionType() int32 {
	if x != nil && x.ActionType != nil {
		return *x.ActionType
	}
	return 0
}

type DouyinFavoriteActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
}

func (x *DouyinFavoriteActionResponse) Reset() {
	*x = DouyinFavoriteActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteActionResponse) ProtoMessage() {}

func (x *DouyinFavoriteActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favorate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteActionResponse) Descriptor() ([]byte, []int) {
	return file_favorate_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinFavoriteActionResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *DouyinFavoriteActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

type DouyinFavoriteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,required" form:"user_id,required" query:"user_id,required"` // 用户id
	Token  *string `protobuf:"bytes,2,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                      // 用户鉴权token
}

func (x *DouyinFavoriteListRequest) Reset() {
	*x = DouyinFavoriteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteListRequest) ProtoMessage() {}

func (x *DouyinFavoriteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_favorate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteListRequest.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteListRequest) Descriptor() ([]byte, []int) {
	return file_favorate_proto_rawDescGZIP(), []int{2}
}

func (x *DouyinFavoriteListRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *DouyinFavoriteListRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type DouyinFavoriteListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32   `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list" form:"video_list" query:"video_list"`                                  // 用户点赞视频列表
}

func (x *DouyinFavoriteListResponse) Reset() {
	*x = DouyinFavoriteListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteListResponse) ProtoMessage() {}

func (x *DouyinFavoriteListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_favorate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteListResponse.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteListResponse) Descriptor() ([]byte, []int) {
	return file_favorate_proto_rawDescGZIP(), []int{3}
}

func (x *DouyinFavoriteListResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *DouyinFavoriteListResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *DouyinFavoriteListResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                                                    // 视频唯一标识
	Author        *User   `protobuf:"bytes,2,req,name=author" json:"author,required" form:"author,required" query:"author,required"`                                                     // 视频作者信息
	PlayUrl       *string `protobuf:"bytes,3,req,name=play_url,json=playUrl" json:"play_url,required" form:"play_url,required" query:"play_url,required"`                                // 视频播放地址
	CoverUrl      *string `protobuf:"bytes,4,req,name=cover_url,json=coverUrl" json:"cover_url,required" form:"cover_url,required" query:"cover_url,required"`                           // 视频封面地址
	FavoriteCount *int64  `protobuf:"varint,5,req,name=favorite_count,json=favoriteCount" json:"favorite_count,required" form:"favorite_count,required" query:"favorite_count,required"` // 视频的点赞总数
	CommentCount  *int64  `protobuf:"varint,6,req,name=comment_count,json=commentCount" json:"comment_count,required" form:"comment_count,required" query:"comment_count,required"`      // 视频的评论总数
	IsFavorite    *bool   `protobuf:"varint,7,req,name=is_favorite,json=isFavorite" json:"is_favorite,required" form:"is_favorite,required" query:"is_favorite,required"`                // true-已点赞，false-未点赞
	Title         *string `protobuf:"bytes,8,req,name=title" json:"title,required" form:"title,required" query:"title,required"`                                                         // 视频标题
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_favorate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_favorate_proto_rawDescGZIP(), []int{4}
}

func (x *Video) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil && x.PlayUrl != nil {
		return *x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil && x.CoverUrl != nil {
		return *x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil && x.FavoriteCount != nil {
		return *x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil && x.CommentCount != nil {
		return *x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil && x.IsFavorite != nil {
		return *x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              *int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                                            // 用户id
	Name            *string `protobuf:"bytes,2,req,name=name" json:"name,required" form:"name,required" query:"name,required"`                                                     // 用户名称
	FollowCount     *int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount" json:"follow_count,omitempty" form:"follow_count" query:"follow_count"`                    // 关注总数
	FollowerCount   *int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount" json:"follower_count,omitempty" form:"follower_count" query:"follower_count"`          // 粉丝总数
	IsFollow        *bool   `protobuf:"varint,5,req,name=is_follow,json=isFollow" json:"is_follow,required" form:"is_follow,required" query:"is_follow,required"`                  // true-已关注，false-未关注
	Avatar          *string `protobuf:"bytes,6,opt,name=avatar" json:"avatar,omitempty" form:"avatar" query:"avatar"`                                                              //用户头像
	BackgroundImage *string `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage" json:"background_image,omitempty" form:"background_image" query:"background_image"` //用户个人页顶部大图
	Signature       *string `protobuf:"bytes,8,opt,name=signature" json:"signature,omitempty" form:"signature" query:"signature"`                                                  //个人简介
	TotalFavorited  *int64  `protobuf:"varint,9,opt,name=total_favorited,json=totalFavorited" json:"total_favorited,omitempty" form:"total_favorited" query:"total_favorited"`     //获赞数量
	WorkCount       *int64  `protobuf:"varint,10,opt,name=work_count,json=workCount" json:"work_count,omitempty" form:"work_count" query:"work_count"`                             //作品数量
	FavoriteCount   *int64  `protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount" json:"favorite_count,omitempty" form:"favorite_count" query:"favorite_count"`         //点赞数量
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_favorate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_favorate_proto_rawDescGZIP(), []int{5}
}

func (x *User) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil && x.FollowCount != nil {
		return *x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil && x.FollowerCount != nil {
		return *x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil && x.IsFollow != nil {
		return *x.IsFollow
	}
	return false
}

func (x *User) GetAvatar() string {
	if x != nil && x.Avatar != nil {
		return *x.Avatar
	}
	return ""
}

func (x *User) GetBackgroundImage() string {
	if x != nil && x.BackgroundImage != nil {
		return *x.BackgroundImage
	}
	return ""
}

func (x *User) GetSignature() string {
	if x != nil && x.Signature != nil {
		return *x.Signature
	}
	return ""
}

func (x *User) GetTotalFavorited() int64 {
	if x != nil && x.TotalFavorited != nil {
		return *x.TotalFavorited
	}
	return 0
}

func (x *User) GetWorkCount() int64 {
	if x != nil && x.WorkCount != nil {
		return *x.WorkCount
	}
	return 0
}

func (x *User) GetFavoriteCount() int64 {
	if x != nil && x.FavoriteCount != nil {
		return *x.FavoriteCount
	}
	return 0
}

var File_favorate_proto protoreflect.FileDescriptor

var file_favorate_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x1b, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5e, 0x0a, 0x1c, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x4a, 0x0a, 0x19, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x1a, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73,
	0x67, 0x12, 0x2e, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0xfa, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x02, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x02, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x08, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xe1,
	0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x18, 0x05, 0x20, 0x02, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0x88, 0x02, 0x0a, 0x0f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x7e, 0x0a, 0x0f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x2e, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x44, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0xd2, 0xc1, 0x18, 0x18, 0x2f, 0x64,
	0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x74, 0x69, 0x65, 0x2f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x12, 0x75, 0x0a, 0x0c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x2e, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0xca, 0xc1, 0x18, 0x16, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x42, 0x24, 0x5a,
	0x22, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x61, 0x74, 0x65,
	0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65,
}

var (
	file_favorate_proto_rawDescOnce sync.Once
	file_favorate_proto_rawDescData = file_favorate_proto_rawDesc
)

func file_favorate_proto_rawDescGZIP() []byte {
	file_favorate_proto_rawDescOnce.Do(func() {
		file_favorate_proto_rawDescData = protoimpl.X.CompressGZIP(file_favorate_proto_rawDescData)
	})
	return file_favorate_proto_rawDescData
}

var file_favorate_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_favorate_proto_goTypes = []interface{}{
	(*DouyinFavoriteActionRequest)(nil),  // 0: favorite.DouyinFavoriteActionRequest
	(*DouyinFavoriteActionResponse)(nil), // 1: favorite.DouyinFavoriteActionResponse
	(*DouyinFavoriteListRequest)(nil),    // 2: favorite.DouyinFavoriteListRequest
	(*DouyinFavoriteListResponse)(nil),   // 3: favorite.DouyinFavoriteListResponse
	(*Video)(nil),                        // 4: favorite.Video
	(*User)(nil),                         // 5: favorite.User
}
var file_favorate_proto_depIdxs = []int32{
	4, // 0: favorite.DouyinFavoriteListResponse.video_list:type_name -> favorite.Video
	5, // 1: favorite.Video.author:type_name -> favorite.User
	0, // 2: favorite.FavoriteHandler.favorite_action:input_type -> favorite.DouyinFavoriteActionRequest
	2, // 3: favorite.FavoriteHandler.FavoriteList:input_type -> favorite.DouyinFavoriteListRequest
	1, // 4: favorite.FavoriteHandler.favorite_action:output_type -> favorite.DouyinFavoriteActionResponse
	3, // 5: favorite.FavoriteHandler.FavoriteList:output_type -> favorite.DouyinFavoriteListResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_favorate_proto_init() }
func file_favorate_proto_init() {
	if File_favorate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_favorate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_favorate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_favorate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_favorate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_favorate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_favorate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_favorate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_favorate_proto_goTypes,
		DependencyIndexes: file_favorate_proto_depIdxs,
		MessageInfos:      file_favorate_proto_msgTypes,
	}.Build()
	File_favorate_proto = out.File
	file_favorate_proto_rawDesc = nil
	file_favorate_proto_goTypes = nil
	file_favorate_proto_depIdxs = nil
}
