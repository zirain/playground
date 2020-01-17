<template>
    <div class="panel panel-primary">
        <div class="panel-heading">
            <h3 class="panel-title">本示例创建了一个可复用组件，可结合外部数据来使用它。</h3>
        </div>
        <div class="panel-body">
            <form id="search">
                Search
                <input name="query" v-model="searchQuery">
            </form>
            <hr />
            <vue-grid :data="gridData" :columns="gridColumns" :filter-key="searchQuery">
            </vue-grid>
        </div>
    </div>
</template>

<script>
import VueGrid from '@/components/grid/VueGrid'

export default {
    name: 'GridDemo',
    data() {
        return {
            searchQuery: '',
            gridColumns: ['name', 'power'],
            gridData: [
                { name: 'Chuck Norris', power: Infinity },
                { name: 'Bruce Lee', power: 9000 },
                { name: 'Jackie Chan', power: 7000 },
                { name: 'Jet Li', power: 8000 }
            ]
        }
    },
    components: {
        'vue-grid': VueGrid
    }
}
</script>

<style scoped>
body {
    font-family: Helvetica Neue, Arial, sans-serif;
    font-size: 14px;
    color: #444;
}
</style>
