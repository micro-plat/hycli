{-{- $table := .}-}
{-{- $ok:= $table.Optrs.All.Find "note"}-}
{-{- if $ok}-}
{-{- $note:= $table.Optrs.All.GetOne "note"}-}
    {-{- $queryMethod:=$note.GetParam "query" ""}-}
    {-{- $saveMethod:=$note.GetParam "save" ""}-}
    {-{- $updateMethod:=$note.GetParam "update" ""}-}
    {-{- $delMethod:=$note.GetParam "del" ""}-}
    {-{- $sortMethod:=$note.GetParam "sort" ""}-}
    {-{- if ne "" $queryMethod}-}
    query_note_{-{$note.UNQ}-}(form) {
        let that = this
        this.$theia.http.post("/{-{$table.Name.MainPath|lower}-}/{-{$queryMethod}-}",form).then(res=>{
        that.note_dataList = res.items
        }).catch(res=>{
            // let code = ((res||{}).response||{}).status||0
            // let msg = `查询数据失败(${code})`
            // that.$notify.error({title: '失败',message: msg,duration:5000})
        });
    },
    {-{- end}-}
    {-{- if ne "" $saveMethod}-}
    save_note_{-{$note.UNQ}-}(form) {
        let that = this
        this.$theia.http.post("/{-{$table.Name.MainPath|lower}-}/{-{$saveMethod}-}",form).then(res=>{
            that.query_note_{-{$note.UNQ}-}()
        }).catch(res=>{
            // let code = ((res||{}).response||{}).status||0
            // let msg = `创建失败(${code})`
            // that.$notify.error({title: '失败',message: msg,duration:5000})
        });
    },
    {-{- end}-}
    {-{- if ne "" $updateMethod}-}
    update_note_{-{$note.UNQ}-}(form) {
        let that = this
        this.$theia.http.post("/{-{$table.Name.MainPath|lower}-}/{-{$updateMethod}-}",form).then(res=>{
            that.query_note_{-{$note.UNQ}-}()
        }).catch(res=>{
            // let code = ((res||{}).response||{}).status||0
            // let msg = `更新内容失败(${code})`
            // that.$notify.error({title: '失败',message: msg,duration:5000})
        });
    },
    {-{- end}-}
    {-{- if ne "" $delMethod}-}
    del_note_{-{$note.UNQ}-}(form) {
        let that = this
        this.$theia.http.post("/{-{$table.Name.MainPath|lower}-}/{-{$delMethod}-}",form).then(res=>{
            that.query_note_{-{$note.UNQ}-}()
        }).catch(res=>{
            // let code = ((res||{}).response||{}).status||0
            // let msg = `删除失败(${code})`
            // that.$notify.error({title: '失败',message: msg,duration:5000})
        });
    },
    {-{- end}-}
    {-{- if ne "" $sortMethod}-}
    changeSort_note_{-{$note.UNQ}-}(form) {
        let that = this
        this.$theia.http.post("/{-{$table.Name.MainPath|lower}-}/{-{$sortMethod}-}",form).then(res=>{
            that.query_note_{-{$note.UNQ}-}()
        }).catch(res=>{
            // let code = ((res||{}).response||{}).status||0
            // let msg = `更新排序失败(${code})`
            // that.$notify.error({title: '失败',message: msg,duration:5000})
        });
    },
    {-{- end}-}
{-{- end}-}