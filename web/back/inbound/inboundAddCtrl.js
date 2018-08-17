'use strict';

angular.module('yunyunApp')
// 视图控制器
.controller('inboundAddCtrl', ['$scope', '$log', '$state', 'inboundAddService', function($scope, $log, $state, inboundAddService) {
	$scope.designSelect = [];
	$scope.colorSelect = [];
	$scope.supplierSelect = [];
	$scope.price = [];
	$scope.note = [];
	inboundAddService.getDesignList().then(function (result) {
		if (result.status == 200) {
			if (result.data.Code === 0) {
				var jsonObj = angular.fromJson(result.data.Msg);
				$scope.designs = jsonObj.designInfos;
				if (jsonObj.designInfos.length != 0) {
					if ($scope.designSelect.length == 0) {
						$scope.designSelect.push($scope.designs[0]);
					}
				}
			} else {
				toastr.error("失败：" + result.data.Msg);
			}
		} else {
			toastr.error("失败：" + result.status + " " + result.statusText);
		}
	});
	inboundAddService.getColorList().then(function (result) {
		if (result.status == 200) {
			if (result.data.Code === 0) {
				var jsonObj = angular.fromJson(result.data.Msg);
				$scope.colors = jsonObj.colorInfos;
				if (jsonObj.colorInfos.length != 0) {
					if ($scope.colorSelect.length == 0) {
						$scope.colorSelect.push($scope.colors[0]);
					}
				}
			} else {
				toastr.error("失败：" + result.data.Msg);
			}
		} else {
			toastr.error("失败：" + result.status + " " + result.statusText);
		}
	});
	inboundAddService.getSupplierList().then(function (result) {
		if (result.status == 200) {
			if (result.data.Code === 0) {
				var jsonObj = angular.fromJson(result.data.Msg);
				$scope.suppliers = jsonObj.supplierInfos;
				if (jsonObj.supplierInfos.length != 0) {
					if ($scope.supplierSelect.length == 0) {
						$scope.supplierSelect.push($scope.suppliers[0]);
					}
				}
			} else {
				toastr.error("失败：" + result.data.Msg);
			}
		} else {
			toastr.error("失败：" + result.status + " " + result.statusText);
		}
	});

	$scope.designColors = [{"0":null}]
	$scope.addDesignColor = function() {
		$scope.quantitys.push([{"0":null}]);
		$scope.designColors.push({"0":null});

		$scope.designSelect.push($scope.designs[0]);
		$scope.colorSelect.push($scope.colors[0]);
		$scope.supplierSelect.push($scope.suppliers[0]);
	};
	$scope.deleteDesignColor = function(index) {
		$scope.designColors.pop();
		$scope.designSelect.splice(index, 1);
		$scope.colorSelect.splice(index, 1);
		$scope.supplierSelect.splice(index, 1);
		$scope.price.splice(index, 1);
		$scope.note.splice(index, 1);
		$scope.quantitys.splice(index, 1);
		$scope.inputQuantitys.splice(index, 1);
	};

	$scope.quantitys = [[{"0":null}]];
	$scope.inputQuantitys = [[]];
	$scope.addQuantity = function(index) {
		$scope.quantitys[index].push({"0":null});
		$scope.inputQuantitys.push([]);
		$scope.inputQuantitys[index].push(null);
	};
	$scope.deleteQuantity = function(parentIndex, index) {
		$scope.quantitys[parentIndex].pop();
		$scope.inputQuantitys[parentIndex].splice(index, 1);
	};

	$scope.onAdd = function() {
		var inboundCloths = [];
		for (var i = 0; i <= $scope.designSelect.length - 1; i++) {
			if ($scope.price[i] < 0) continue;

			var inputQuantitysLen = $scope.inputQuantitys[i].length;
			for (var j = 0; j <= $scope.inputQuantitys[i].length - 1; j++) {
				if ($scope.inputQuantitys[i][j] <= 0) {
					inputQuantitysLen--;
				}
			}
			if (inputQuantitysLen <= 0) continue;

			var cloth = {
				"designId": $scope.designSelect[i].id,
				"colorId": $scope.colorSelect[i].id,
				"supplierId": $scope.supplierSelect[i].id,
				"quantitys": $scope.inputQuantitys[i],
				"price": $scope.price[i],
				"note": $scope.note[i]
			};
			inboundCloths.push(cloth);
		}
		if (inboundCloths.length <= 0) {
			toastr.info("请输入有效的数据！！！");
			return;
		};
		inboundAddService.add(inboundCloths).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					toastr.error("新增成功！！！");
					$state.reload('back.inboundAdd');
				} else {
					toastr.error("新增失败：" + result.data.Msg);
				}
			} else {
				toastr.error("新增失败：" + result.status + " " + result.statusText);
			}
		});
	};
}])
// 后台交互服务
.factory('inboundAddService', ['$q', '$filter', '$timeout', '$http', '$cookies', function ($q, $filter, $timeout, $http, $cookies) {
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
			'size':   99999
		};
		return iHttp("getSupplierList", url, params);
	}

	function add(inboundCloths) {
		var url = '/crm/inboundCloth/add';
		var params = {
			'inboundCloths': inboundCloths
		};
		return iHttp("add", url, params);
	}

	// 接口函数实现 end
	///////////////////////////////////////////////////

	return {
		getDesignList : getDesignList,
		getColorList : getColorList,
		getSupplierList : getSupplierList,
		add : add
	};
}]);