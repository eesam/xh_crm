angular.module('yunyunApp')
.directive('assertSameAs', [function() {
    return {
        restrict: 'A',
        require: 'ngModel',
        link: function(scope, elm, attrs, ngModel) {
            var validateFn = function (viewValue) {
                var password=scope.$eval(attrs.assertSameAs);
                if(typeof(password)== "undefined") {
                    return viewValue;
                }else if(password==viewValue){
                    ngModel.$setValidity('same', true);
                }else{
                    ngModel.$setValidity('same', false);
                }
                return viewValue;
            };

            scope.$watch(
                function () {
                    return scope.$eval(attrs.assertSameAs);
                }, function () {
                    validateFn(ngModel.$viewValue);
                }
            );

            ngModel.$parsers.push(validateFn);
            ngModel.$formatters.push(validateFn);
        }
    };
}]);
