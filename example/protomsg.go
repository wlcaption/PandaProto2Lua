package protomsg

const (
	C2S_ACCNT_ENTER_GAME = 1001 // 进入游戏

	C2S_USER_HEAD_ID_CHANGE       = 1011 // 更改头像
	C2S_USER_HEAD_FRAME_ID_CHANGE = 1012 //更改头像框
	C2S_USER_USER_GUIDE           = 1021 // 新手引导
	C2S_USER_NAME_CHANGE          = 1031 // 更改名字
	C2S_USER_FIRM_NAME_CHANGE     = 1032 // 更改公司名字
	C2S_USER_GET_BRIEF_INFO       = 1041 // 获取玩家简明信息
	C2S_USER_SET_KEY_VALUE        = 1042 // key value 设置
	C2S_USER_GET_KEY_VALUES       = 1043 // 获取key values 信息

	C2S_GM_CMD_MGR  = 1051 // GM命令 - 物品管理
	C2S_GM_RECHARGE = 1052 // GM命令 - 充值

	C2S_USER_CHECK_SYS_REAL_TIME = 1061 // 校对系统时间

	C2S_USER_RECHARGE               = 1071 // 用户充值
	C2S_USER_PREORDER               = 1072 // 用户预充值
	C2S_USER_DELIVER_ORDER          = 1073 // 用户充值
	C2S_USER_QUERY_DELIVER_ORDER    = 1074 // 用户充值补单
	C2S_USER_PRIVILEDGE_CARD_SUBMIT = 1075 // 特权卡每日领取
	C2S_USER_WEEK_CARD_SUBMIT       = 1076 // 周卡每日领取
	C2S_USER_GET_RED_DOT_INFO       = 1077 // 部分红点需求拉取 for5点

	C2S_USER_TASK_DO            = 1101 // 开始普通/精英副本
	C2S_USER_TASK_LEAVE         = 1102 // 离开普通/精英副本
	C2S_USER_TASK_SWEEP         = 1111 // 扫荡普通/精英副本
	C2S_USER_TASK_BUY_ELITE_CNT = 1121 // 购买精英副本次数
	C2S_USER_TASK_REWARD_SUBMIT = 1131 // 领取任务奖励
	C2S_USER_TASK_LINE_SUBMIT   = 1136 // 提交副本主线

	C2S_USER_MALE_ROLE_TASK_DO            = 1151 // 开始人物副本
	C2S_USER_MALE_ROLE_TASK_LEAVE         = 1152 // 离开人物副本
	C2S_USER_MALE_ROLE_TASK_SWEEP         = 1161 // 扫荡人物副本
	C2S_USER_MALE_ROLE_TASK_BUY_CNT       = 1171 // 购买人物副本次数
	C2S_USER_MALE_ROLE_TASK_REWARD_SUBMIT = 1181 // 领取人物副本奖励
	C2S_USER_MALE_ROLE_TASK_LINE_SUBMIT   = 1186 // 提交人物副本主线

	C2S_USER_FIRM_PROP_VOUCHER_ARRANGE = 1201 // 安排学习券
	C2S_USER_FIRM_PROP_VOUCHER_SUBMIT  = 1202 // 完成学习券
	C2S_USER_FIRM_PROP_VOUCHER_CANCEL  = 1203 // 取消学习券
	C2S_USER_FIRM_PROP_VOUCHER_QUICK   = 1204 // 加速学习券

	C2S_USER_GOODS_SELL             = 1301 // 出售道具
	C2S_USER_GOODS_EXCHANGE         = 1321 // 道具兑换
	C2S_USER_GOODS_CONSUME          = 1341 // 道具使用
	C2S_USER_GOODS_EXCHANGE_REFRESH = 1342 // 道具兑换刷新

	C2S_USER_MAIL_BRIEF_INFOS                    = 1402 // 获取邮件简明信息
	C2S_USER_MAIL_DETAIL_INFO                    = 1403 // 获取邮件详情
	C2S_USER_MAIL_RECV_REWARDS                   = 1404 // 领取邮件奖励
	C2S_USER_MAIL_FRIEND_MAIL_RECV_STATUS_CHANGE = 1411 // 好友邮件接受状态变更

	C2S_USER_CARD_UPGRADE                = 1501 // 卡牌升级
	C2S_USER_CARD_DEPLOY                 = 1502 // 卡牌装备
	C2S_USER_CARD_RESOLVE                = 1503 // 卡牌分解
	C2S_USER_CARD_ADVANCE                = 1504 // 卡牌进阶
	C2S_USER_CARD_DESIGN                 = 1505 // 卡牌设计
	C2S_USER_CARD_EVOLUTION              = 1507 // 卡牌进化
	C2S_USER_CARD_SKILL_UPGRADE          = 1508 // 卡牌技能升级
	C2S_USER_CARD_ADVANCE_REWARDS_SUBMIT = 1509 // 卡牌进阶奖励领取

	C2S_USER_CHECK_IN_SUBMIT         = 1701 // 提交签到奖励
	C2S_USER_CHECK_IN_REWARD_SUBMIT  = 1702 // 领取签到奖励
	C2S_USER_CHECK_IN_APPEND         = 1703 // 补签签到奖励
	C2S_USER_CHECK_IN_MONTHLY_REWARD = 1704 // 领取签到月奖励
	C2S_USER_CHECK_IN_GET_INFO       = 1705 // 获取当前签到信息

	C2S_USER_HEART_BUY     = 1711 // 购买体力
	C2S_USER_GOLD_BUY      = 1712 // 购买金币
	C2S_USER_STAR_EXCHANGE = 1721 // 星星兑换
	C2S_USER_FILM_BUY      = 1722 // 购买胶片

	C2S_USER_CARD_CALL_SUBMIT              = 1801 // 抽卡
	C2S_USER_ACHIEVEMENT_SUBMIT            = 1811 // 成就奖励
	C2S_USER_DAILY_QUEST_SUBMIT            = 1821 // 日常任务奖励
	C2S_USER_LIVENESS_REWARD_SUBMIT        = 1822 // 领取活跃度奖励
	C2S_USER_WEEKLY_LIVENESS_REWARD_SUBMIT = 1823 // 领取活跃度奖励

	C2S_USER_REDEEM_CODE_SUBMIT = 1831 // 兑换码奖励领取

	C2S_USER_RANK_TASK_LIST     = 1901 // 任务排行榜
	C2S_USER_RANK_CARD_CNT_LIST = 1902 // 卡牌数量排行榜

	C2S_USER_RANK_ARENA_LIST     = 1921 // 获取竞技场排行榜
	C2S_USER_ARENA_CARD_DEPLOY   = 1922 // 竞技场卡牌上阵
	C2S_USER_ARENA_GET_PEER      = 1923 // 获取竞技场对手
	C2S_USER_ARENA_PVP_BATTLE    = 1924 // 挑战对手
	C2S_USER_ARENA_RANK_INFOS    = 1925 // 获取竞技场奖励列表
	C2S_USER_ARENA_REWARD_SUBMIT = 1926 // 领取竞技场奖励
	C2S_USER_ARENA_BUY_CNT       = 1927 // 购买竞技场次数

	/* v2 arena */
	C2S_USER_ARENA_V2_RANK_ARENA_LIST = 1931 // 获取竞技场排行榜
	C2S_USER_ARENA_V2_CARD_DEPLOY     = 1932 // 竞技场卡牌上阵
	C2S_USER_ARENA_REFRESH_PEER       = 1933 // 刷新竞技场对手
	C2S_USER_ARENA_V2_PVP_BATTLE      = 1934 // 挑战对手
	C2S_USER_ARENA_V2_REWARD_SUBMIT   = 1936 // 领取竞技场奖励
	C2S_USER_ARENA_V2_BUY_CNT         = 1937 // 购买竞技场次数
	C2S_USER_ARENA_MAIN_INFO          = 1938 // 获取竞技场界面信息

	C2S_USER_ACTIVITY_GET_INFOS                    = 2005 // 获取活动详细信息
	C2S_USER_ACTIVITY_DAILY_REWARD_SUBMIT          = 2011 // 领取每日活动奖励  //2017-10-31  现在用这个每日活动
	C2S_USER_ACTIVITY_GROUP_REWARD_SUBMIT          = 2021 // 领取组活动奖励
	C2S_USER_ACTIVITY_CARD_CALL_SUBMIT             = 2031 // 活动抽卡
	C2S_USER_ACTIVITY_TASK_DO                      = 2041 // 活动开始任务
	C2S_USER_ACTIVITY_TASK_LEAVE                   = 2042 // 活动离开任务
	C2S_USER_ACTIVITY_TASK_SWEEP                   = 2043 // 活动扫荡任务
	C2S_USER_ACTIVITY_TASK_BUY_CNT                 = 2044 // 活动购买副本次数
	C2S_USER_ACTIVITY_TASK_LINE_SUBMIT             = 2045 // 活动提交副本主线
	C2S_USER_ACTIVITY_TASK_DAILY_REWARD_SUBMIT     = 2046 // 活动中附带的每日奖励,每天只能领取固定奖励,没有领取不补发
	C2S_USER_ACTIVITY_TIME_REWARD_SUBMIT           = 2051 // 领取定时活动奖励
	C2S_USER_ACTIVITY_DAILY_EX_REWARD_SUBMIT       = 2061 // 领取每日活动奖励
	C2S_USER_ACTIVITY_MAIL_REWARD_SUBMIT           = 2071 // 领取邮件活动奖励
	C2S_USER_ACTIVITY_GET_PACK_INFOS               = 2081 // 获取礼包活动详细信息
	C2S_USER_ACTIVITY_BUY_PACK                     = 2082 // 购买礼包
	C2S_USER_ACTIVITY_COMMON_REWARD_SUBMIT         = 2083 // 领取通用活动奖励
	C2S_USER_ACTIVITY_SHARE_SUBMIT                 = 2084 // 提交分享结果
	C2S_USER_ACTIVITY_SHARE_REWARD_SUBMIT          = 2085 // 领取分享奖励
	C2S_USER_ACTIVITY_COMMENT_SUBMIT               = 2086 // 好评引导活动提交
	C2S_USER_ACTIVITY_GOODS_EXCHANGE_RANK          = 2087 // 生日榜
	C2S_USER_ACTIVITY_GOODS_EXCHANGE_REWARD_SUBMIT = 2088 // 生日榜奖励领取
	C2S_USER_ACTIVITY_INVITE_SUBMIT                = 2089 // 生日邀请

	/* 充值类活动 */
	C2S_USER_ACTIVITY_FUND_REWARD_SUBMIT             = 2091 // 领取基金活动奖励
	C2S_USER_ACTIVITY_RECHARGE_REWARD_SUBMIT         = 2092 // 领取累充活动奖励
	C2S_USER_ACTIVITY_FIRSTCHARGE_PACK_REWARD_SUBMIT = 2093 // 领取首充礼包奖励

	C2S_USER_ACTIVITY_ARENA_QUEST_SUBMIT           = 2095 // 领取竞技场任务奖励
	C2S_USER_ACTIVITY_ARENA_EXCHANGE_SUBMIT        = 2096 // 领取竞技场兑换奖励
	C2S_USER_ACTIVITY_LOOP_QUEST_SUBMIT            = 2097 // 领取跑环掉落奖励
	C2S_USER_ACTIVITY_SPRING_FORTUNE_GET_SUBMIT    = 2098 // 获取今日签
	C2S_USER_ACTIVITY_SPRING_FORTUNE_REWARD_SUBMIT = 2099 // 获取签到奖励 (红包/签到)
	C2S_USER_ACTIVITY_SPRING_FORTUNE_SHARE_SUBMIT  = 2090 // 春节红包分享
	/* 充值类活动 */

	C2S_USER_SECRET_TASK_START         = 2101 // 暗访副本开始
	C2S_USER_SECRET_TASK_RESET         = 2102 // 暗访副本重置
	C2S_USER_SECRET_TASK_CARD_CHANGE   = 2103 // 暗访副本更换卡牌
	C2S_USER_SECRET_TASK_CARD_DEPLOY   = 2104 // 暗访副本上阵卡牌
	C2S_USER_SECRET_TASK_DO            = 2111 // 暗访副本执行
	C2S_USER_SECRET_TASK_LEAVE         = 2112 // 暗访副本离开
	C2S_USER_SECRET_TASK_REWARD_SUBMIT = 2121 // 暗访副本奖励领取

	C2S_USER_STAFF_REFRESH = 2201 // 刷新职员商店
	C2S_USER_STAFF_HIRE    = 2211 // 雇佣职员
	C2S_USER_STAFF_FIRE    = 2212 // 解雇职员
	C2S_USER_STAFF_UPGRADE = 2213 // 升级职员
	C2S_USER_STAFF_DEPLOY  = 2214 // 特聘职员

	C2S_USER_SHORT_MSG_GET_HISTORY_LIST    = 2301 // 获取短信消息历史列表
	C2S_USER_SHORT_MSG_GET_HISTORY_MSG     = 2302 // 获取历史短信消息
	C2S_USER_SHORT_MSG_GET_SEND_LIST       = 2303 // 获取可发送短信消息列表
	C2S_USER_SHORT_MSG_GET_RECV_MSGS       = 2311 // 获取最新短信消息
	C2S_USER_SHORT_MSG_SEND_MSG            = 2312 // 发送短信消息
	C2S_USER_SHORT_MSG_REPLY_MSG           = 2313 // 回复短信消息
	C2S_USER_SHORT_MSG_OPTION_PARAM_CHANGE = 2314 // 变更短信消息选项参数

	C2S_USER_PHONE_MSG_GET_RECV_MSGS   = 2351 // 获取电话消息列表
	C2S_USER_PHONE_MSG_SET_MSG_STATUS  = 2352 // 变更电话消息状态
	C2S_USER_PHONE_MSG_GET_HISTORY_MSG = 2353 // 获取历史电话消息
	C2S_USER_PHONE_MSG_REPLY_MSG       = 2354 // 回复电话消息

	C2S_USER_PUBLIC_MSG_GET_HISTORY_LIST = 2401 // 获取公众号消息历史列表
	C2S_USER_PUBLIC_MSG_GET_RECV_MSGS    = 2402 // 获取最新公众号消息
	C2S_USER_PUBLIC_MSG_SET_MSG_STATUS   = 2403 // 变更公众号消息状态

	C2S_USER_FRIEND_MSG_GET_RECV_MSGS  = 2451 // 获取朋友圈消息列表
	C2S_USER_FRIEND_MSG_GET_SEND_LIST  = 2452 // 获取可发送朋友圈消息列表
	C2S_USER_FRIEND_MSG_SEND_MSG       = 2461 // 发送朋友圈消息
	C2S_USER_FRIEND_MSG_REPLY_MSG      = 2462 // 回复朋友圈消息
	C2S_USER_FRIEND_MSG_SET_MSG_STATUS = 2463 // 变更朋友圈消息状态

	C2S_USER_MALE_ROLE_MODIFY_NOTE         = 2501 // 男主修改备注
	C2S_USER_MALE_ROLE_CHANGE_RELATED_CARD = 2502 // 男主修改关联卡牌
	C2S_USER_MALE_ROLE_RECOVER_NOTE        = 2503 // 男主备注还原

	C2S_USER_ENGAGEMENT_SUBMIT  = 2551 // 完成约会
	C2S_USER_EXTRA_STORY_SUBMIT = 2552 // 解锁小传

	C2S_USER_TRACK_TASK_CARD_DEPLOY   = 2601 // 跟踪副本上阵卡牌
	C2S_USER_TRACK_TASK_GENERATE      = 2602 // 跟踪副本生成副本
	C2S_USER_TRACK_TASK_DO            = 2611 // 跟踪副本执行
	C2S_USER_TRACK_TASK_LEAVE         = 2612 // 跟踪副本离开
	C2S_USER_TRACK_TASK_DROP_SUBMIT   = 2613 // 跟踪副本掉落奖励领取
	C2S_USER_TRACK_TASK_REWARD_SUBMIT = 2621 // 跟踪副本奖励领取

	C2S_USER_FRIEND_APPLY_INFOS      = 2701 // 获取申请好友列表
	C2S_USER_FRIEND_APPLY_FRIEND     = 2702 // 申请好友
	C2S_USER_FRIEND_APPLY_MANAGE     = 2703 // 好友申请处理
	C2S_USER_FRIEND_FRIEND_INFOS     = 2751 // 获取好友列表
	C2S_USER_FRIEND_REMOVE_FRIEND    = 2752 // 删除好友
	C2S_USER_FRIEND_RECOMMEND_FRIEND = 2761 // 推荐好友
	C2S_USER_FRIEND_SEND_MAIL        = 2771 // 发送好友邮件
	C2S_USER_FRIEND_BATCH_SEND_MAIL  = 2772 // 发送好友邮件

	C2S_USER_LOOP_TASK_STATUS          = 2801 // 跑环任务状态
	C2S_USER_LOOP_TASK_SUBMIT          = 2802 // 跑环任务提交
	C2S_USER_LOOP_TASK_LOOP_RWD_SUBMIT = 2803 // 跑环任务领取奖励
	C2S_USER_LOOP_TASK_LOOP_RESET      = 2804 // 跑环任务重置
	C2S_USER_LOOP_TASK_LOOP_SWEEP      = 2805 // 跑环扫荡

	C2S_USER_RES_DUNGEON_INFO                = 2810 // 获取资源副本面板信息
	C2S_USER_GET_RES_DUNGEON_THEME_REWARD    = 2811 // 领取资源副本奖励
	C2S_USER_RES_DUNGEON_THEME_BATTLE        = 2812 // 资源副本战斗
	C2S_USER_RES_DUNGEON_THEME_RESET         = 2813 // 资源副本重置
	C2S_USER_RES_DUNGEON_UPDATE_DEFENCE      = 2814 // 资源副本上阵守护
	C2S_USER_RES_DUNGEON_THEME_BATTLE_GOBACK = 2815 // 资源副本战斗回撤
	C2S_USER_RES_DUNGEON_SWEEP               = 2816 // 资源副本扫荡
	C2S_USER_GET_RES_DUNGEON_THEME_REWARDS   = 2817 // 领取资源副本奖励

	C2S_USER_STONE_RANDOM = 2820 // 随机石头

	C2S_USER_BIRTH_EXAM_SET           = 2821 // 生日问卷设置
	C2S_USER_BIRTH_REWARD_SUBMIT      = 2822 //生日奖励领取
	C2S_USER_BIRTH_EXAM_REWARD_SUBMIT = 2823 //生日问卷奖励领取
	C2S_USER_BIRTH_DISPLAY_REWARD     = 2824 //生日问卷奖励查看

	C2S_USER_RES_DUNGEON_BOX_DISPATCH = 2830 //资源副本 派遣宝箱
	C2S_USER_RES_DUNGEON_BOX_SPEEDUP  = 2831 //资源副本 加速宝箱
	C2S_USER_RES_DUNGEON_BOX_CANCEL   = 2832 //资源副本 取消派遣
	C2S_USER_RES_DUNGEON_BOX_FINISH   = 2833 //资源副本 完成派遣

	C2S_USER_ACTIVITY_WATER_PLANT_SUBMIT           = 4001 // 浇花活动浇水 施肥请求
	C2S_USER_ACTIVITY_WATER_PLANT_REWARD_SUBMIT    = 4002 // 浇花活动活动奖励（羁绊）
	C2S_USER_ACTIVITY_GAIN_CONSUME_REWARD_SUBMIT   = 4003 // 日常/累计消耗活动
	C2S_USER_ACTIVITY_DEFEND_TASK_BATTLE           = 4004 // 保卫活动 战斗
	C2S_USER_ACTIVITY_DEFEND_TASKLINE_SUBMIT       = 4005 // 保卫活动 提交活动任务主线
	C2S_USER_ACTIVITY_DEFEND_TASK_REWARD_SUBMIT    = 4006 // 保卫活动 活动奖励领取
	C2S_USER_ACTIVITY_ANSWER_EXAM_SUBMIT           = 4007 // 答题协议
	C2S_USER_ACTIVITY_REVIEW_EXAM_SUBMIT           = 4008 // 复习协议
	C2S_USER_ACTIVITY_EVENT_STORY_SUBMIT           = 4009 // 剧情事件请求
	C2S_USER_ACTIVITY_EVENT_TUJIAN_SUBMIT          = 4010 // 事件3 图鉴展示, 答题
	C2S_USER_ACTIVITY_EVENT_TUJIAN_GETREWARD       = 4011 // 事件3 获取奖励
	C2S_USER_ACTIVITY_OBTAIN_BUFF_REWARD_SUBMIT    = 4012 // 事件3 领取Buff奖励
	C2S_USER_ACTIVITY_SINGLE_BUY_GOODS             = 4013 // 单人卡池商店购买
	C2S_USER_ACTIVITY_RECALL_BIND_CODE             = 4020 // 绑定code
	C2S_USER_ACTIVITY_RECALL_GET_USER_BRIEF_INFO   = 4021 // 获取玩家基本信息
	C2S_USER_ACTIVITY_RECALL_GET_LOGS              = 4022 // 获取召回积分日志
	C2S_USER_ACTIVITY_RECALL_SHARE                 = 4023 // 召回分享
	C2S_USER_ACTIVITY_PRIVILEGE_CARD_REWARD_SUBMIT = 4024 // 特权卡活动奖励领取

	C2S_USER_ACTIVITY_ANSWER_EXAM3_SUBMIT        = 4025 // 答题3协议
	C2S_USER_ACTIVITY_REVIEW_EXAM3_SUBMIT        = 4026 // 答题3复习协议
	C2S_USER_ACTIVITY_ANSWER_EXAM3_REWARD_SUBMIT = 4027 // 答题3领取羁绊奖励
	C2S_USER_ACTIVITY_MAIN_QUEST_WISH            = 4028 // 主线卡池许愿
	C2S_USER_ACTIVITY_MAIN_QUEST_EXCHANGE        = 4029 // 主线卡池兑换
	C2S_USER_ACTIVITY_MAIN_QUEST_INFO            = 4030 // 主线卡信息
)

