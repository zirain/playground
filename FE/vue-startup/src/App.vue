<template>
  <div id="app">
    <nav class="navbar navbar-inverse" role="navigation">
      <div class="container">
        <!-- Brand and toggle get grouped for better mobile display -->
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="#">{{ title }}</a>
        </div>
        <!-- Collect the nav links, forms, and other content for toggling -->
        <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
          <ul class="nav navbar-nav">
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Basic
                <span class="caret"></span>
              </a>
              <ul class="dropdown-menu">
                <li>
                  <a href="#/markdown">Markdown</a>
                </li>
                <li>
                  <a href="#/githubcommits">GitHubCommits</a>
                </li>
              </ul>
            </li>
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Components
                <span class="caret"></span>
              </a>
              <ul class="dropdown-menu">
                <li>
                  <a href="#/grid">Grid</a>
                </li>
                <li>
                  <a href="#/tree">Tree</a>
                </li>
                <li>
                  <a href="#/svg">Svg</a>
                </li>
                <li>
                  <a href="#/modal">Modal</a>
                </li>
                <li>
                  <a href="#/DraggableHeader">Draggable Header</a>
                </li>
                <li>
                  <a href="#/WrapperComponent">Wrapper Component</a>
                </li>
              </ul>
            </li>
            <li>
              <a href="#/realtime">RealTime</a>
            </li>
            <li>
              <a href="#/validation">Validation</a>
            </li>
            <li>
              <a href="#/todo">ToDoMvc</a>
            </li>
          </ul>
        </div>
        <!-- /.navbar-collapse -->
      </div>
      <!-- /.container-fluid -->
    </nav>
    <div class="container">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: 'app',
  data() {
    return {
      img: require('./assets/logo.png'),
      title: 'Vue-Startup'
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
