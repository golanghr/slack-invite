angular.module('golanghr.slackinvite').factory("TeamInfo", function($rootScope, $http, $q) {
  return {
    getData: function() {
      var defer = $q.defer();

      $http.get("//"+window.location.hostname+":8500/api/slack/stats", {cache: false})
      .success(function(data){
        defer.resolve(data);
      });

      return defer.promise
    }
  }
});
