var app = angular.module("kanbanApp", ["storyList.controller"])

app.controller("KanbanController", [function() {}]);

var mod = angular.module("state.services", [])

mod.service("State", ["$rootScope", "$http", function($rootScope, $http){
  var service = {
    states: []
  }

  $http.get("/states").success(function(data){
    console.log('data', data)
    service.states = data
    $rootScope.$broadcast("states.update")
  })

  return service
}])

var mod = angular.module("story.directives", [])

mod.directive("storyView", function() {
  return {
    restrict: 'E',
    templateUrl: 'javascript/angular/story/StoryView.html',
    scope: true,
    transclude : false
  };
});

mod.directive("draggable", function(){
  return function(scope, element){
    var el = element[0]
    var story = scope.story

    el.draggable = true

    el.addEventListener("dragstart", function(e){
      e.dataTransfer.effectAllowed = "move"
      e.dataTransfer.setData("Title", story.title)
      e.dataTransfer.setData("State", scope.$parent.state.name)
      this.classList.add("drag")
      return false
    }, false)
  }
})

var mod = angular.module("story.services", [])

mod.service("Story", ["$rootScope", "$http", function($rootScope, $http){
  var service = {
    stories: {},
    update: updateStory,
    remove: deleteStory
  }

  $http.get("/stories").success(function(data){
    service.stories = data
    $rootScope.$broadcast("stories.update")
  })

  return service

  function updateStory(story, state){
    $http.post("/stories/" + story.id, story).success(function(data){
      if(state.name != data.state){
        console.error("state not updated")
      }
    })
  }

  function deleteStory(id){
    $http.delete("/stories/" + id).success(function(){
      console.log(arguments)
    })
  }
}])

var mod = angular.module("storyList.controller", ["storyList.directives", "story.services", "state.services"])

mod.controller("StoryListCtrl", ["$scope", "$rootScope", "Story", "State", function($scope, $rootScope, Story, State){
  $scope.$on("states.update", function(){
    $scope.states = State.states
    console.log('states.update', $scope.states)
  })

  $scope.$on("stories.update", function(){
    $scope.stories = Story.stories
    console.log('stories.update', $scope.stories)
  })

  $scope.$on("story.move", function(event, args){
    console.log('$on fired')
    var idx = null
    var stories = $scope.stories[args.prev]

    stories.forEach(function(item, i){
      if(item.title == args.title) return idx = i
    })

    var story = stories.splice(idx, 1)[0]
    story.state = args.state.id
    $scope.stories[args.state.name].push(story)
    $scope.$apply()

    Story.update(story, args.state)
  })
}])

var mod = angular.module("storyList.directives", ["story.directives"])

mod.directive("storyListView", function() {
  return {
    restrict: 'E',
    templateUrl: 'javascript/angular/storyList/StoryListView.html',
    replace: true,
    scope: true,
    transclude: true
  };
});

mod.directive("droppable", ["$rootScope", function($rootScope) {
  return function(scope, element) {
    // again we need the native object
    var el = element[0]

    el.addEventListener("dragover", function(e){
      e.dataTransfer.dropEffect = "move"
      if(e.preventDefault) e.preventDefault()

      this.classList.add("over")
      return false
    }, false)

    el.addEventListener("dragenter", function() {
      this.classList.add("over")
      return false
    }, false);

    el.addEventListener("dragleave", function() {
      this.classList.remove("over")
      return false
    }, false);

    el.addEventListener("drop", function(e){
      if(e.stopPropagation) e.stopPropagation()

      this.classList.remove("over")

      var title = e.dataTransfer.getData("Title")
      var state = e.dataTransfer.getData("State")
      if(state === scope.state){
        return
      }

      var args = { title: title, prev: state, state: scope.state }
      console.log('broadcasted')
      $rootScope.$broadcast("story.move", args)
      return false;
    }, false)
  }
}])
