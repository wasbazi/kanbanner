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
