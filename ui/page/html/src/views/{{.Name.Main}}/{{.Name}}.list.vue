<template>
  {{ $table := .|fltrUITable}}
  {{ $qrow := $table.QRows}}
  {{ $lstRow := $table.LRows}}
  <div>
    <div class="query">
      <el-form :model="form" inline size="small">
        <!-- 日期范围控件 -->
        {{ range $i,$c:= $qrow }}
        {{- if eq "daterange" $c.RType.Type  -}}
        <el-form-item>
          <el-date-picker
            style="width: 200px"
            v-model="form.{{$c.Name}}"
            type="daterange"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        {{- end -}}{{ end }}

        <!-- 日期控件 -->
        {{ range $i,$c:= $qrow }}
        {{- if eq "date" $c.RType.Type  -}}
        <el-form-item>
          <el-date-picker
            style="width: 120px"
            v-model="form.{{$c.Name}}"
            clearable
            type="date"
            placeholder="{{.Desc}}"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        {{- end -}}{{ end }}

        <!-- 下拉控件 -->
        {{ range $i,$c:= $qrow }}
        {{- if eq "select" $c.RType.Type  -}}
        <el-form-item>
          <el-select
            v-model="form.{{.Name}}"
            style="width: 120px"
            clearable
            placeholder="{{.Desc}}"
          >
            <el-option
              v-for="item in {{.Name}}List"
              :key="item.value"
              :label="item.name"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        {{- end -}}{{ end }}

        <!-- 下拉控件多选 -->
        {{ range $i,$c:= $qrow }}
        {{- if eq "multiselect" $c.RType.Type  -}}
        <el-form-item>
          <el-select
            v-model="form.{{.Name}}"
            clearable
            multiple
            collapse-tags
            placeholder="{{.Desc}}"
          >
            <el-option
              v-for="item in {{.Name}}List"
              :key="item.value"
              :label="item.name"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        {{- end -}}{{ end }}

        <!--  数字控件 -->
        {{ range $i,$c:= $qrow }}
        {{- if eq "number" $c.RType.Type  -}}
        <el-form-item label="{{$c.Desc}}">
          <el-input-number v-model="form.{{ $c.Name }}"
          {{ if ne "" $c.RType.Min}} :min="{{ $c.RType.Min }}" {{ end }}
          {{ if ne "" $c.RType.Max}} :max="{{ $c.RType.Max }}" {{ end }}
          />
        </el-form-item>
        {{- end -}} {{ end }}
        <!-- 输入控件 -->
        {{ range $i,$c:= $qrow }}
        {{- if eq "input" $c.RType.Type  -}}
        <el-form-item>
          <el-input
            clearable
            style="width: 120px"
            v-model="form.{{$c.Name}}"
            maxlength="{{$c.RType.Len}}"
            placeholder="{{$c.Desc}}"
          />
        </el-form-item>
        {{- end -}} {{ end }}
        <el-form-item>
          <el-button type="primary" :icon="Search" @click="onQuery"
            >查询</el-button
          >
          <el-button type="success" :icon="Search" @click="showAdd"
            >添加</el-button
          >
        </el-form-item>
      </el-form>
    </div>
    <div class="list">
      <el-table
        :data="dataList"
        v-loading="conf.loading"
        stripe
        style="width: 100%"
        size="small"
        empty-text="无数据"
        height="540px"
      >
        {{- range $i,$c:=$lstRow -}}
        {{- if eq "switch" $c.RType.Type}}
        <el-table-column align="center" label="{{$c.Desc}}">
          <template #default="scope">
            <el-switch disabled v-model="scope.row.{{$c.Name}}_switch" />
          </template>
        </el-table-column>
        {{- else if eq "progress" $c.RType.Type}}
        <el-table-column align="center" label="{{$c.Desc}}">
          <template #default="scope">
            <el-progress
              :percentage="scope.row.{{$c.Name}}"
              :color="conf.progressColor"
            />
          </template>
        </el-table-column>
        {{- else if eq "link" $c.RType.Type}}
        <el-table-column align="center" label="{{$c.Desc}}">
          <template #default="scope">
            <el-link
              type="primary"
              :href="scope.row.{{$c.Name}}"
              :title="scope.row.{{$c.Name}}"
            >查看</el-link>
          </template>
        </el-table-column>
        {{- else if eq "select" $c.RType.Type}}
        {{- if eq true $c.Ext.ColorType -}}
        <el-table-column align="center" label="{{$c.Desc}}">
          <template #default="scope">
            <span
              :class="scope.row.{{$c.Name}}==0?'text-success':'text-warning'"
              v-text="scope.row.{{$c.Name}}_label"
            ></span>
          </template>
        </el-table-column>
        {{- else}}
        <el-table-column
          prop="{{$c.Name}}_label"
          align="center"
          label="{{$c.Desc}}"
        />
        {{- end -}}
        {{- else}}
        <el-table-column
          prop="{{$c.Name}}"
          align="center"
          label="{{$c.Desc}}"
        />
        {{- end -}}
        {{- end }}
      </el-table>
      <el-pagination
        :currentPage="form.pi"
        :page-size="form.ps"
        style="position: absolute; right: 30px"
        :page-sizes="[10, 20, 50, 100]"
        :background="false"
        small
        layout="total,sizes,prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
    <CAdd ref="cadd" @onsaved="onQuery"></CAdd>
  </div>
