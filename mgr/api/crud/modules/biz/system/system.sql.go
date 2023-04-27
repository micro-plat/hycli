package system

//getSystemInfo 获取系统信息
const getSystemInfo = `
select 
	t.id,
	t.name,
	t.index_url,
	t.logo,
	t.theme,
	t.ident,
	t.login_url,
	t.bg_url
from sso_system_info t
where t.ident = @ident 
and t.status = 0
`
