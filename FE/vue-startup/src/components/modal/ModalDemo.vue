<template>
    <div id="app">
        <blockquote>
            <p>使用到的特性：组件，prop 传递，内容插入(content insertion)，过渡（transitions）。</p>
        </blockquote>
        <button id="show-modal" class="btn btn-default btn-primary" @click="showModal = true">Show Modal</button>
        <!-- use the modal component, pass in the prop -->
        <modal v-if="showModal" @close="showModal = false">
            <!--
                                    you can use custom content here to overwrite
                                    default content
                                -->
            <h3 slot="header">custom header</h3>
        </modal>
    </div>
</template>
<script>
import MyModal from '@/components/modal/MyModal'
export default {
    name: 'ModelDemo',
    data() {
        return {
            showModal: false
        }
    },
    components: {
        modal: MyModal
    }
}
</script>
<style>

</style>
