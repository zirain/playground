<template>
    <li>
        <div :class="{bold: isFolder}" @click="toggle" @dblclick="changeType">
            {{model.name}}
            <span v-if="isFolder">[{{open ? '-' : '+'}}]</span>
        </div>
        <ul v-show="open" v-if="isFolder">
            <item class="item" v-for="model in model.children" :key="model.name" :model="model">
            </item>
            <li class="add" @click="addChild">+</li>
        </ul>
    </li>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({
    name: 'item',
    props: {
        model: Object
    },
    data: function () {
        return {
            open: false
        }
    },
    computed: {
        isFolder: function () {
            return this.model.children &&
                this.model.children.length
        }
    },
    methods: {
        toggle: function () {
            if (this.isFolder) {
                this.open = !this.open
            }
        },
        changeType: function () {
            if (!this.isFolder) {
                Vue.set(this.model, 'children', [])
                this.addChild()
                this.open = true
            }
        },
        addChild: function () {
            this.model.children.push({
                name: 'new stuff'
            })
        }
    }
})
</script>
<style scoped>
.item {
    cursor: pointer;
}

.bold {
    font-weight: bold;
}
</style>
