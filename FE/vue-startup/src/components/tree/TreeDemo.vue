<template>
    <ul id="demo">
        <item class="item" :model="treeData">
        </item>
    </ul>
</template>
<script>
import MyTree from '@/components/tree/MyTree'
export default {
    name: 'TreeDemo',
    data() {
        return {
            treeData: {
                name: 'My Tree',
                children: [
                    { name: 'hello' },
                    { name: 'wat' },
                    {
                        name: 'child folder',
                        children: [
                            {
                                name: 'child folder',
                                children: [
                                    { name: 'hello' },
                                    { name: 'wat' }
                                ]
                            },
                            { name: 'hello' },
                            { name: 'wat' },
                            {
                                name: 'child folder',
                                children: [
                                    { name: 'hello' },
                                    { name: 'wat' }
                                ]
                            }
                        ]
                    }
                ]
            }
        }
    },
    components: {
        'item': MyTree
    }
}
</script>
<style scoped>
body {
    font-family: Menlo, Consolas, monospace;
    color: #444;
}

ul {
    padding-left: 1em;
    line-height: 1.5em;
    list-style-type: dot;
    text-align: left;
}
</style>
