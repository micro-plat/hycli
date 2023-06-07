<template>
    <div>
        <el-image v-if="type.toLowerCase()=='img'"
        :style="{ margin: margin, height: height, width: width }"
        :src="url"
        :zoom-rate="1.2"
        :preview-src-list="[url]"
        :initial-index="4"
        fit="cover"
      />
    
    </div>

</template>

<script>
export default {
    props: {
        type: {
            type: String,
            default:"img"
        },
        url: {
            type: String,
        },
        height: {
            type: String,
            default: "auto"
        },
        width: {
            type: String,
            default: "100%"
        },
        rows: {
            type: String,
            default: "2"
        },
        margin: {
            type: String,
            default: "8px 0px 16px 0px"
        },
    },
    data() {

        return {
           
        }
    },
    methods: {
        show(queryForm = {}) {
           
        }
    },

}
</script>

<style></style>