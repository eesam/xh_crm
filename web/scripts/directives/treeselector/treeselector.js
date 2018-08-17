'use strict';


angular.module('yunyunApp')
.directive('treeselector', ["$compile", function($compile) {
    return {
        restrict: 'AE',
        templateUrl:'scripts/directives/treeselector/treeselector.html',
        scope: {
            'data': '=',
            'modal':'='
        },
        controller: function($scope) {
             $scope.isleave=true;
             $scope.isShow=false;
             $scope.change=function(){
                $scope.isleave=!$scope.isleave;
             };

             $scope.showSelect=function($event){
                $($event.target).next().css("width",$($event.target).css("width"));
                $scope.isShow=true;
             };

             $scope.hideSelect=function($event){
                    if($scope.isleave) {
                    $scope.isShow = false;
                }else{
                    $($event.target).focus();
                }
             };

             $scope.sureSelected=function(node){
                 $scope.modal.pname = node.name;
                 $scope.modal.pid = node.id;
                 $scope.isShow = false;
             };
        }
    };
}]);
