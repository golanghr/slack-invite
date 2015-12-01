angular.module('golanghr.slackinvite').controller('MainCtrl', function ($rootScope, $scope, TeamInfo, ngProgressFactory) {
  $scope.progressbar = ngProgressFactory.createInstance();
  $scope.progressbar.setColor('#02bbff');

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
  setInterval(propagateTeamInfo, 10000);


  $scope.$on('$routeChangeStart', function(next, current) {
    $scope.progressbar.start();
  });

  $scope.$on('$routeChangeSuccess', function () {
		$scope.progressbar.complete();
	});
});
