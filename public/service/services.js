angular.module('golanghr.slackinvite').factory("TeamInfo", function($rootScope, $http, $q) {
  return {
    getData: function() {
      var defer = $q.defer();

      $http.get("//"+api_host+"/api/slack/stats", {cache: false})
      .success(function(data){
        defer.resolve(data);
      });

      return defer.promise
    }
  }
});
