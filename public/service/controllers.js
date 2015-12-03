angular.module('golanghr.slackinvite').controller('MainCtrl', function ($rootScope, $scope, $q, $http, $interval, $window, TeamInfo, ngProgressFactory) {
  $scope.progressbar = ngProgressFactory.createInstance();
  $scope.progressbar.setColor('#02bbff');

  // Using directive so just wanna make sure it's standardized.
  $rootScope.pageLoaded = $scope.pageLoaded;
  $rootScope.invitationSuccess = false;
  $scope.firstName = "";
  $scope.lastName = "";
  $scope.emailAddress = "";

  $window.addEventListener("DOMContentLoaded", function(event) {
    console.log("DOM fully loaded and parsed");
    $rootScope.pageLoaded = true;
  });

  $scope.requestInvitation = function() {
    $scope.progressbar.start();
    var defer = $q.defer();

    $http({method: "GET", url:"//"+api_host+"/api/slack/invite?first_name="+$scope.firstName+"&last_name="+$scope.lastName+"&email="+$scope.emailAddress, cache: false}).
    then(function(response){
      defer.resolve(response);
      $scope.progressbar.complete();
    }, function(response){
      defer.resolve(response);
      $scope.progressbar.complete();
    });

    return defer.promise
  };

  $scope.submit = function() {
    $scope.requestInvitation().then(function(response){
      if (response.data.Ok && response.data.Ok == true) {
        $rootScope.invitationSuccess = true;
        return
      }

      if (response.data.error && response.data.error.indexOf("First name must be provided") > -1) {

      } else if (response.data.error && response.data.error.indexOf("Last name must be provided") > -1) {

      }
    });
  };


  $scope.away_members = [];
  $scope.active_members = [];
  $scope.admins = [];
  $scope.total_members = 0;

  function propagateTeamInfo() {
    TeamInfo.getData().then(function(data){
      $scope.away_members = data.Away;
      $scope.active_members = data.Active;
      $scope.admins = data.Admins;
      $scope.total_members = data.Total;
    });
  }

  propagateTeamInfo();
  $interval(propagateTeamInfo, 10000);

  $scope.$on('$routeChangeStart', function(next, current) {
    $scope.progressbar.start();
  });

  $scope.$on('$routeChangeSuccess', function () {
		$scope.progressbar.complete();
	});
});
