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
}])
