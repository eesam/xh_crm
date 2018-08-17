'use strict';

angular.module('yunyunApp')
.directive('multiselector', ['$location', function($location) {
	return {
		templateUrl:'scripts/directives/multiselector/multiselector.html',
		restrict: 'E',
		replace: true,
		scope: {
			'leftModel': '=',
			'rightModel': '='
		},
		controller: function($scope) {
			// 全不选
			$scope.selectNone = function() {
				for (var i = 0; i < $scope.rightModel.length; i++) {
					$scope.leftModel.push($scope.rightModel[i]);
				}
				$scope.rightModel = [];
			};

			// 全选
			$scope.selectAll = function() {
				for (var i = 0; i < $scope.leftModel.length; i++) {
					$scope.rightModel.push($scope.leftModel[i]);
				}
				$scope.leftModel = [];
			};

			// 移动到左侧
			$scope.toLeft = function() {
				// 将右边选中的加入到左边
				for (var i = 0; i < $scope.rightSelected.length; i++) {
					$scope.leftModel.push($scope.rightSelected[i]);
				}

				// 将右边选中的删掉
				for (var k = 0; k < $scope.rightSelected.length; k++) {
					for (var j = 0; j < $scope.rightModel.length; j++) {
						if ($scope.rightSelected[k].communityId == $scope.rightModel[j].communityId) {
							$scope.rightModel.splice(j, 1);
							break;
						}
					}
				}

				// 选中清空
				$scope.rightSelected = [];
			};

			// 移动到右侧
			$scope.toRight = function() {
				// 将左边选中的加入到右边
				for (var i = 0; i < $scope.leftSelected.length; i++) {
					$scope.rightModel.push($scope.leftSelected[i]);
				}

				// 将左边选中的删掉
				for (var k = 0; k < $scope.leftSelected.length; k++) {
					for (var j = 0; j < $scope.leftModel.length; j++) {
						if ($scope.leftSelected[k].communityId == $scope.leftModel[j].communityId) {
							$scope.leftModel.splice(j, 1);
							break;
						}
					}
				}

				// 选中清空
				$scope.leftSelected = [];
			};
		}
	};
}]);
