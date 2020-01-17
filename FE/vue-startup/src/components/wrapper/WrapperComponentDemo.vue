<template>
    <div id="el">
        <div>
            <p>Selected: {{ selected }}</p>
            <select2 :options="options" v-model="selected">
                <option disabled value="0">Select one</option>
            </select2>
        </div>
    </div>
</template>
<script>
import select2 from '@/components/wrapper/Select2Component'
export default {
    name: 'wrapper-component',
    data() {
        return {
            selected: 2,
            options: [
                { id: 1, text: 'Hello' },
                { id: 2, text: 'World' }
            ]
        }
    },
    components: {
        'select2': select2
    }
}
</script>
<style>
select {
    min-width: 300px;
}
</style>