</template>


<script>
import CAdd from "./{{.Name}}.add"
export default {
   components: {
		CAdd,
  },
  data() {
    return {
      conf:{
        loading:false,
        progressColor: this.$theia.env.conf.progress||[]
      },
      form: {
        pi: 1,
        ps: 10,
        {{- range $i,$c := $qrow}}
        {{$c.Name}}:"",
        {{- end}}
        },
        {{- range $i,$c := $qrow -}}
        {{- if or (eq "select" $c.RType.Type) (eq "multiselect" $c.RType.Type)}}
        {{.Name}}List:this.$theia.enum.get("{{$c.Ext.SLType}}"),
        {{- end}}
        {{- end}}
        dataList:[],
        total:0,
    };
  },
    created() {
      this.queryData()
    },
  methods:{
    queryData(){

      //构建查询参数
      let that = this
       {{ range $i,$c:= $qrow }}
        {{- if eq "daterange" $c.RType.Type  -}}
     this.form.start_{{$c.Name}} = null
      this.form.end_{{$c.Name}} = null
      if(this.form.{{- $c.Name -}} && this.form.{{$c.Name}}.length>1){
        this.form.start_{{$c.Name}}=this.form.{{$c.Name}}[0]
        this.form.end_{{$c.Name}}=this.form.{{$c.Name}}[1]
      }
        {{end -}}
        {{- end -}}
      //发送查询请求
      this.conf.loading = true
      this.$theia.http.get("/{{.Name.CName|lower}}/query",this.form).then(res=>{
        that.conf.loading = false
        that.dataList = res.items
        that.total = res.count
        that.dataList.forEach(item => {
        {{- range $i,$c := $qrow -}}
        {{- if eq "select" $c.RType.Type}}
          item.{{$c.Name}}_label = that.$theia.enum.getName("{{$c.Ext.SLType}}",item.{{$c.Name}})
        {{- else if eq "switch" $c.RType.Type}}
          item.{{$c.Name}}_switch = item.{{$c.Name}} == 0
        {{- end}}
        {{- end}}
        });
      })
    },
    handleSizeChange(ps){
      this.form.ps = ps
      this.queryData()
    },
    handleCurrentChange(pi){
      this.form.pi = pi
      this.queryData()
    },
    onQuery(){
      this.form.pi = 1
      this.queryData()
    },
    showAdd(){
       this.$refs.cadd.show()
    }
  },
};
</script>



<style>
.el-form-item {
  margin-right: 10px !important;
}
</style>