const (
	S2C_SYS_REAL_TIME = 1 // 服务器与客户端同步时间
	S2C_SYS_MSG       = 2 // 系统广播消息

	S2C_ACCNT_ENTER_GAME_RET = 1001 // 进入游戏返回
	S2C_ACCNT_KICK_OUT       = 1002 // 账号踢下线

	S2C_USER_HEAD_ID_CHANGE_RET       = 1011 // 更改头像
	S2C_USER_HEAD_FRAME_ID_CHANGE_RET = 1012 // 更改头像框
	S2C_USER_USER_GUIDE_RET           = 1021 // 新手引导
	S2C_USER_NAME_CHANGE_RET          = 1031 // 更改名字
	S2C_USER_FIRM_NAME_CHANGE_RET     = 1032 // 更改公司名字
	S2C_USER_GET_BRIEF_INFO_RET       = 1041 // 获取玩家简明信息
	S2C_USER_SET_KEY_VALUE_RET        = 1042 // key value 设置
	S2C_USER_GET_KEY_VALUES_RET       = 1043 // 获取key values 信息

	S2C_GM_CMD_MGR_RET  = 1051 // GM命令 - 物品管理
	S2C_GM_RECHARGE_RET = 1052 // GM命令 - 充值

	S2C_USER_CHECK_SYS_REAL_TIME_RET = 1061 // 校对系统时间

	S2C_USER_RECHARGE_RET               = 1071 // 用户充值
	S2C_USER_PREORDER_RET               = 1072 // 用户预充值
	S2C_USER_DELIVER_ORDER_RET          = 1073 // 用户充值
	S2C_USER_QUERY_DELIVER_ORDER_RET    = 1074 // 用户充值补单
	S2C_USER_PRIVILEDGE_CARD_SUBMIT_RET = 1075 // 特权卡每日领取
	S2C_USER_WEEK_CARD_SUBMIT_RET       = 1076 // 周卡每日领取
	S2C_USER_GET_RED_DOT_INFO_RET       = 1077 // 部分红点逻辑

	S2C_USER_TASK_DO_RET            = 1101 // 开始普通/精英副本
	S2C_USER_TASK_LEAVE_RET         = 1102 // 离开普通/精英副本
	S2C_USER_TASK_SWEEP_RET         = 1111 // 扫荡普通/精英副本
	S2C_USER_TASK_BUY_ELITE_CNT_RET = 1121 // 购买精英副本次数
	S2C_USER_TASK_REWARD_SUBMIT_RET = 1131 // 领取任务奖励
	S2C_USER_TASK_LINE_SUBMIT_RET   = 1136 // 提交副本主线

	S2C_USER_MALE_ROLE_TASK_DO_RET            = 1151 // 开始人物副本
	S2C_USER_MALE_ROLE_TASK_LEAVE_RET         = 1152 // 离开人物副本
	S2C_USER_MALE_ROLE_TASK_SWEEP_RET         = 1161 // 扫荡人物副本
	S2C_USER_MALE_ROLE_TASK_BUY_CNT_RET       = 1171 // 购买人物副本次数
	S2C_USER_MALE_ROLE_TASK_REWARD_SUBMIT_RET = 1181 // 领取人物副本奖励
	S2C_USER_MALE_ROLE_TASK_LINE_SUBMIT_RET   = 1186 // 提交人物副本主线

	S2C_USER_FIRM_PROP_VOUCHER_ARRANGE_RET = 1201 // 安排学习券
	S2C_USER_FIRM_PROP_VOUCHER_SUBMIT_RET  = 1202 // 完成学习券
	S2C_USER_FIRM_PROP_VOUCHER_CANCEL_RET  = 1203 // 取消学习券
	S2C_USER_FIRM_PROP_VOUCHER_QUICK_RET   = 1204 // 加速学习券

	S2C_USER_GOODS_SELL_RET             = 1301 // 出售道具
	S2C_USER_GOODS_EXCHANGE_RET         = 1321 // 道具兑换
	S2C_USER_GOODS_CONSUME_RET          = 1341 // 道具使用
	S2C_USER_GOODS_EXCHANGE_REFRESH_RET = 1342 // 道具兑换刷新

	S2C_USER_MAIL_BRIEF_INFOS_RET                    = 1402 // 获取邮件信息
	S2C_USER_MAIL_DETAIL_INFO_RET                    = 1403 // 获取邮件详情
	S2C_USER_MAIL_RECV_REWARDS_RET                   = 1404 // 领取邮件奖励
	S2C_USER_MAIL_FRIEND_MAIL_RECV_STATUS_CHANGE_RET = 1411 // 好友邮件接受状态变更

	S2C_USER_CARD_UPGRADE_RET                = 1501 // 卡牌升级
	S2C_USER_CARD_DEPLOY_RET                 = 1502 // 卡牌装备
	S2C_USER_CARD_RESOLVE_RET                = 1503 // 卡牌分解
	S2C_USER_CARD_ADVANCE_RET                = 1504 // 卡牌进阶
	S2C_USER_CARD_DESIGN_RET                 = 1505 // 卡牌设计
	S2C_USER_CARD_EVOLUTION_RET              = 1507 // 卡牌进化
	S2C_USER_CARD_SKILL_UPGRADE_RET          = 1508 // 卡牌技能升级
	S2C_USER_CARD_ADVANCE_REWARDS_SUBMIT_RET = 1509 // 卡牌进阶奖励领取

	S2C_USER_CHECK_IN_SUBMIT_RET         = 1701 // 提交签到奖励
	S2C_USER_CHECK_IN_REWARD_SUBMIT_RET  = 1702 // 领取签到奖励
	S2C_USER_CHECK_IN_APPEND_RET         = 1703 // 补签签到奖励
	S2C_USER_CHECK_IN_MONTHLY_REWARD_RET = 1704 // 领取签到月奖励
	S2C_USER_CHECK_IN_GET_INFO_RET       = 1705 // 获取签到信息

	S2C_USER_HEART_BUY_RET     = 1711 // 购买体力
	S2C_USER_GOLD_BUY_RET      = 1712 // 购买金币
	S2C_USER_STAR_EXCHANGE_RET = 1721 // 星星兑换
	S2C_USER_FILM_BUY_RET      = 1722 // 购买胶片

	S2C_USER_CARD_CALL_SUBMIT_RET              = 1801 // 抽卡
	S2C_USER_ACHIEVEMENT_SUBMIT_RET            = 1811 // 成就奖励
	S2C_USER_DAILY_QUEST_SUBMIT_RET            = 1821 // 日常任务奖励
	S2C_USER_LIVENESS_REWARD_SUBMIT_RET        = 1822 // 领取活跃度奖励
	S2C_USER_WEEKLY_LIVENESS_REWARD_SUBMIT_RET = 1823 // 领取周活跃度奖励

	S2C_USER_REDEEM_CODE_SUBMIT_RET = 1831 // 兑换码奖励领取

	S2C_USER_MAIN_LINE_REWARD_SUBMIT_RET = 1851 // 领取主线奖励

	S2C_USER_RANK_TASK_LIST_RET     = 1901 // 任务排行榜
	S2C_USER_RANK_CARD_CNT_LIST_RET = 1902 // 卡牌数量排行榜

	S2C_USER_RANK_ARENA_LIST_RET     = 1921 // 获取竞技场排行榜
	S2C_USER_ARENA_CARD_DEPLOY_RET   = 1922 // 竞技场卡牌上阵
	S2C_USER_ARENA_GET_PEER_RET      = 1923 // 获取竞技场对手
	S2C_USER_ARENA_PVP_BATTLE_RET    = 1924 // 挑战对手
	S2C_USER_ARENA_RANK_INFOS_RET    = 1925 // 获取竞技场排行信息
	S2C_USER_ARENA_REWARD_SUBMIT_RET = 1926 // 领取竞技场奖励
	S2C_USER_ARENA_BUY_CNT_RET       = 1927 // 购买竞技场次数
	/* v2 arena */
	S2C_USER_ARENA_V2_RANK_ARENA_LIST_RET = 1931 // 获取竞技场排行榜
	S2C_USER_ARENA_V2_CARD_DEPLOY_RET     = 1932 // 竞技场卡牌上阵
	S2C_USER_ARENA_REFRESH_PEER_RET       = 1933 // 刷新竞技场对手
	S2C_USER_ARENA_V2_PVP_BATTLE_RET      = 1934 // 挑战对手
	S2C_USER_ARENA_V2_REWARD_SUBMIT_RET   = 1936 // 领取竞技场奖励
	S2C_USER_ARENA_V2_BUY_CNT_RET         = 1937 // 购买竞技场次数
	S2C_USER_ARENA_MAIN_INFO_RET          = 1938 // 获取竞技场界面信息

	S2C_USER_ACTIVITY_GET_INFOS_RET                    = 2005 // 获取活动详细静态信息
	S2C_USER_ACTIVITY_DAILY_REWARD_SUBMIT_RET          = 2011 // 领取每日活动奖励
	S2C_USER_ACTIVITY_GROUP_REWARD_SUBMIT_RET          = 2021 // 领取组活动奖励
	S2C_USER_ACTIVITY_CARD_CALL_SUBMIT_RET             = 2031 // 活动抽卡
	S2C_USER_ACTIVITY_TASK_DO_RET                      = 2041 // 活动开始任务
	S2C_USER_ACTIVITY_TASK_LEAVE_RET                   = 2042 // 活动离开任务
	S2C_USER_ACTIVITY_TASK_SWEEP_RET                   = 2043 // 活动扫荡任务
	S2C_USER_ACTIVITY_TASK_BUY_CNT_RET                 = 2044 // 活动购买精英副本次数
	S2C_USER_ACTIVITY_TASK_LINE_SUBMIT_RET             = 2045 // 活动提交副本主线
	S2C_USER_ACTIVITY_TASK_DAILY_REWARD_SUBMIT_RET     = 2046 // 活动中附带的每日奖励,每天只能领取固定奖励,返回结果
	S2C_USER_ACTIVITY_TIME_REWARD_SUBMIT_RET           = 2051 // 领取定时活动奖励
	S2C_USER_ACTIVITY_DAILY_EX_REWARD_SUBMIT_RET       = 2061 // 领取每日活动奖励
	S2C_USER_ACTIVITY_MAIL_REWARD_SUBMIT_RET           = 2071 // 领取邮件活动奖励
	S2C_USER_ACTIVITY_GET_PACK_INFOS_RET               = 2081 // 获取礼包活动详细信息
	S2C_USER_ACTIVITY_BUY_PACK_RET                     = 2082 // 购买礼包
	S2C_USER_ACTIVITY_COMMON_REWARD_SUBMIT_RET         = 2083 // 领取通用活动奖励
	S2C_USER_ACTIVITY_SHARE_SUBMIT_RET                 = 2084 // 提交分享结果
	S2C_USER_ACTIVITY_SHARE_REWARD_SUBMIT_RET          = 2085 // 领取分享奖励
	S2C_USER_ACTIVITY_COMMENT_SUBMIT_RET               = 2086 // 好评引导活动结果
	S2C_USER_ACTIVITY_GOODS_EXCHANGE_RANK_RET          = 2087 // 生日榜
	S2C_USER_ACTIVITY_GOODS_EXCHANGE_REWARD_SUBMIT_RET = 2088 // 生日榜奖励领取
	S2C_USER_ACTIVITY_INVITE_SUBMIT_RET                = 2089 // 生日邀请

	/* 充值类活动 */
	S2C_USER_ACTIVITY_FUND_REWARD_SUBMIT_RET             = 2091 // 领取基金奖励
	S2C_USER_ACTIVITY_RECHARGE_REWARD_SUBMIT_RET         = 2092 // 领取累充奖励
	S2C_USER_ACTIVITY_FIRSTCHARGE_PACK_REWARD_SUBMIT_RET = 2093 // 领取首充礼包奖励

	S2C_USER_ACTIVITY_ARENA_QUEST_SUBMIT_RET           = 2095 // 领取竞技场任务奖励
	S2C_USER_ACTIVITY_ARENA_EXCHANGE_SUBMIT_RET        = 2096 // 领取竞技场兑换奖励
	S2C_USER_ACTIVITY_LOOP_QUEST_SUBMIT_RET            = 2097 // 领取跑环掉落活动奖励
	S2C_USER_ACTIVITY_SPRING_FORTUNE_GET_SUBMIT_RET    = 2098 // 获取今日签
	S2C_USER_ACTIVITY_SPRING_FORTUNE_REWARD_SUBMIT_RET = 2099 // 获取签到奖励 (红包/签到)
	S2C_USER_ACTIVITY_SPRING_FORTUNE_SHARE_SUBMIT_RET  = 2090 // 春节分享
	/* 充值类活动 */

	S2C_USER_SECRET_TASK_START_RET         = 2101 // 暗访副本开始
	S2C_USER_SECRET_TASK_RESET_RET         = 2102 // 暗访副本重置
	S2C_USER_SECRET_TASK_CARD_CHANGE_RET   = 2103 // 暗访副本更换卡牌
	S2C_USER_SECRET_TASK_CARD_DEPLOY_RET   = 2104 // 暗访副本上阵卡牌
	S2C_USER_SECRET_TASK_DO_RET            = 2111 // 暗访副本执行
	S2C_USER_SECRET_TASK_LEAVE_RET         = 2112 // 暗访副本离开
	S2C_USER_SECRET_TASK_REWARD_SUBMIT_RET = 2121 // 暗访副本奖励领取

	S2C_USER_STAFF_REFRESH_RET = 2201 // 刷新职员商店
	S2C_USER_STAFF_HIRE_RET    = 2211 // 雇佣职员
	S2C_USER_STAFF_FIRE_RET    = 2212 // 解雇职员
	S2C_USER_STAFF_UPGRADE_RET = 2213 // 升级职员
	S2C_USER_STAFF_DEPLOY_RET  = 2214 // 特聘职员

	S2C_USER_SHORT_MSG_GET_HISTORY_LIST_RET    = 2301 // 获取短信消息历史列表
	S2C_USER_SHORT_MSG_GET_HISTORY_MSG_RET     = 2302 // 获取历史短信消息
	S2C_USER_SHORT_MSG_GET_SEND_LIST_RET       = 2303 // 获取可发送短信消息列表
	S2C_USER_SHORT_MSG_GET_RECV_MSGS_RET       = 2311 // 获取最新短信消息
	S2C_USER_SHORT_MSG_SEND_MSG_RET            = 2312 // 发送短信消息
	S2C_USER_SHORT_MSG_REPLY_MSG_RET           = 2313 // 回复短信消息
	S2C_USER_SHORT_MSG_OPTION_PARAM_CHANGE_RET = 2314 // 变更短信消息选项参数

	S2C_USER_PHONE_MSG_GET_RECV_MSGS_RET   = 2351 // 获取电话消息列表
	S2C_USER_PHONE_MSG_SET_MSG_STATUS_RET  = 2352 // 变更电话消息状态
	S2C_USER_PHONE_MSG_GET_HISTORY_MSG_RET = 2353 // 获取历史电话消息
	S2C_USER_PHONE_MSG_REPLY_MSG_RET       = 2354 // 回复电话消息

	S2C_USER_PUBLIC_MSG_GET_HISTORY_LIST_RET = 2401 // 获取公众号消息历史列表
	S2C_USER_PUBLIC_MSG_GET_RECV_MSGS_RET    = 2402 // 获取最新公众号消息
	S2C_USER_PUBLIC_MSG_SET_MSG_STATUS_RET   = 2403 // 变更公众号消息状态

	S2C_USER_FRIEND_MSG_GET_RECV_MSGS_RET  = 2451 // 获取朋友圈消息列表
	S2C_USER_FRIEND_MSG_GET_SEND_LIST_RET  = 2452 // 获取可发送朋友圈消息列表
	S2C_USER_FRIEND_MSG_SEND_MSG_RET       = 2461 // 发送朋友圈消息
	S2C_USER_FRIEND_MSG_REPLY_MSG_RET      = 2462 // 回复朋友圈消息
	S2C_USER_FRIEND_MSG_SET_MSG_STATUS_RET = 2463 // 变更朋友圈消息状态

	S2C_USER_MALE_ROLE_MODIFY_NOTE_RET         = 2501 // 男主修改备注
	S2C_USER_MALE_ROLE_CHANGE_RELATED_CARD_RET = 2502 // 男主修改关联卡牌
	S2C_USER_MALE_ROLE_RECOVER_NOTE_RET        = 2503 // 男主备注还原

	S2C_USER_ENGAGEMENT_SUBMIT_RET  = 2551 // 完成约会
	S2C_USER_EXTRA_STORY_SUBMIT_RET = 2552 // 小传解锁

	S2C_USER_TRACK_TASK_CARD_DEPLOY_RET   = 2601 // 跟踪副本上阵卡牌
	S2C_USER_TRACK_TASK_GENERATE_RET      = 2602 // 跟踪副本生成副本
	S2C_USER_TRACK_TASK_DO_RET            = 2611 // 跟踪副本执行
	S2C_USER_TRACK_TASK_LEAVE_RET         = 2612 // 跟踪副本继续
	S2C_USER_TRACK_TASK_DROP_SUBMIT_RET   = 2613 // 跟踪副本掉落奖励领取
	S2C_USER_TRACK_TASK_REWARD_SUBMIT_RET = 2621 // 跟踪副本奖励领取
	S2C_USER_TRACK_TASK_LEAVE_CHECK_RET   = 2622 // 跟踪副本离开校验

	S2C_USER_FRIEND_APPLY_INFOS_RET      = 2701 // 获取申请好友列表
	S2C_USER_FRIEND_APPLY_FRIEND_RET     = 2702 // 申请好友
	S2C_USER_FRIEND_APPLY_MANAGE_RET     = 2703 // 好友申请管理
	S2C_USER_FRIEND_FRIEND_INFOS_RET     = 2751 // 获取好友列表
	S2C_USER_FRIEND_REMOVE_FRIEND_RET    = 2752 // 删除好友
	S2C_USER_FRIEND_RECOMMEND_FRIEND_RET = 2761 // 推荐好友
	S2C_USER_FRIEND_SEND_MAIL_RET        = 2771 // 发送好友邮件
	S2C_USER_FRIEND_BATCH_SEND_MAIL_RET  = 2772 // (批量)发送好友邮件

	S2C_USER_LOOP_TASK_STATUS_RET          = 2801 // 跑环任务状态
	S2C_USER_LOOP_TASK_SUBMIT_RET          = 2802 // 跑环任务提交
	S2C_USER_LOOP_TASK_LOOP_RWD_SUBMIT_RET = 2803 // 跑环任务领取奖励
	S2C_USER_LOOP_TASK_LOOP_RESET_RET      = 2804 // 跑环任务重置
	S2C_USER_LOOP_TASK_LOOP_SWEEP_RET      = 2805 // 扫荡奖励回包

	S2C_USER_RES_DUNGEON_INFO_RET              = 2810 // 获取资源副本面板信息
	S2C_USER_GET_RES_DUNGEON_THEME_REWARD_RET  = 2811 // 领取资源副本奖励
	S2C_USER_RES_DUNGEON_THEME_BATTLE_RET      = 2812 // 资源副本战斗
	S2C_USER_RES_DUNGEON_THEME_RESET_RET       = 2813 // 资源副本重置
	S2C_USER_RES_DUNGEON_UPDATE_DEFENCE        = 2814 // 资源副本上阵守护
	S2C_USER_RES_DUNGEON_THEME_BATTLE_GOBACK   = 2815 // 资源副本战斗重置
	S2C_USER_RES_DUNGEON_SWEEP_RET             = 2816 // 资源副本扫荡
	S2C_USER_GET_RES_DUNGEON_THEME_REWARDS_RET = 2817 // 批量领取资源副本奖励

	S2C_USER_STONE_RANDOM_RET             = 2820 // 随机石头
	S2C_USER_BIRTH_EXAM_SET_RET           = 2821 // 生日问卷设置
	S2C_USER_BIRTH_REWARD_SUBMIT_RET      = 2822 //生日奖励领取
	S2C_USER_BIRTH_EXAM_REWARD_SUBMIT_RET = 2823 //生日问卷奖励领取
	S2C_USER_BIRTH_DISPLAY_REWARD_RET     = 2824 //生日问卷奖励查看

	S2C_USER_RES_DUNGEON_BOX_DISPATCH_RET = 2830 //资源副本 派遣宝箱
	S2C_USER_RES_DUNGEON_BOX_SPEEDUP_RET  = 2831 //资源副本 加速宝箱
	S2C_USER_RES_DUNGEON_BOX_CANCEL_RET   = 2832 //资源副本 取消派遣
	S2C_USER_RES_DUNGEON_BOX_FINISH_RET   = 2833 //资源副本 完成派遣

	S2C_USER_ACTIVITY_WATER_PLANT_SUBMIT_RET           = 4001 // 浇花活动浇水施肥
	S2C_USER_ACTIVITY_WATER_PLANT_REWARD_SUBMIT_RET    = 4002 // 浇花活动领取羁绊
	S2C_USER_ACTIVITY_GAIN_CONSUME_REWARD_SUBMIT_RET   = 4003 // 日常/累计消耗活动
	S2C_USER_ACTIVITY_DEFEND_TASK_BATTLE_RET           = 4004 // 保卫活动 战斗
	S2C_USER_ACTIVITY_DEFEND_TASKLINE_SUBMIT_RET       = 4005 // 保卫活动 提交活动任务主线
	S2C_USER_ACTIVITY_DEFEND_TASK_REWARD_SUBMIT_RET    = 4006 // 保卫活动 活动奖励领取
	S2C_USER_ACTIVITY_ANSWER_EXAM_SUBMIT_RET           = 4007 // 答题协议
	S2C_USER_ACTIVITY_REVIEW_EXAM_SUBMIT_RET           = 4008 // 复习协议
	S2C_USER_ACTIVITY_EVENT_STORY_SUBMIT_RET           = 4009 // 返回事件剧情协议
	S2C_USER_ACTIVITY_EVENT_TUJIAN_SUBMIT_RET          = 4010 // 事件3,图鉴事件返回
	S2C_USER_ACTIVITY_EVENT_TUJIAN_GETREWARD_RET       = 4011 // 事件3 领取奖励返回
	S2C_USER_ACTIVITY_OBTAIN_BUFF_REWARD_RET           = 4012 //  领取Buff奖励返回
	S2C_USER_ACTIVITY_SINGLE_BUY_GOODS_RET             = 4013 // 单人卡池购买商店物品
	S2C_USER_ACTIVITY_RECALL_BIND_CODE_RET             = 4020 // 绑定code
	S2C_USER_ACTIVITY_RECALL_GET_USER_BRIEF_INFO_RET   = 4021 // 获取玩家基本信息
	S2C_USER_ACTIVITY_RECALL_GET_LOGS_RET              = 4022 // 获取召回积分日志
	S2C_USER_ACTIVITY_RECALL_SHARE_RET                 = 4023 // 召回分享
	S2C_USER_ACTIVITY_PRIVILEGE_CARD_REWARD_SUBMIT_RET = 4024 // 特权卡活动奖励领取

	S2C_USER_ACTIVITY_ANSWER_EXAM3_SUBMIT_RET        = 4025 // 答题3协议
	S2C_USER_ACTIVITY_REVIEW_EXAM3_SUBMIT_RET        = 4026 // 答题3复习协议
	S2C_USER_ACTIVITY_ANSWER_EXAM3_REWARD_SUBMIT_RET = 4027 // 答题3活动领取羁绊
	S2C_USER_ACTIVITY_MAIN_QUEST_WISH_RET            = 4028 // 主线卡池许愿
	S2C_USER_ACTIVITY_MAIN_QUEST_EXCHANGE_RET        = 4029 // 主线卡池兑换
	S2C_USER_ACTIVITY_MAIN_QUEST_INFO_RET            = 4030 // 主线卡信息

	S2C_USER_STATUS_CHANGE              = 5201 // 状态变化
	S2C_USER_MONEY_CHANGE               = 5202 // 货币变化
	S2C_USER_HEART_CHANGE               = 5203 // 体力变化
	S2C_USER_EXP_CHANGE                 = 5204 // 等级经验变化
	S2C_USER_GOODS_CHANGE               = 5205 // 道具变化
	S2C_USER_FIRM_PROP_CHANGE           = 5206 // 公司属性变化
	S2C_USER_CARD_CHANGE                = 5208 // 卡牌变化
	S2C_USER_VIP_EXP_CHANGE             = 5209 // VIP经验变化
	S2C_USER_RECHARGE_EXP_CHANGE        = 5210 // 充值经验变化
	S2C_USER_PRIVILEDGE_CARD_CHANGE     = 5211 // 特权卡变化
	S2C_USER_FILM_CHANGE                = 5212 // 胶片变化
	S2C_USER_SUBSCRIPT_EXPIRE_CHNAGE    = 5213 // 玩家订阅时间变化
	S2C_USER_ACTIVITES_SCORE_CHANGE     = 5214 // 玩家活动积分变化
	S2C_USER_MALE_ROLE_FAVOR_EXP_CHANGE = 5215 // 男主好感度变化
	S2C_USER_HEAD_FRAMES_CHANGE         = 5216 // 玩家头像框变化
	S2C_USER_ACTIVITY_BUFF_CHANGE       = 5217 // 活动buff变化
	S2C_USER_ACTIVITY_EXTRASTORY_CHANGE = 5218 // 活动小传变化

	S2C_USER_SHORT_MSG_CHANGE  = 5221 // 短消息变化
	S2C_USER_PHONE_MSG_CHANGE  = 5222 // 电话消息变化
	S2C_USER_PUBLIC_MSG_CHANGE = 5223 // 公众号消息变化
	S2C_USER_FRIEND_MSG_CHANGE = 5224 // 朋友圈消息变化

	S2C_USER_MAIL_STATUS_RET = 5231 // 邮件状态变化

	S2C_USER_ACTIVITY_SWITCH_CHANGE_RET = 5241 // 活动类型信息变更
	S2C_USER_ACTIVITY_STATUS_CHANGE_RET = 5242 // 活动信息变更
	S2C_USER_ANTI_ADDICTION_RET         = 5243 // 防沉迷通知

	S2C_SVR_VERSION_CHANGE = 5301 // 服务器版本变化
)

