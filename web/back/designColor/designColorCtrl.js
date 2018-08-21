'use strict';

angular.module('yunyunApp')
// 视图控制器
.controller('designColorCtrl', ['$scope', '$modal', '$log', 'designColorService', '$stateParams', '$window', 'FileUploader',
	function($scope, $modal, $log, designColorService, $stateParams, $window, FileUploader) {
	// console.log("designId=", $stateParams.designId);
	if ($stateParams.designId == null) {
		$stateParams.designId = $window.localStorage["designId"];
		console.log("1 designId=", $stateParams.designId);
	} else {
		$window.localStorage["designId"] = $stateParams.designId;
		console.log("2 designId=", $stateParams.designId);
	}

	var ctrl = this;

	ctrl.displayed = [];
	ctrl.start = 0;
	ctrl.number = 0;
	ctrl.initList = function initList(tableState) {
		var pagination = tableState.pagination;

		// 同一页不刷新
		if ( (ctrl.start == pagination.start) && (ctrl.number == pagination.number) ) {
			return;
		}

		ctrl.start = pagination.start || 0;
		ctrl.number = pagination.number || 10;

		ctrl.isLoading = true;
		designColorService.getList(ctrl.start, ctrl.number, $stateParams.designId).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					ctrl.displayed = jsonObj.designColorInfos;
					tableState.pagination.numberOfPages = Math.ceil(jsonObj.totalCount/10);
				} else {
					toastr.error("获取列表失败：" + result.data.errorMsg);
				}
			} else {
				toastr.error("获取列表失败：" + result.status + " " + result.statusText);
			}

			ctrl.isLoading = false;
		});
	};

	ctrl.frushList = function frushList() {
		ctrl.isLoading = true;
		designColorService.getList(ctrl.start, ctrl.number, $stateParams.designId).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					ctrl.displayed = jsonObj.designColorInfos;
					// tableState.pagination.numberOfPages = Math.ceil(jsonObj.totalCount/10);
				} else {
					toastr.error("刷新列表失败：" + result.data.errorMsg);
				}
			} else {
				toastr.error("刷新列表失败：" + result.status + " " + result.statusText);
			}

			ctrl.isLoading = false;
		});
	};

	ctrl.isEmpty = function isEmpty() {
		if ( (false == ctrl.isLoading) && (ctrl.displayed.length == 0) ) {
			return true;
		}
		return false;
	}

	ctrl.openAdd = function(size) {
		var selectrow = {
			designId: $stateParams.designId,
			designName: $stateParams.designName
		};
		var modalInstance = $modal.open({
			templateUrl : 'back/designColor/designColorAdd.html',
			controller : 'designColorModal',
			size : size,
			backdrop : 'static',
			resolve : {
				selectrow : function() { return selectrow; },
				modaltype : function() { return 0; }
			}
		});

		modalInstance.result.then(function(selectedItem) {
			$scope.selected = selectedItem;
			ctrl.frushList();
		}, function() {
			$log.info('Modal dismissed at: ' + new Date());
		});
	};

	ctrl.openEdit = function(size, selectrow) {
		var modalInstance = $modal.open({
			templateUrl : 'back/designColor/designColorEdit.html',
			controller : 'designColorModal',
			size : size,
			backdrop : 'static',
			resolve : {
				selectrow : function() { return selectrow; },
				modaltype : function() { return 1; }
			}
		});

		modalInstance.result.then(function(selectedItem) {
			$scope.selected = selectedItem;
			ctrl.frushList();
		}, function() {
			$log.info('Modal dismissed at: ' + new Date());
		});
	};

	ctrl.openDelete = function(size, selectrow) {
		var modalInstance = $modal.open({
			templateUrl : 'back/designColor/designColorDelete.html',
			controller : 'designColorModal',
			size : size,
			backdrop : 'static',
			resolve : {
				selectrow : function() { return selectrow; },
				modaltype : function() { return 2; }
			}
		});

		modalInstance.result.then(function(selectedItem) {
			$scope.selected = selectedItem;
			ctrl.frushList();
		}, function() {
			$log.info('Modal dismissed at: ' + new Date());
		});
	};

	ctrl.openLgPic = function(size, selectrow) {
		var modalInstance = $modal.open({
			templateUrl : 'back/designColor/designColorPic.html',
			controller : 'designColorModal',
			size : size,
			backdrop : 'static',
			resolve : {
				selectrow : function() { return selectrow; },
				modaltype : function() { return 3; }
			}
		});

		modalInstance.result.then(function(selectedItem) {
			$scope.selected = selectedItem;
		}, function() {
			$log.info('Modal dismissed at: ' + new Date());
		});
	};


}])
// 模态框控制器
.controller('designColorModal', ['$scope', '$modalInstance', 'designColorService', 'selectrow', 'modaltype', 'FileUploader',
	function($scope, $modalInstance, designColorService, selectrow, modaltype, FileUploader) {

	if (modaltype == 0 || modaltype == 1) {
		designColorService.getColorList().then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					$scope.colors = jsonObj.colorInfos;
					if (jsonObj.colorInfos.length != 0) {
						if (modaltype == 0) {
							$scope.colorSelect = $scope.colors[0];
						} else {
							for (var i = $scope.colors.length - 1; i >= 0; i--) {
								if ($scope.colors[i].id == selectrow.colorId) {
									$scope.colorSelect = $scope.colors[i];
									break;
								}
							}
						}
					}
				} else {
					toastr.error("失败：" + result.data.Msg);
				}
			} else {
				toastr.error("失败：" + result.status + " " + result.statusText);
			}
		});


		///////////////////////////////////////////////////////////////////

		var uploader = $scope.uploader = new FileUploader({
			url: '/crm/designColor/uploadPic',
			queueLimit: 1
		});

		// FILTERS

		uploader.filters.push({
			name: 'imageFilter',
			fn: function(item /*{File|FileLikeObject}*/, options) {
				var type = '|' + item.type.slice(item.type.lastIndexOf('/') + 1) + '|';
				return '|jpg|png|jpeg|bmp|gif|'.indexOf(type) !== -1;
			}
		});

		// CALLBACKS

		uploader.onWhenAddingFileFailed = function(item /*{File|FileLikeObject}*/, filter, options) {
			console.info('onWhenAddingFileFailed', item, filter, options);
		};
		uploader.onAfterAddingFile = function(fileItem) {
			console.info('onAfterAddingFile', fileItem);
		};
		uploader.onAfterAddingAll = function(addedFileItems) {
			console.info('onAfterAddingAll', addedFileItems);
		};
		uploader.onBeforeUploadItem = function(item) {
			console.info('onBeforeUploadItem', item);
		};
		uploader.onProgressItem = function(fileItem, progress) {
			console.info('onProgressItem', fileItem, progress);
		};
		uploader.onProgressAll = function(progress) {
			console.info('onProgressAll', progress);
		};
		uploader.onSuccessItem = function(fileItem, response, status, headers) {
			// console.info('onSuccessItem', fileItem, response, status, headers);
			$scope.item.picUrl = response.Msg;
		};
		uploader.onErrorItem = function(fileItem, response, status, headers) {
			console.info('onErrorItem', fileItem, response, status, headers);
		};
		uploader.onCancelItem = function(fileItem, response, status, headers) {
			console.info('onCancelItem', fileItem, response, status, headers);
		};
		uploader.onCompleteItem = function(fileItem, response, status, headers) {
			console.info('onCompleteItem', fileItem, response, status, headers);
		};
		uploader.onCompleteAll = function() {
			console.info('onCompleteAll');
		};

		console.info('uploader', uploader);

		$scope.clearItems = function() { //重新选择文件时，清空队列，达到覆盖文件的效果
			uploader.clearQueue();
		}

		///////////////////////////////////////////////////////////////////
	}

	if (modaltype == 0) {
		$scope.item = {
			designId: selectrow.designId,
			designName: selectrow.designName,
			picUrl: "",
			note: ""
		};
	}

	if (modaltype == 3) {
		$scope.picUrl = selectrow.picUrl;
	}

	// 新增
	$scope.onAdd = function() {
		var designId = $scope.item.designId,
			colorId = $scope.colorSelect.id,
			picUrl = $scope.item.picUrl,
			note = $scope.item.note;
		designColorService.add(designId, colorId, picUrl, note).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					toastr.error("新增成功！！！");
					$modalInstance.close($scope.selected);
				} else {
					toastr.error("新增失败：" + result.data.Msg);
				}
			} else {
				toastr.error("新增失败：" + result.status + " " + result.statusText);
			}
		});
	};

	if (modaltype == 1) {
		$scope.item = {
			designId: selectrow.designId,
			designName: selectrow.designName,
			picUrl: selectrow.picUrl,
			note: selectrow.note
		};
	}

	// 编辑
	$scope.onEdit = function() {
		var designId = $scope.item.designId,
			colorId = $scope.colorSelect.id,
			picUrl = $scope.item.picUrl,
			note = $scope.item.note;
		designColorService.edit(designId, colorId, picUrl, note).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					toastr.error("编辑成功！！！");
					$modalInstance.close($scope.selected);
				} else {
					toastr.error("编辑失败：" + result.data.Msg);
				}
			} else {
				toastr.error("编辑失败：" + result.status + " " + result.statusText);
			}
		});
	};

	// 删除
	$scope.onDelete = function() {
		designColorService.remove(selectrow.id).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					toastr.error("删除成功！！！");
					$modalInstance.close($scope.selected);
				} else {
					toastr.error("删除失败：" + result.data.Msg);
				}
			} else {
				toastr.error("删除失败：" + result.status + " " + result.statusText);
			}
		});
	};

	// cancel click
	$scope.cancel = function() {
		$modalInstance.dismiss('cancel');
	};
}])
// 后台交互服务
.factory('designColorService', ['$q', '$filter', '$timeout', '$http', '$cookies', function ($q, $filter, $timeout, $http, $cookies) {
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


	function getColorList() {
		var url = '/crm/color/getList';
		var params = {
			'offset': 0,
			'size':   99999
		};
		return iHttp("getColorList", url, params);
	}


	function getList(offset, size, designId) {
		var url = '/crm/designColor/getList';
		var params = {
			'designId': designId,
			'offset': offset,
			'size':   size
		};
		return iHttp("getList", url, params);
	}

	function add(designId, colorId, picUrl, note) {
		var url = '/crm/designColor/add';
		var params = {
			'designId': designId,
			'colorId': colorId,
			'picUrl': picUrl,
			'note': note
		};
		return iHttp("add", url, params);
	}

	function edit(designId, colorId, picUrl, note) {
		var url = '/crm/designColor/edit';
		var params = {
			'designId': designId,
			'colorId': colorId,
			'picUrl': picUrl,
			'note': note
		};
		return iHttp("edit", url, params);
	}

	function remove(id) {
		var url = '/crm/designColor/delete';
		var params = {
			'id': id
		};
		return iHttp("remove", url, params);
	}

	// 接口函数实现 end
	///////////////////////////////////////////////////

	return {
		getColorList : getColorList,
		getList : getList,
		add : add,
		edit : edit,
		remove : remove
	};
}]);