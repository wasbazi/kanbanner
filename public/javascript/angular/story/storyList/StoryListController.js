var mod = angular.module("storyList", ["story.services", "story.directives", "state.services"])

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
