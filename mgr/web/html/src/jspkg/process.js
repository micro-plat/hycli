export function process() {}

process.prototype.getDays = function (s, e) {
  let start = s ? new Date(s.replace(/-/g, "/")) : null;
  let end = e ? new Date(e.replace(/-/g, "/")) : null;
  let now = new Date();

  //结束时间小于当前时间
  if (end && end < now) {
    return 100;
  }

  //开始时间为空，或大于当前时间
  if (!start || start > now) {
    return 0;
  }

  //结束时间为空，开始时间不为空，默认为50%
  if (!end) {
    return 50;
  }

  //开始 当前时间 结束时间
  let total = this.getWeekHours(start, end);
  let pass = this.getWeekHours(start, now);
  let pv = (pass / total) * 100;
  let persent = pv > 100 ? "100" : pv + "";
  let lst = persent.split(".");
  return lst.length == 0 ? persent : lst[0];
};
process.prototype.getWeekHours = function (beginDate, endDate) {
  //日期差值,即包含周六日、以天为单位的工时，86400000=1000*60*60*24.
  var workDayVal = (endDate - beginDate) / 86400000 + 1;
  //工时的余数
  var remainder = workDayVal % 7;
  //工时向下取整的除数
  var divisor = Math.floor(workDayVal / 7);
  var weekendDay = 2 * divisor;
  //起始日期的星期，星期取值有（1,2,3,4,5,6,0）
  var nextDay = beginDate.getDay();
  //从起始日期的星期开始 遍历remainder天
  for (var tempDay = remainder; tempDay >= 1; tempDay--) {
    //第一天不用加1
    if (tempDay == remainder) {
      nextDay = nextDay + 0;
    } else if (tempDay != remainder) {
      nextDay = nextDay + 1;
    }
    //周日，变更为0
    if (nextDay == 7) {
      nextDay = 0;
    }
    //周六日
    if (nextDay == 0 || nextDay == 6) {
      weekendDay = weekendDay + 1;
    }
  }
  //实际工时（天） = 起止日期差 - 周六日数目。
  let rd = workDayVal - weekendDay;
  return rd < 0 ? 0 : rd * 8; //换算为每天8小时
};
