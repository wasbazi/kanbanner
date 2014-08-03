var app = angular.module('kanbanApp', [])
app.controller('KanbanController', ['$scope', '$http', function($scope, $http) {
}]);

app.controller('StoryListCtrl', ['$scope', '$rootScope', 'Story', function($scope, $rootScope, Story){
  // $scope.stories = [ {title: 'learn stuff', body: 'learn a whole bunch of stuff'}]
  $scope.$on('stories.update', function(){
    $scope.stories = Story.stories
  })

  $scope.$on('story.move', function(event, args){
    var idx = null
    var stories = $scope.stories[args.prev]
    console.log(args.key)
    stories.forEach(function(item, i){
      if(item.title == args.title) return idx = i
    })

    if(idx === null) {
      console.log('whoops')
    }

    var story = stories.splice(idx, 1)[0]
    story.state = args.state
    $scope.stories[args.state].push(story)
    $scope.$apply()
  })
}])

app.service('Story', ['$rootScope', '$http', function($rootScope, $http){
  var service = { stories: {} }

  $http.get('/stories').success(function(data){
    service.stories = data
    $rootScope.$broadcast('stories.update')
  })

  return service
}])

app.directive('draggable', function(){
  return function(scope, element){
    var el = element[0]
    var story = scope.story

    el.draggable = true

    el.addEventListener('dragstart', function(e){
      console.log(scope)
      e.dataTransfer.effectAllowed = 'move'
      e.dataTransfer.setData('Title', story.Title)
      e.dataTransfer.setData('State', scope.$parent.state)
      this.classList.add('drag')
      return false
    }, false)
  }
})

app.directive('droppable', ['$rootScope', function($rootScope) {
  return function(scope, element) {
    // again we need the native object
    var el = element[0]

    el.addEventListener('dragover', function(e){
      e.dataTransfer.dropEffect = 'move'
      if(e.preventDefault) e.preventDefault()

      this.classList.add('over')
      return false
    }, false)

    el.addEventListener('dragenter', function() {
      this.classList.add('over')
      return false
    }, false);

    el.addEventListener('dragleave', function() {
      this.classList.remove('over')
      return false
    }, false);

    el.addEventListener('drop', function(e){
      if(e.stopPropagation) e.stopPropagation()

      this.classList.remove('over')

      var title = e.dataTransfer.getData('Title')
      var state = e.dataTransfer.getData('State')
      if(state === scope.state){
        return
      }

      var args = { title: title, prev: state, state: scope.state }
      $rootScope.$broadcast('story.move', args)
      return false;
    }, false)
  }
}])
