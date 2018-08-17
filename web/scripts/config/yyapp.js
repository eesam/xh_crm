// 孝相伴app.js

'use strict';

angular.module('yunyunApp', [
	'oc.lazyLoad',
	'ui.router',
	'ui.bootstrap',
	'ngAnimate',
	'ngSanitize',
	'ngCookies',
	'jcs-autoValidate',
	'angular-md5'
])
// 配置
.config(['$httpProvider', '$stateProvider', '$urlRouterProvider', '$ocLazyLoadProvider', '$cookiesProvider',
	function($httpProvider, $stateProvider, $urlRouterProvider, $ocLazyLoadProvider, $cookiesProvider) {

	/**
	 * [loadModule 加载模块]
	 * @param  {[type]} $stateProvider
	 * @param  {[type]} routes         模块路由
	 * @param  {[type]} modulepath     模块路径
	 * @param  {[type]} routeprefix    路由前缀
	 * @return {[type]}
	 */
	var loadModule = function($stateProvider, routes, modulepath, routeprefix) {
		for (var i = 0; i < routes.length; i++) {
			var route;
			var url;
			if (routes[i].route !== '') {
				if (routeprefix !== '') {
					route = routeprefix + '.' + routes[i].route;
				} else {
					route = routes[i].route;
				}
				var routeArray = routes[i].route.split('.');
				url = '/' + routeArray[routeArray.length-1];
			} else {
				route = routeprefix;
				url = '/' + routeprefix;
			}

			var templateUrl = routes[i].templateUrl;
			if (typeof(routes[i].templateUrl) != "undefined") {
				templateUrl = modulepath + '/' + routes[i].templateUrl;
			}
			var template = routes[i].template;
			var hasOwnControl = routes[i].hasOwnControl,
				abstract      = routes[i].abstract,
				lazyLoad      = routes[i].lazyLoad,
				title         = routes[i].title;

			if (typeof(hasOwnControl) === "undefined") {
				hasOwnControl = true;
			}
			if (typeof(abstract) === "undefined") {
				abstract = false;
			}
			if (typeof(lazyLoad) === "undefined") {
				lazyLoad = [];
			}
			if (typeof(title) === "undefined") {
				title = "";
			}

			var controller = '';
			if (hasOwnControl) {
				var ctrlPath = templateUrl.split('.')[0] + 'Ctrl.js';
				var tempArray = ctrlPath.split('.')[0].split('/');
				var selfCtrl = {
					name: 'yunyunApp',
					files: [ctrlPath]
				};
				lazyLoad.push(selfCtrl);
				controller = tempArray[tempArray.length-1] + ' as ctrl';
			}

			// console.debug(routes[i].idata);
			$stateProvider.state(route, {
				title		: title,
				templateUrl	: templateUrl,
				template	: template,
				url			: url,
				lazyLoad	: lazyLoad,
				controller	: controller,
				idata		: routes[i].idata,
				params      : routes[i].params,
				resolve		: {
					loadMyFile: function ($ocLazyLoad) {
						return $ocLazyLoad.load(this.lazyLoad); // 该函数参数支持列表
					},
					idata : function($q, $http) {
						// if (typeof(this.idata) === "undefined") {
						// 	return;
						// }
						// var deferred = $q.defer();
						// var promise = $http({
						// 	method: 'POST',
						// 	// url: '/trade/base/getSimRealTimeStatus.do',
						// 	url: '/mgr/system/getSystemRole.do',
						// 	params: {}
						// });
						// promise.then(function (result) {
						// 	yy_systemRole = result.data.responseParams.systemRole;
						// 	yy_username = result.data.responseParams.username;
						// 	deferred.resolve(result.data.responseParams);
						// }, function (result) {
						// 	deferred.reject(result);
						// });

						// return deferred.promise;
						return "idata";
					}
				}
			});
		}
	};


	$httpProvider.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=utf-8';

	$httpProvider.interceptors.push('businessInterceptor');
	// $httpProvider.interceptors.push('retryInterceptor');

	$cookiesProvider.defaults.path = "/";
	// {
	// 	path: ''				// 字符串，cookies只在这个路径及其子路径可用。默认情况下，这个将会是出现在你基础标签上的网址路径。
	// 	// domain: yourDomain,	// 字符串，cookies只在这个域及其子域可用。为了安全问题，如果当前域不是需求域的或者其子域，那么用户代理不会接受cookies。
	// 	// expires: expireDate,	// 字符串，日期。"Wdy, DD Mon YYYY HH:MM:SS GMT"格式的字符串或者一个日期对象表示cookies将在这个确切日期/时间过期。
	// 	// secure: true/false	// boolean，该cookies将只在安全连接中被提供。
	// };

	$ocLazyLoadProvider.config( { debug:false, events:true } );

	$urlRouterProvider.otherwise('/back/color');

	// 加载后台模块
	loadModule($stateProvider, yy_route_back, 'back', 'back');
}])
// 校验配置
.run(['$state', 'Yyservice', '$rootScope', 'bootstrap3ElementModifier', 'validator', 'defaultErrorMessageResolver',
	function ($state, Yyservice, $rootScope, bootstrap3ElementModifier, validator, defaultErrorMessageResolver) {
	//是否开启有效状态和无效状态样式
	validator.setValidElementStyling(true);
	validator.setInvalidElementStyling(true);
	//设置错误提示语言
	defaultErrorMessageResolver.setI18nFileRootPath('bower_components/angular-auto-validate/dist/lang');
	defaultErrorMessageResolver.setCulture('zh-cn');
	//开启图标
	// bootstrap3ElementModifier.enableValidationStateIcons(true);

	defaultErrorMessageResolver.getErrorMessages().then(function (errorMessages) {
		errorMessages.password = '请输入6-16位字符，由数据，字母组成';
		errorMessages.same = '两次密码输入不一致';
		errorMessages.phone = '电话号码不正确';
		errorMessages.unique = '该用户已存在';
	});

}]);