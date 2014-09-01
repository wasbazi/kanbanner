var mod = angular.module("story.services", [])

mod.service("Story", ["$rootScope", "$http", function($rootScope, $http){
  var service = {
    stories: {},
    update: updateStory
  }

  $http.get("/stories").success(function(data){
    service.stories = data
    $rootScope.$broadcast("stories.update")
  })

  return service

  function updateStory(story, state){
    console.log("story", story.state)
    $http.post("/story/" + story.id, story).success(function(data){
      console.log('response', data)
      if(state.name != data.state){
        console.error("state not updated")
      }
    })
  }
}])
