'use strict';

angular.module('yunyunApp')
.directive('fileModel', [function(){
    return {
        scope: {
            fileModel: "=",
            fileName: "="
        },
        link: function (scope, element, attributes) {
            element.bind("change", function (changeEvent) {
                scope.$apply(function () {
                    scope.fileModel = changeEvent.target.files[0];
                    console.debug(scope.fileModel);
                    if (!!scope.fileModel && !!scope.fileModel.name) {
                        scope.fileName = scope.fileModel.name;
                    }
                    // or all selected files:
                    // scope.fileread = changeEvent.target.files;
                });
            });
        }
    }
}]);
