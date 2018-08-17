'use strict';

angular.module('yunyunApp')
// 视图控制器
.controller('outboundCtrl', ['$scope', '$modal', '$log', 'outboundService', function($scope, $modal, $log, outboundService) {
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
		outboundService.getList(ctrl.start, ctrl.number).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					ctrl.displayed = jsonObj.outboundInfos;
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
		outboundService.getList(ctrl.start, ctrl.number).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					ctrl.displayed = jsonObj.outboundInfos;
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
			templateUrl : 'back/outbound/outboundAdd.html',
			controller : 'outboundModal',
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
			templateUrl : 'back/outbound/outboundEdit.html',
			controller : 'outboundModal',
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
			templateUrl : 'back/outbound/outboundDelete.html',
			controller : 'outboundModal',
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
.controller('outboundModal', ['$scope', '$modalInstance', 'outboundService', 'selectrow', 'modaltype', function($scope, $modalInstance, outboundService, selectrow, modaltype) {
	if (modaltype == 0 || modaltype == 1) {
		outboundService.getDesignList().then(function (result) {
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
		outboundService.getColorList().then(function (result) {
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
		outboundService.getCustomerList().then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					$scope.customers = jsonObj.customerInfos;
					if (jsonObj.customerInfos.length != 0) {
						if (modaltype == 0) {
							$scope.customerSelect = $scope.customers[0];
						} else {
							for (var i = $scope.customers.length - 1; i >= 0; i--) {
								if ($scope.customers[i].id == selectrow.customerId) {
									$scope.customerSelect = $scope.customers[i];
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
			customerId = $scope.customerSelect.id,
			quantity = $scope.item.quantity,
			price = $scope.item.price,
			time = $scope.item.time,
			note = $scope.item.note;
		outboundService.add(designId, colorId, customerId, quantity, price, time, note).then(function (result) {
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
			customerId = $scope.customerSelect.id,
			quantity = $scope.item.quantity,
			price = $scope.item.price,
			time = $scope.item.time,
			note = $scope.item.note;
		outboundService.edit(id, designId, colorId, customerId, quantity, price, time, note).then(function (result) {
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
		outboundService.remove(selectrow.id).then(function (result) {
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
.factory('outboundService', ['$q', '$filter', '$timeout', '$http', '$cookies', function ($q, $filter, $timeout, $http, $cookies) {
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

	function getCustomerList() {
		var url = '/crm/customer/getList';
		var params = {
			'offset': 0,
			'size':   9999
		};
		return iHttp("getCustomerList", url, params);
	}

	function getList(offset, size) {
		var url = '/crm/outbound/getList';
		var params = {
			'offset': offset,
			'size':   size
		};
		return iHttp("getList", url, params);
	}

	function add(designId, colorId, customerId, quantity, price, time, note) {
		var url = '/crm/outbound/add';
		var params = {
			'designId': designId,
			'colorId': colorId,
			'customerId': customerId,
			'quantity': quantity,
			'price': price,
			'time': time,
			'note': note
		};
		return iHttp("add", url, params);
	}

	function edit(id, designId, colorId, customerId, quantity, price, time, note) {
		var url = '/crm/outbound/edit';
		var params = {
			'id': id,
			'designId': designId,
			'colorId': colorId,
			'customerId': customerId,
			'quantity': quantity,
			'price': price,
			'time': time,
			'note': note
		};
		return iHttp("edit", url, params);
	}

	function remove(id) {
		var url = '/crm/outbound/delete';
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
		getCustomerList : getCustomerList,
		getList : getList,
		add : add,
		edit : edit,
		remove : remove
	};
}]);