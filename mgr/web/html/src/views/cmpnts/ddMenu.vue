<template>
  <span class="ddmenu">
    <el-dropdown>
      <span class="el-dropdown-link">
        {{ label }}
        <el-icon>
          <arrow-down />
        </el-icon>
      </span>
      <template #dropdown>
        <el-menu :collapse="true" style="width:160px">
          <ddTree @itemChanged="itemChanged" :menuList="menuList"></ddTree>
        </el-menu>
      </template>
    </el-dropdown>
  </span>
</template>

   

<script>
import { nextTick } from 'vue'
import ddTree from './ddTree.vue'

export default {
  components: { ddTree },
  name: "ddmenu",
  props: {
    menuList: {
      type: Array,
      default: [],
    },
    menuType: {
      type: String,
      default: ""
    },
    modelValue: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      label: "请选择..."
    }
  },
  watch: {
    'modelValue' () {
      this.setLabel(this.menuType, this.modelValue)
    }
  },
  mounted() {
        this.setLabel(this.menuType, this.modelValue)
  },
  methods: {
    itemChanged(item) {
      this.setLabel(item.type, item.value)
      this.$emit("update:modelValue", item.value)
      this.$emit("valueChanged", item.value)
    },
    setLabel(tp, id) {
      let label = this.$theia.enum.getTreeNames(tp, id)
      if (!label) {
        return
      }
      this.label = label
    },
  }
}
</script>

<style>
:root {
  --el-menu-item-height: 38px;
  --el-menu-sub-item-height: 38px;
}

.ddmenu {
  margin-top: 8px;
  margin-left: 10px;
}

.ddmenu .el-dropdown-link {
  cursor: pointer;
  color: var(--el-color-primary);
  font-size: 0.7rem;
  margin-top: 9px;
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>