// 孝相伴服务js

'use strict';

angular.module('yunyunApp')
// 拦截器

// 请求/回应处理器
// .config(['$httpProvider', function($httpProvider) {
// 	$httpProvider.defaults.transformRequest = [
// 		function(request) {
// 			return request;
// 		}
// 	];
// 	$httpProvider.defaults.transformResponse = [
// 		function(response) {
// 			return response;
// 		}
// 	];
// }])
// 后台交互服务
.factory('Yyservice', ['$q', '$filter', '$timeout', '$http', 'md5', function ($q, $filter, $timeout, $http, md5) {
	function getBaseUrl() {
		return "/task/";
	}

	///////////////////////////////////////////////////
	// 接口函数实现

	function iHttp(funcname, url, params) {
		console.debug(funcname + "...");
		var deferred = $q.defer();
		var promise = $http({
			method: 'POST',
			url: url,
			params: params
		});
		promise.then(function (resp) {
			deferred.resolve(resp);
		},function (resp) {
			deferred.reject(resp);
		});

		return deferred.promise;
	}

	// 帐号登出
	function logout() {
		var url = getBaseUrl() + 'doLogout.do';
		var params = {};
		return iHttp("logout", url, params);
	}

	// 获取基本资料
	function getAdminInfo() {
		var url = getBaseUrl() + 'getAdminInfo.do';
		var params = {};
		return iHttp("getAdminInfo", url, params);
	}
	// 基本资料录入
	function editAdminInfo(info) {
		var url = getBaseUrl() + 'editAdminInfo.do';
		var params = {
			'communityAdminId': info.id,
			'mobile': info.mobile,
			'email': info.email,
			'note': info.note
		};
		return iHttp("setBaseinfo", url, params);
	}
	// 修改密码
	function changePassword(password, newPassword, newPasswordAgain) {
		var url = getBaseUrl() + 'changePassword.do';
		var params = {
			'password': md5.createHash(password),
			'newPassword': md5.createHash(newPassword),
			'newPasswordAgain': md5.createHash(newPasswordAgain)
		};
		return iHttp("changePassword", url, params);
	}

	function getAdminCommunitList(offset, size) {
		var url = '/mgr/community/getAdminCommunityList4Web.do';
		var params = {
			'offset': offset,
			'size': size
		};
		return iHttp("device addCommunityMember", url, params);
	}
	// 接口函数实现 end
	///////////////////////////////////////////////////

	return {
		logout				: logout,
		getAdminInfo		: getAdminInfo,
		editAdminInfo		: editAdminInfo,
		changePassword		: changePassword,
		getAdminCommunitList: getAdminCommunitList
	};
	// 接口函数实现 end
	///////////////////////////////////////////////////

}])
.filter('deviceTypeFilter', function () {
	return function (value) {
		if (value === "1") return "手表";
		if (value === "2") return "萤石";
		if (value === "3") return "机器人";
		return "未知";
	};
})
// 1-	男性
// 2-	女性
.filter('sexFilter', function () {
	return function (value) {
		if (value === "1") return "男";
		if (value === "2") return "女";
		return "111";
	};
});
