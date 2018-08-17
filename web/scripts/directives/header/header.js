'use strict';

angular.module('yunyunApp')
.directive('header', [function(){
	return {
		templateUrl:'scripts/directives/header/header.html',
		restrict: 'E',
		replace: true,
		controller: ['$scope', 'Yyservice', '$state', '$cookies', function($scope, Yyservice, $state, $cookies) {
			$scope.username = yy_username;

			// $scope.myconsole = {
			// 	id: 0,
			// 	fullname: "我的控制台",
			// 	shortname: "我的控制台",
			// 	pid: 0
			// };

			// Yyservice.getRecentMessage().then(function (result) {
			// 	$scope.msgs = result.data.msgs;
			// });

			// Yyservice.getOrgList(0, 100).then(function (result) {
			// 	$scope.consoles = [];
			// 	for (var i = 0; i < result.data.orgs.length; i++) {
			// 		if (result.data.orgs[i].pid === 0) {
			// 			$scope.consoles.push(result.data.orgs[i]);
			// 		}
			// 	}
			// 	if ($scope.consoles.length === 0) {
			// 		$scope.consoles.push($scope.myconsole);
			// 	}
			// 	$scope.consoleselected = $scope.consoles[0];
			// });

			// $scope.onSelectOrg = function(console) {
			// 	$scope.consoleselected = console;
			// };

			$scope.loginOut = function() {
				Yyservice.logout().then(function (result) {
					if (result.status == 200) {
						if (result.data.errorCode === 0) {
							if (result.data.responseParams.systemRole === "SystemAdmin") {
								window.location.href = './loginSysAdmin.html';
							} else if (result.data.responseParams.systemRole === "CommunityAdmin") {
								window.location.href = './loginCommunityAdmin.html';
							}
						} else {
							toastr.error("退出失败：" + result.data.errorMsg);
						}
					} else {
						toastr.error("退出失败：" + result.status + " " + result.statusText);
					}
				});
			};
		}]
	};
}]);