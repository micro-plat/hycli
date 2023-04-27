package enum

type LoginCode int

//用户名密码错误
const ErrUserPwdNotExists = 901

//用户已锁定暂时无法登录
const ErrUserLocked = 902

//用户已禁用暂时无法登录
const ErrUserDisabled = 903

//角色已锁定暂时无法登录
const ErrRoleLocked = 904

//系统已锁定暂时无法登录
const ErrSystemLocked = 905

//ErrSystemBusy 系统繁忙无法登录
const ErrSystemBusy = 999

const Normal = 0

const BoolTrue = 0
