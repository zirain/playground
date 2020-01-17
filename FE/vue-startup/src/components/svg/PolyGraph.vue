<template>
    <g>
        <polygon :points="points"></polygon>
        <circle cx="100" cy="100" r="80"></circle>
        <axis-label v-for="(stat, index) in stats" :key="stat.name" :stat="stat" :index="index" :total="stats.length">
        </axis-label>
    </g>
</template>
<script>
import Vue from 'vue'
import AxisLabel from '@/components/svg/AxisLabel'
export default Vue.extend({
    name: 'ploygraph',
    props: ['stats'],
    computed: {
        // a computed property for the polygon's points
        points: function () {
            var total = this.stats.length
            var parent = this
            return this.stats.map(function (stat, i) {
                var point = parent.valueToPoint(stat.value, i, total)
                return point.x + ',' + point.y
            }).join(' ')
        }
    },
    components: {
        // a sub component for the labels
        'axis-label': AxisLabel
    }
})
</script>
