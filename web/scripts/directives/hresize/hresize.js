'use strict';

/**
 * 高度重设
 */
angular.module('yunyunApp')
.directive('hresize', ['$$utilDebounce', '$window', function($$utilDebounce, $window) {
	return function (scope, element, attr) {
		var w = angular.element($window);
		scope.$watch(function () {
			return {
				'h': window.innerHeight
			};
		}, function (newValue, oldValue) {
			// console.log(newValue, oldValue);
			scope.windowHeight = newValue.h;
			// scope.windowWidth = newValue.w;
			scope.hresize = function () {
				return {
					'height': (newValue.h) + 'px',
					'overflow-y': 'auto'
				};
			};
			// scope.hresizeWithOffset = function (offsetH) {
			// 	scope.$eval(attr.notifier);
			// 	return {
			// 		'height': (newValue.h - offsetH) + 'px'
			// 	};
			// };
		}, true);

		var onResize = $$utilDebounce(function() {
			scope.$apply();
		}, 50);

		w.bind('resize', onResize);

		scope.$on('destroy', function() {
			w.off('resize', onResize);
		});
	};
}]);