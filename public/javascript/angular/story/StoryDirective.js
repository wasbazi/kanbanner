var mod = angular.module("story.directives", [])

mod.directive("storyView", function() {
  return {
    restrict: 'A',
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
      $rootScope.$broadcast("story.move", args)
      return false;
    }, false)
  }
}])
