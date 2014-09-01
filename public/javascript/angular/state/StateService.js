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
