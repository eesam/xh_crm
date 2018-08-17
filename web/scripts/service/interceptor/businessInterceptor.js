'use strict';

angular.module('yunyunApp')
.factory('businessInterceptor', ['$q', function ($q) {
	var business = {
		response: function(response) {
			// 如果请求的是接口,并且存在状态码
			if (!angular.isUndefined(response.data.errorCode)){
				// 登录状态检测
				// if (response.data.errorCode === 6) {
				// 	if (yy_systemRole === "SystemAdmin") {
				// 		window.location.href = './loginSysAdmin.html';
				// 	} else {
				// 		window.location.href = './loginCommunityAdmin.html';
				// 	}
				// }
			}
			return response;
		}
	};
	return business;
}]);