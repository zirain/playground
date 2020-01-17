<template>
    <div class="group pubsub">
        <div class="half left">
            <h2>Publish</h2>
            <button class="half left" id="send-event" @click="handleClick()">Send test-event with</button>
            <input type="text" class="half" id="event-data" v-model="value" />
        </div>
        <div class="half">
            <h2>Subscribe</h2>
            <ul id="events-received">
                <template v-for="event in eventsReceived">
                    <li> {{event}} </li>
                </template>
            </ul>
        </div>
    </div>
</template>
<script>
import Vue from 'vue'
export default Vue.extend({
    name: 'event',
    props: ['ds'],
    data() {
        return {
            eventsReceived: [],
            value: ''
        }
    },
    created() {
        this.event = this.ds.event
        this.event.subscribe('test-event', value => {
            this.eventsReceived.push(value)
        })
    },
    methods: {
        handleClick() {
            this.event.emit('test-event', this.value)
        }
    }
})
</script>

