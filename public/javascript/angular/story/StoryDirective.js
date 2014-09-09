var mod = angular.module("story.directives", []) //["story.controller"])

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
