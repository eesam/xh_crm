angular.module('yunyunApp')
.directive('emailSameAs', ["$q", "$http", function($q, $http) {
    var link = function(scope, element, attrs, ngModel){
        ngModel.$asyncValidators.unique = function(modelValue, viewValue){
            return $http.post('/checkName.txt', {username: viewValue}).then(
                function(response) {
                    for(var i=0;i<response.data.length;i++){
                        if(response.data[i].user==viewValue)
                            return $q.reject("error");
                    }
                    return true;
                }
            );
        };
    };

    return {
        restrict: 'A',
        require: 'ngModel',
        link: link
    };
}]);
