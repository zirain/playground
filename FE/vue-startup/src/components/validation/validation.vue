<template>
    <div id="app">
        <ul is="transition-group">
            <li v-for="user in users" class="user" :key="user['.key']">
                <span>{{user.name}} - {{user.email}}</span>
                <button class="btn btn-sm btn-danger" v-on:click="removeUser(user)">X</button>
            </li>
        </ul>
        <form id="form" v-on:submit.prevent="addUser">
            <input type="text" v-model="newUser.name" placeholder="Username">
            <input type="email" v-model="newUser.email" placeholder="email@email.com">
            <input type="submit" class="btn btn-sm btn-primary" value="Add User">
        </form>
        <ul class="errors">
            <li v-show="!validation.name">Name cannot be empty.</li>
            <li v-show="!validation.email">Please provide a valid email address.</li>
        </ul>
    </div>
</template>

<script>
import firebase from 'firebase'
var emailRE = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/

// Setup Firebase
var config = {
    apiKey: 'AIzaSyAi_yuJciPXLFr_PYPeU3eTvtXf8jbJ8zw',
    authDomain: 'vue-demo-537e6.firebaseapp.com',
    databaseURL: 'https://vue-demo-537e6.firebaseio.com'
}
firebase.initializeApp(config)
var usersRef = firebase.database().ref('users')

export default {
    data() {
        return {
            newUser: {
                name: '',
                email: ''
            }
        }
    },
    // firebase binding
    // https://github.com/vuejs/vuefire
    firebase: {
        users: usersRef
    },
    // computed property for form validation state
    computed: {
        validation: function () {
            return {
                name: !!this.newUser.name.trim(),
                email: emailRE.test(this.newUser.email)
            }
        },
        isValid: function () {
            var validation = this.validation
            return Object.keys(validation).every(function (key) {
                return validation[key]
            })
        }
    },
    // methods
    methods: {
        addUser: function () {
            if (this.isValid) {
                usersRef.push(this.newUser)
                this.newUser.name = ''
                this.newUser.email = ''
            }
        },
        removeUser: function (user) {
            usersRef.child(user['.key']).remove()
        }
    }
}
</script>

<style>
ul {
    padding: 0;
}

.user {

    line-height: 30px;
    padding: 10px;
    border-top: 1px solid #eee;
    overflow: hidden;
    transition: all .25s ease;
}

.user:last-child {
    border-bottom: 1px solid #eee;
}

.v-enter,
.v-leave-active {
    height: 0;
    padding-top: 0;
    padding-bottom: 0;
    border-top-width: 0;
    border-bottom-width: 0;
}

.errors {
    color: #f00;
}
</style>
