package member

//sqlUserLogin 用户登录
const sqlUserLogin = `
select t.user_id,t.user_name,t.name,t.password,
t.mobile,t.status u_status,ro.status r_status,ri.status ri_status,
s.status s_status,s.index_url,s.login_url,s.ident,
s.name sys_name,s.theme,ri.name role_name,
if(TIMESTAMPDIFF(DAY, t.last_failed_time, now())<1 and t.curr_failed_cnt>=t.max_failed_cnt,0,1) is_lock
from sso_user_account t
left join sso_system_info s on s.ident=@ident
left join sso_user_role ro on t.user_id=ro.user_id and ro.sys_id=s.id
left join sso_role_info ri on ri.role_id=ro.role_id
where t.user_name = @user_name
and t.status =0
and ri.status = 0
and s.status = 0`

const sqlUpdateLoginBySuccess = `
update sso_user_account t
set t.last_login_time = now(),
t.curr_failed_cnt = 0
where t.user_id=@user_id
`
const sqlUpdateLoginByFailed = `
update sso_user_account t
set t.last_login_time = now(),
t.curr_failed_cnt = if (TIMESTAMPDIFF(DAY, t.last_failed_time, now()) > 1,1,t.curr_failed_cnt + 1),
t.last_failed_time = now()
where t.user_name = @user_name
`
