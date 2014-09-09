var mod = angular.module("story.controller", ["story.services", "story.directives", "state.services"])

mod.controller("StoryCtrl", ["$scope", "$rootScope", "Story", function($scope, $rootScope, Story){
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
