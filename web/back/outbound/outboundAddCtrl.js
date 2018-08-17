'use strict';

angular.module('yunyunApp')
// 视图控制器
.controller('outboundAddCtrl', ['$scope', '$log', '$state', 'outboundAddService', function($scope, $log, $state, outboundAddService) {
	$scope.designSelect = [];
	$scope.colorSelect = [];
	$scope.customerSelect = [];
	$scope.note = [];
	outboundAddService.getDesignList().then(function (result) {
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
	outboundAddService.getColorList().then(function (result) {
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
	outboundAddService.getCustomerList().then(function (result) {
		if (result.status == 200) {
			if (result.data.Code === 0) {
				var jsonObj = angular.fromJson(result.data.Msg);
				$scope.customers = jsonObj.customerInfos;
				if (jsonObj.customerInfos.length != 0) {
					if ($scope.customerSelect.length == 0) {
						$scope.customerSelect.push($scope.customers[0]);
					}
				}
			} else {
				toastr.error("失败：" + result.data.Msg);
			}
		} else {
			toastr.error("失败：" + result.status + " " + result.statusText);
		}
	});

	////////////////////////////////////////
	$scope.designColors = [{"0":null}]
	$scope.addDesignColor = function() {
		$scope.designColors.push({"0":null});

		$scope.designSelect.push($scope.designs[0]);
		$scope.colorSelect.push($scope.colors[0]);
		$scope.customerSelect.push($scope.customers[0]);
		$scope.outboundCloths.push([]);
	};
	$scope.deleteDesignColor = function(index) {
		$scope.designColors.pop();
		$scope.designSelect.splice(index, 1);
		$scope.colorSelect.splice(index, 1);
		$scope.customerSelect.splice(index, 1);
		$scope.note.splice(index, 1);
		$scope.outboundCloths.splice(index, 1);
	};
	////////////////////////////////////////

	$scope.outboundCloths = [[]];
	$scope.isCheck = [];
	$scope.updateAllSelection = function(index) {
		for (var i = 0; i <= $scope.outboundCloths[index].length - 1; i++) {
			$scope.outboundCloths[index][i].isCheck = $scope.isCheck[index];
		}
	};

	$scope.designSelectChange = function(index) {
		console.log("designSelect change");
		$scope.queryStock(index);
	};

	$scope.colorSelectChange = function(index) {
		console.log("colorSelect change");
		$scope.queryStock(index);
	};

	$scope.queryStock = function(index) {
		$scope.outboundCloths[index].splice(0, $scope.outboundCloths[index].length);
		var designId = $scope.designSelect[index].id,
			colorId = $scope.colorSelect[index].id;
		outboundAddService.getInboundClothList(designId, colorId).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					var jsonObj = angular.fromJson(result.data.Msg);
					if (jsonObj.inboundClothInfos.length == 0) {
						toastr.info("暂无数据");
					}
					for (var i = 0; i <= jsonObj.inboundClothInfos.length - 1; i++) {
						var ic = {
							"inboundClothId": jsonObj.inboundClothInfos[i].id,
							"inboundPrice": jsonObj.inboundClothInfos[i].price,
							"remainQuantity": jsonObj.inboundClothInfos[i].remainQuantity,
							"outboundPrice": jsonObj.inboundClothInfos[i].price,
							"outboundQuantity": jsonObj.inboundClothInfos[i].remainQuantity,
							"isCheck": false
						};
						$scope.outboundCloths[index].push(ic);
					}
				} else {
					toastr.error("失败：" + result.data.Msg);
				}
			} else {
				toastr.error("失败：" + result.status + " " + result.statusText);
			}
		});
	};

	$scope.onAdd = function() {
		var outboundCloths = [];
		for (var i = 0; i <= $scope.designSelect.length - 1; i++) {
			var cloths = [];
			cloths.splice(0, cloths.length);
			var inboundClothsLen = $scope.outboundCloths[i].length;
			for (var j = 0; j <= $scope.outboundCloths[i].length - 1; j++) {
				if ($scope.outboundCloths[i][j].isCheck == false ||
					$scope.outboundCloths[i][j].remainQuantity <= 0 ||
					$scope.outboundCloths[i][j].outboundPrice < 0 ||
					$scope.outboundCloths[i][j].outboundQuantity > $scope.outboundCloths[i][j].remainQuantity) {
					inboundClothsLen--;
				} else {
					var cloth = {
						"inboundClothId": $scope.outboundCloths[i][j].inboundClothId,
						"outboundPrice": $scope.outboundCloths[i][j].outboundPrice,
						"outboundQuantity": $scope.outboundCloths[i][j].outboundQuantity
					};
					cloths.push(cloth);
				}
			}
			if (inboundClothsLen <= 0) continue;

			var outboundCloth = {
				// "designId": $scope.designSelect[i].id,
				// "colorId": $scope.colorSelect[i].id,
				"customerId": $scope.customerSelect[i].id,
				"note": $scope.note[i],
				"cloths": cloths
			};

			outboundCloths.push(outboundCloth);
		}

		if (outboundCloths.length <= 0) {
			toastr.info("请输入有效的数据！！！");
			return;
		};
		outboundAddService.add(outboundCloths).then(function (result) {
			if (result.status == 200) {
				if (result.data.Code === 0) {
					toastr.error("新增成功！！！");
					$state.reload('back.outboundAdd');
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
.factory('outboundAddService', ['$q', '$filter', '$timeout', '$http', '$cookies', function ($q, $filter, $timeout, $http, $cookies) {
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
			'size':   99999
		};
		return iHttp("getCustomerList", url, params);
	}

	function getInboundClothList(designId, colorId) {
		var url = '/crm/inboundCloth/getList';
		var params = {
			'designId': designId,
			'colorId': colorId
		};
		return iHttp("add", url, params);
	}

	function add(outboundCloths) {
		var url = '/crm/outboundCloth/add';
		var params = {
			'outboundCloths': outboundCloths
		};
		return iHttp("add", url, params);
	}

	// 接口函数实现 end
	///////////////////////////////////////////////////

	return {
		getDesignList : getDesignList,
		getColorList : getColorList,
		getCustomerList : getCustomerList,
		getInboundClothList : getInboundClothList,
		add : add
	};
}]);