angular.module('golanghr.slackinvite').directive("pagePreloader", function($rootScope) {
  return {
    restrict: "E",
    replace:  true,
    template: '<div class="animate-hide" id="page-preloader" ng-hide="pageLoaded"></div>'
  }
});
