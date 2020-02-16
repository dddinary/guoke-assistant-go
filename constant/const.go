package constant

const (

	ConfigFilePathDevelop 				= "develop.yaml"
	ConfigFilePathRelease 				= "release.yaml"

	RedisKeyLecture 					= "LectureData"
	RedisKeyPrefixToken 				= "Token:"

	PageSize							= 15
	PostMaxLen							= 300
	CommentMaxLen						= 300

	StudentStatusCommon 				= 0
	StudentStatusBlocked 				= 1

	PostKindAll							= 0 // 全部
	PostKindNotice						= 1 // 通知活动
	PostKindAskAndReply					= 2 // 果壳问问
	PostKindAnonymous 					= 3 // 匿名树洞
	PostKindFriends						= 4 // 约伴交友
	PostKindSecondHand					= 5 // 二手市场
	PostKindLostAndFound				= 6 // 失物招领

	NotificationKindLikePost			= 1
	NotificationKindCommentPost			= 2
	NotificationKindLikeComment			= 3
	NotificationKindCommentComment		= 4
	NotificationKindAdminDeletePost		= 11
	NotificationKindAdminDeleteComment	= 12
	NotificationKindAdminBlock			= 13
	NotificationKindAdminUnblock		= 14
	NotificationAdminSay				= 15

	NotificationStatusUnread			= 0
	NotificationStatusRead				= 1
	NotificationStatusDeleted			= 2

	ContextKeyUid						= "reqUid"
	ContextKeyBlocked					= "reqBlocked"
)
