package menu

const getMenuByUserId = `
select m.id,m.name,m.parent,m.icon,m.url path,m.is_open
from sso_system_menu m
inner join sso_system_info s on s.id = m.sys_id and s.ident=@ident
inner join sso_user_role r on r.sys_id=m.sys_id and r.user_id=@user_id
inner join sso_role_menu u on u.role_id = r.role_id and u.sys_id=m.sys_id and FIND_IN_SET(m.id,u.menu_id)
where s.status = 0
and m.menu_type <= 2
order by m.parent asc,m.sortrank asc
`
