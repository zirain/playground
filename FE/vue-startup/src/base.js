exports.install = function (Vue, options) {
  // 1. 添加全局方法或属性
  Vue.valueToPoint = function (value, index, total) {
    var x = 0
    var y = -value * 0.8
    var angle = Math.PI * 2 / total * index
    var cos = Math.cos(angle)
    var sin = Math.sin(angle)
    var tx = x * cos - y * sin + 100
    var ty = x * sin + y * cos + 100
    return {
      x: tx,
      y: ty
    }
  }
  // 2. 添加全局资源
  Vue.directive('my-directive', {
    bind (el, binding, vnode, oldVnode) {
      // 逻辑...
    }
  })
  // 3. 注入组件
  Vue.mixin({
    created: function () {
      // 逻辑...
    }
  })
  // 4. 添加实例方法
  Vue.prototype.valueToPoint = function (value, index, total) {
    var x = 0
    var y = -value * 0.8
    var angle = Math.PI * 2 / total * index
    var cos = Math.cos(angle)
    var sin = Math.sin(angle)
    var tx = x * cos - y * sin + 100
    var ty = x * sin + y * cos + 100
    return {
      x: tx,
      y: ty
    }
  }
}
