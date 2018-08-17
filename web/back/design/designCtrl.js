'use strict';

angular.module('yunyunApp')
// 视图控制器
.controller('designCtrl', ['$scope', '$modal', '$log', '$state', 'designService', function($scope, $modal, $log, $state, designService) {
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
		designService.getList(ctrl.start, ctrl.number).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					ctrl.displayed = jsonObj.designInfos;
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
		designService.getList(ctrl.start, ctrl.number).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					ctrl.displayed = jsonObj.designInfos;
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
		var modalInstance = $modal.open({
			templateUrl : 'back/design/designAdd.html',
			controller : 'designModal',
			size : size,
			backdrop : 'static',
			resolve : {
				selectrow : function() {},
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
			templateUrl : 'back/design/designEdit.html',
			controller : 'designModal',
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
			templateUrl : 'back/design/designDelete.html',
			controller : 'designModal',
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

	ctrl.openUploadPic = function(size, selectrow) {
		$state.go('back.designColor', {designId: selectrow.id, designName: selectrow.name});
	};

}])
// 模态框控制器
.controller('designModal', ['$scope', '$modalInstance', 'designService', 'selectrow', 'modaltype', function($scope, $modalInstance, designService, selectrow, modaltype) {
	// 新增
	$scope.onAdd = function() {
		var name = $scope.item.name,
			note = $scope.item.note;
		designService.add(name, note).then(function (result) {
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
			id: selectrow.id,
			name: selectrow.name,
			quantity: selectrow.quantity,
			note: selectrow.note
		};
	}

	// 编辑
	$scope.onEdit = function() {
		var id = $scope.item.id,
			name = $scope.item.name,
			quantity = $scope.item.quantity,
			note = $scope.item.note;
		designService.edit(id, name, quantity, note).then(function (result) {
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
		designService.remove(selectrow.id).then(function (result) {
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

	if (modaltype == 3) {
		designService.getColorList().then(function (result) {
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
	}
}])
// 后台交互服务
.factory('designService', ['$q', '$filter', '$timeout', '$http', '$cookies', function ($q, $filter, $timeout, $http, $cookies) {
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

	function getList(offset, size) {
		var url = '/crm/design/getList';
		var params = {
			'offset': offset,
			'size':   size
		};
		return iHttp("getList", url, params);
	}

	function getColorList() {
		var url = '/crm/color/getList';
		var params = {
			'offset': 0,
			'size':   99999
		};
		return iHttp("getColorList", url, params);
	}

	function add(name, note) {
		var url = '/crm/design/add';
		var params = {
			'name': name,
			'note': note
		};
		return iHttp("add", url, params);
	}

	function edit(id, name, quantity, note) {
		var url = '/crm/design/edit';
		var params = {
			'id': id,
			'name': name,
			'quantity': quantity,
			'note': note
		};
		return iHttp("edit", url, params);
	}

	function remove(id) {
		var url = '/crm/design/delete';
		var params = {
			'id': id
		};
		return iHttp("remove", url, params);
	}

	// 接口函数实现 end
	///////////////////////////////////////////////////

	return {
		getList : getList,
		getColorList: getColorList,
		add : add,
		edit : edit,
		remove : remove
	};
}]);