<template>
    <div id="note-editor">
        <textarea :value="activeNoteText" @input="editNote" class="form-control">
        </textarea>
    </div>
</template>

<script>

import { mapState, mapGetters, mapActions } from 'vuex'

export default {
    computed: {
        ...mapGetters(['activeNoteText'])
    },
    methods: mapActions(['editNote'])
}
</script>
