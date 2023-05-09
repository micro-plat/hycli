package system

// insertUserLog 保存用户日志数据
const insertUserLog = `
insert into sso_user_log
(
	user_id,
	title,
	content
)
values
(
	@user_id,
	@title,
	@content
)`
