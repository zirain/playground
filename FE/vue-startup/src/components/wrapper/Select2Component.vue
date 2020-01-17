<template>
    <select>
        <slot></slot>
    </select>
</template>
<script>
import Vue from 'vue'
import $ from 'jquery'
import Select2 from 'select2'
import 'select2/dist/css/select2.css'
import 'select2-bootstrap-theme/dist/select2-bootstrap.css'

export default Vue.extend({
    name: 'select2',
    props: ['options', 'value'],
    mounted: function () {
        var vm = this
        $(this.$el)
            // init select2
            .select2({
                data: this.options,
                theme: 'bootstrap'
            })
            .val(this.value)
            .trigger('change')
            // emit event on change.
            .on('change', function () {
                vm.$emit('input', this.value)
            })
    },
    watch: {
        value: function (value) {
            // update value
            $(this.$el).val(value).trigger('change')
        },
        options: function (options) {
            // update options
            $(this.$el).select2({ data: options })
        }
    },
    destroyed: function () {
        $(this.$el).off().select2('destroy')
    }
})
</script>
