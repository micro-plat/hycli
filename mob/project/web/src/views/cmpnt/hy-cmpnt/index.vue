<template>
    <cmpnt :is="component" v-if="component" />
  </template>
  <script>

  
  export default {
    name: 'cmpnt',
    props: [ 'type'],
    data() {
      return {
        component: null
      }
    },
    computed: {
      loader() {
        return () => import(`@/views/cmpnt/hy-orderbar`)
      }
    },
    mounted() {
      this.loader()
        .then(() => {
          this.component = () => this.loader()
        })
        .catch(() => {
          this.component = () =>
          import(`@/views/cmpnt/hy-userhead`)
        })
    }
  }
  </script>
  
  