import Vue from 'vue'
import Vuex from 'vuex'
import * as actions from './action'
import * as getters from './getter'

Vue.use(Vuex)

const state = {
  notes: [],
  activeNote: {},
  show: 'all'
}

const mutations = {
  ADD_NOTE(state) {
    const newNote = {
      text: 'New note',
      favorite: false
    }
    state.notes.push(newNote)
    state.activeNote = newNote
  },

  EDIT_NOTE(state, e) {
    state.activeNote.text = e.target.value
  },

  DELETE_NOTE(state) {
    state.notes.splice(state.notes.indexOf(state.activeNote), 1)
    state.activeNote = state.notes[0] || {}
  },

  TOGGLE_FAVORITE(state) {
    state.activeNote.favorite = !state.activeNote.favorite
  },

  SET_ACTIVE_NOTE(state, note) {
    state.activeNote = note
  },

  SET_SHOW(state, show) {
    state.show = show
  }
}

export default new Vuex.Store({
  state,
  mutations,
  actions,
  getters
})
