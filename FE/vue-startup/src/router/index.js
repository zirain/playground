import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Markdown from '@/components/markdown/Markdown'
import GitHubCommits from '@/components/GitHubCommits/GitHubCommits'
import GridDemo from '@/components/grid/GridDemo'
import TreeDemo from '@/components/tree/TreeDemo'
import SvgDemo from '@/components/svg/SvgDemo'
import ModalDemo from '@/components/modal/ModalDemo'
import DraggableHeaderDemo from '@/components/draggableheader/DraggableHeaderDemo'
import WrapperComponentDemo from '@/components/wrapper/WrapperComponentDemo'
import RealTime from '@/components/realtime/deepstreamHub'
import Validation from '@/components/Validation/Validation'
import ToDoList from '@/components/todolist/todolist'

Vue.use(Router)

export default new Router({
  routes: [{
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/markdown',
      name: 'markdown',
      component: Markdown
    },
    {
      path: '/GitHubCommits',
      name: 'GitHubCommits',
      component: GitHubCommits
    },
    {
      path: '/Grid',
      name: 'GridDemo',
      component: GridDemo
    },
    {
      path: '/Tree',
      name: 'TreeDemo',
      component: TreeDemo
    },
    {
      path: '/Svg',
      name: 'SvgDemo',
      component: SvgDemo
    },
    {
      path: '/Modal',
      name: 'ModalDemo',
      component: ModalDemo
    },
    {
      path: '/DraggableHeader',
      name: 'DraggableHeaderDemo',
      component: DraggableHeaderDemo
    },
    {
      path: '/WrapperComponent',
      name: 'WrapperComponentDemo',
      component: WrapperComponentDemo
    },
    {
      path: '/RealTime',
      name: 'RealTime',
      component: RealTime
    },
    {
      path: '/Validation',
      name: 'Validation',
      component: Validation
    },
    {
      path: '/todo',
      name: 'ToDoList',
      component: ToDoList
    }
  ]
})
