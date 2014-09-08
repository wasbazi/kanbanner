var mod = angular.module("storyList.directives", [])

mod.directive("storyListView", function() {
  return {
    restrict: 'A',
    templateUrl: 'javascript/angular/storyList/StoryListView.html',
    replace: true,
    scope: true,
    transclude: true
  };
});
