<template>
    <div id="demo">
        <div class="text-center">
            <!-- Use the component -->
            <svg width="200" height="200">
                <polygraph :stats="stats"></polygraph>
            </svg>
        </div>
        <div class="row">
            <div class="col-lg-9">
                <!-- controls -->
                <div v-for="stat in stats" :key="stat.name">
                    <label>{{stat.label}}</label>
                    <input type="range" v-model="stat.value" min="0" max="100">
                    <span>{{stat.value}}</span>
                    <button @click="remove(stat)" class="remove">X</button>
                </div>
                <form id="add">
                    <input name="newlabel" v-model="newLabel">
                    <button class="btn btn-sm btn-primary" @click="add">Add a Stat</button>
                </form>
            </div>
            <div class="col-lg-3">
                <pre id="raw">{{ stats }}</pre>
            </div>
        </div>
    </div>
</template>
<script>
import PolyGraph from '@/components/svg/PolyGraph'
export default {
    name: 'svgdemo',
    data() {
        return {
            newLabel: '',
            stats: [
                { label: 'A', value: 100 },
                { label: 'B', value: 100 },
                { label: 'C', value: 100 },
                { label: 'D', value: 100 },
                { label: 'E', value: 100 },
                { label: 'F', value: 100 }
            ]
        }
    },
    methods: {
        add: function (e) {
            e.preventDefault()
            if (!this.newLabel) return
            this.stats.push({
                label: this.newLabel,
                value: 100
            })
            this.newLabel = ''
        },
        remove: function (stat) {
            if (this.stats.length > 3) {
                this.stats.splice(this.stats.indexOf(stat), 1)
            } else {
                alert('Can\'t delete more!')
            }
        }
    },
    components: {
        'polygraph': PolyGraph
    }
}
</script>
<style>
body {
    font-family: Helvetica Neue, Arial, sans-serif;
}

polygon {
    fill: #42b983;
    opacity: .75;
}

circle {
    fill: transparent;
    stroke: #999;
}

text {
    font-family: Helvetica Neue, Arial, sans-serif;
    font-size: 10px;
    fill: #666;
}

label {
    display: inline-block;
    margin-left: 10px;
    width: 20px;
}

#raw {
    top: 0;
    left: 300px;
}
</style>
