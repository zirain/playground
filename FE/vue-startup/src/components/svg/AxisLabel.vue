<template>
    <text :x="point.x" :y="point.y">
        {{stat.label}}</text>
</template>
<script>
import Vue from 'vue'
export default Vue.extend({
    name: 'axis-label',
    props: {
        stat: Object,
        index: Number,
        total: Number
    },
    computed: {
        point: function () {
            return this.valueToPoint(
                +this.stat.value + 10,
                this.index,
                this.total
            )
        }
    }
})
</script>
