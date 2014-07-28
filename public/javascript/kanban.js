var app = angular.module('kanbanApp', [])
  app.controller('KanbanController', ['$scope', '$http', function($scope, $http) {
      // $http.get('/stories').success(function(data){
      //   $scope.stories = data
      // })
  }]);

  app.controller('StoryListCtrl', ['$scope', '$http', function($scope, $http){
    $scope.stories = [ {title: 'learn stuff', body: 'learn a whole bunch of stuff'}]
  }])
