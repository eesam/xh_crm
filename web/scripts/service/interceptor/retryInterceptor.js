'use strict';

angular.module('yunyunApp')
.factory('retryInterceptor', ['$injector', '$q', '$timeout', function ($injector, $q, $timeout) {
	var retry = {
		responseError: function(response) {
			// 状态异常
			if (response.status !== 200) {
				//发送次数小于3次，重新尝试发送
				if (angular.isUndefined(response.config.times) || response.config.times < 3) {
					var deferred = $q.defer();
					if (angular.isUndefined(response.config.times)) {
						response.config.times = 1;
					} else {
						++response.config.times;
					}
					//return $injector.get('$http')(response.config);
					$timeout(function () {
						deferred.resolve();
					}, 1000);
					return deferred.promise.then(function () {
						return $injector.get('$http')(response.config);
					});
				} else {
					// 根据服务器返回错误码判断
					toastr.error(response.status + " " + response.statusText);
					return $q.reject(response);
				}
			} else {
				return $q.reject(response);
			}
		}
	};
	return retry;
}]);