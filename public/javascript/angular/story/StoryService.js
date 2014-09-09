var mod = angular.module("story.services", [])

mod.service("Story", ["$rootScope", "$http", function($rootScope, $http){
  var service = {
    stories: {},
    update: updateStory,
    remove: deleteStory
  }

  $http.get("/stories").success(function(data){
    service.stories = data
    $rootScope.$broadcast("stories.update")
  })

  return service

  function updateStory(story, state){
    $http.post("/stories/" + story.id, story).success(function(data){
      if(state.name != data.state){
        console.error("state not updated")
      }
    })
  }

  function deleteStory(id){
    $http.delete("/stories/" + id).success(function(){
      console.log(arguments)
    })
  }
}])
