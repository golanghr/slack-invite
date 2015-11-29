angular.module('golanghr.slackinvite', ['ui.bootstrap', 'ngProgress']);

angular.module('golanghr.slackinvite').controller('IndexCtrl', function ($rootScope, $scope, ngProgressFactory) {
  $scope.progressbar = ngProgressFactory.createInstance();
  $scope.progressbar.setColor('#02bbff');

  $scope.$on('$routeChangeStart', function(next, current) {
    $scope.progressbar.start();
  });

  $scope.$on('$routeChangeSuccess', function () {
		$scope.progressbar.complete();
	});
});
