'use strict';

angular.module('yunyunApp')
.directive('content', [function(){
	return {
		templateUrl: 'scripts/directives/content/content.html',
		restrict: 'E',
		replace: true,
		controller:function(){
			var height = ((window.innerHeight > 0) ?window.innerHeight : this.screen.height) - 1;
			$("#page-wrapper").css("height", (height-50) + "px");

			$(window).bind("load resize", function() {
				height = ((this.window.innerHeight > 0) ? this.window.innerHeight : this.screen.height) - 1;
				if (height < 1) height = 1;
				$("#page-wrapper").css("height", (height-50) + "px");
			});
		}
	};
}]);