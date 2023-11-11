<template>
    <div>
        <template v-for="(item) in menuList">
            <!-- 有次级菜单的,则展开子选项-->
            <el-sub-menu v-if="item.children && item.children.length > 0" :key="item" :index="item.id">
                <template #title>
                    <span>{{ item.text }}</span>
                </template>
                <!-- 递归,实现无限级展开 -->
                <dd-tree @itemChanged="itemChanged" :menuList="item.children"></dd-tree>
            </el-sub-menu>

            <!-- 没有次级菜单的 -->
            <el-menu-item @click="itemChanged(item)" v-if="!item.children || item.children.length == 0" :key="item"
                :index="item.id">
                <span>{{ item.text }}</span>
            </el-menu-item>
        </template>
    </div>
</template>

<script>
export default {
    name: "ddTree",
    props: {
        menuList: {
            type: Array,
            default: [],
        },

    },
    data() {
        return {
            value: ""
        }
    },
    methods: {
        itemChanged(item) {
            this.$emit("itemChanged", item)
            this.value = item.value
            this.$emit("valueChanged", this.value)
        }
    },
}
</script>

<style></style>