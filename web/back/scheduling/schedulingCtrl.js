'use strict';

angular.module('yunyunApp')
// 视图控制器
.controller('schedulingCtrl', ['$scope', '$modal', '$log', 'schedulingService', function($scope, $modal, $log, schedulingService) {
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
		schedulingService.getList(ctrl.start, ctrl.number).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					ctrl.displayed = jsonObj.schedulingInfos;
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
		schedulingService.getList(ctrl.start, ctrl.number).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					ctrl.displayed = jsonObj.schedulingInfos;
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
			templateUrl : 'back/scheduling/schedulingAdd.html',
			controller : 'schedulingModal',
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
			templateUrl : 'back/scheduling/schedulingEdit.html',
			controller : 'schedulingModal',
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
			templateUrl : 'back/scheduling/schedulingDelete.html',
			controller : 'schedulingModal',
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
}])
// 模态框控制器
.controller('schedulingModal', ['$scope', '$modalInstance', 'schedulingService', 'selectrow', 'modaltype', function($scope, $modalInstance, schedulingService, selectrow, modaltype) {
	if (modaltype == 0 || modaltype == 1) {
		schedulingService.getDesignList().then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					$scope.designs = jsonObj.designInfos;
					if (jsonObj.designInfos.length != 0) {
						if (modaltype == 0) {
							$scope.designSelect = $scope.designs[0];
						} else {
							for (var i = $scope.designs.length - 1; i >= 0; i--) {
								if ($scope.designs[i].id == selectrow.designId) {
									$scope.designSelect = $scope.designs[i];
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
		schedulingService.getColorList().then(function (result) {
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
		schedulingService.getSupplierList().then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					$scope.suppliers = jsonObj.supplierInfos;
					if (jsonObj.supplierInfos.length != 0) {
						if (modaltype == 0) {
							$scope.supplierSelect = $scope.suppliers[0];
						} else {
							for (var i = $scope.suppliers.length - 1; i >= 0; i--) {
								if ($scope.suppliers[i].id == selectrow.supplierId) {
									$scope.supplierSelect = $scope.suppliers[i];
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

	// 新增
	$scope.onAdd = function() {
		var designId = $scope.designSelect.id,
			colorId = $scope.colorSelect.id,
			supplierId = $scope.supplierSelect.id,
			quantity = $scope.item.quantity,
			price = $scope.item.price,
			time = $scope.item.time,
			note = $scope.item.note;
		schedulingService.add(designId, colorId, supplierId, quantity, price, time, note).then(function (result) {
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
			quantity: selectrow.quantity,
			price: selectrow.price,
			time: selectrow.time,
			note: selectrow.note
		};
	}

	// 编辑
	$scope.onEdit = function() {
		var id = $scope.item.id,
			designId = $scope.designSelect.id,
			colorId = $scope.colorSelect.id,
			supplierId = $scope.supplierSelect.id,
			quantity = $scope.item.quantity,
			price = $scope.item.price,
			time = $scope.item.time,
			note = $scope.item.note;
		schedulingService.edit(id, designId, colorId, supplierId, quantity, price, time, note).then(function (result) {
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
		schedulingService.remove(selectrow.id).then(function (result) {
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
.factory('schedulingService', ['$q', '$filter', '$timeout', '$http', '$cookies', function ($q, $filter, $timeout, $http, $cookies) {
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

	function getDesignList() {
		var url = '/crm/design/getList';
		var params = {
			'offset': 0,
			'size':   99999
		};
		return iHttp("getDesignList", url, params);
	}

	function getColorList() {
		var url = '/crm/color/getList';
		var params = {
			'offset': 0,
			'size':   99999
		};
		return iHttp("getColorList", url, params);
	}

	function getSupplierList() {
		var url = '/crm/supplier/getList';
		var params = {
			'offset': 0,
			'size':   9999
		};
		return iHttp("getSupplierList", url, params);
	}

	function getList(offset, size) {
		var url = '/crm/scheduling/getList';
		var params = {
			'offset': offset,
			'size':   size
		};
		return iHttp("getList", url, params);
	}

	function add(designId, colorId, supplierId, quantity, price, time, note) {
		var url = '/crm/scheduling/add';
		var params = {
			'designId': designId,
			'colorId': colorId,
			'supplierId': supplierId,
			'quantity': quantity,
			'price': price,
			'time': time,
			'note': note
		};
		return iHttp("add", url, params);
	}

	function edit(id, designId, colorId, supplierId, quantity, price, time, note) {
		var url = '/crm/scheduling/edit';
		var params = {
			'id': id,
			'designId': designId,
			'colorId': colorId,
			'supplierId': supplierId,
			'quantity': quantity,
			'price': price,
			'time': time,
			'note': note
		};
		return iHttp("edit", url, params);
	}

	function remove(id) {
		var url = '/crm/scheduling/delete';
		var params = {
			'id': id
		};
		return iHttp("remove", url, params);
	}

	// 接口函数实现 end
	///////////////////////////////////////////////////

	return {
		getDesignList : getDesignList,
		getColorList : getColorList,
		getSupplierList : getSupplierList,
		getList : getList,
		add : add,
		edit : edit,
		remove : remove
	};
}]);