const (
	S2S_ACCNT_KICK_OUT = 1 // 玩家踢下线

	S2S_DB_GET_MAIL = 1001 // DB获取玩家邮件信息

	S2S_SYS_NONE      = 1100 // SUB_MSG_ID使用的占位符
	S2S_SYS_BROADCAST = 1101 // 系统广播消息

	S2S_SYS_MAIL_SEND  = 1103 // 系统发送邮件
	S2S_SYS_MSG_SEND   = 1104 // 系统发送广播消息
	S2S_SYS_VER_CHANGE = 1105 // 服务器版本变更
	S2S_SYS_MAIL_DEL   = 1106 // 系统删除邮件
	S2S_SYS_MAIL_LIST  = 1107 // 系统邮件列表

	S2S_SYS_ACTIVITY_SWITCH_CHANGE = 1111 // 活动开关状态变更
	S2S_SYS_ACTIVITY_STATUS_CHANGE = 1112 // 活动状态变更

	S2S_SYS_ACCNT_ITEM_MGR    = 1201 // 更改玩家ITEM信息
	S2S_SYS_ACCNT_AUTH_MGR    = 1202 // 更改玩家状态
	S2S_SYS_ACCNT_TASK_UNLOCK = 1203 // 玩家关卡解锁

	S2S_SYS_ACCNT_USER_GUIDE_UNLOCK = 1211 // 玩家新手引导解锁

	S2S_SYS_ACCNT_ARENA_CNT_ADD        = 1221 // 玩家竞技场次数增加
	S2S_SYS_ACCNT_ARENA_WIN_CNT_MODIFY = 1222 // 玩家竞技场连胜次数修改
)
