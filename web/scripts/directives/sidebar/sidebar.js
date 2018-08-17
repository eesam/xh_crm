'use strict';

// 1、根据当前页面路由自动调整样式
// 2、可折叠所有
angular.module('yunyunApp')
.directive('sidebar', ['Yyservice', '$cookies', function(Yyservice, $cookies) {
	return {
		templateUrl:'scripts/directives/sidebar/sidebar.html',
		restrict: 'E',
		replace: true,
		scope: {},
		link: function(scope, element, attrs) {
			// 平台管理员
			scope.systemAdminlevels = [
				{
					name        : '资源',
					collapse    : false,
					secondlevels: [
						{
							name   : '花型列表',
							uisref : 'back.design',
							icon   : 'bars'
						},
						{
							name   : '客户列表',
							uisref : 'back.customer',
							icon   : 'bars'
						},
						{
							name   : '供应商列表',
							uisref : 'back.supplier',
							icon   : 'bars'
						},
						{
							name   : '颜色列表',
							uisref : 'back.color',
							icon   : 'bars'
						},
					]
				},
				{
					name        : '工作',
					collapse    : false,
					secondlevels: [
						{
							name   : '新增入库单',
							uisref : 'back.inboundAdd',
							icon   : 'angle-right'
						},
						{
							name   : '新增出库单',
							uisref : 'back.outboundAdd',
							icon   : 'angle-right'
						},
						{
							name   : '入库',
							uisref : 'back.inbound',
							icon   : 'angle-right'
						},
						{
							name   : '出库',
							uisref : 'back.outbound',
							icon   : 'angle-right'
						},
						{
							name   : '排单',
							uisref : 'back.scheduling',
							icon   : 'angle-right'
						}
					]
				}
			];
			scope.firstlevels = scope.systemAdminlevels;


			// 检查该节点是否没有第三级目录 并且具有href属性
			scope.checkNodeA = function(node) {
				if ( (typeof(node.thirdlevels) === 'undefined') || (node.thirdlevels.length === 0) ) {
					if (typeof(node.href) !== 'undefined') {
						return true;
					}
				}
				return false;
			};
			// 检查该节点是否没有第三级目录 并且不具有href属性
			scope.checkNodeB = function(node) {
				if ( (typeof(node.thirdlevels) === 'undefined') || (node.thirdlevels.length === 0) ) {
					if (typeof(node.href) === 'undefined') {
						return true;
					}
				}
				return false;
			};

			// 检查该节点是否具有第三级目录 并且不具有uisref/href属性
			scope.checkNodeC = function(node) {
				if ( (typeof(node.thirdlevels) === 'undefined') || (node.thirdlevels.length === 0) ) {
					return false;
				}
				if ( (typeof(node.uisref) !== 'undefined') || (typeof(node.href) !== 'undefined') ) {
					return false;
				}
				return true;
			};

			// 检查该节点是否具有第三级目录 并且具有href属性
			scope.checkNodeD = function(node) {
				if ( (typeof(node.thirdlevels) !== 'undefined') && (node.thirdlevels.length !== 0) && (typeof(node.href) !== 'undefined') ) {
					return true;
				}
				return false;
			};

			// 检查该节点是否具有第三级目录 并且具有uisref属性
			scope.checkNodeE = function(node) {
				if ( (typeof(node.thirdlevels) !== 'undefined') && (node.thirdlevels.length !== 0) && (typeof(node.uisref) !== 'undefined') ) {
					return true;
				}
				return false;
			};

			// 检查该节点是否具有href属性
			scope.hasHrefAttr = function(node) {
				if (typeof(node.href) === 'undefined') {
					return false;
				}
				return true;
			};


			scope.sidebarActive = '';
			scope.onSelect = function(tag) {
				// console.debug("sidebarActive:" + scope.sidebarActive + " tag:" + tag);
				scope.sidebarActive = tag;
			};

			scope.secondlevelToggleSelect = function(firstindex, secondindex) {
				scope.secondlevelToggle(firstindex, secondindex);
				scope.onSelect(firstindex + '-' + secondindex);
			};

			// 一级目录展开折叠切换
			scope.firstlevelToggle = function(index) {
				scope.firstlevels[index].collapse = !scope.firstlevels[index].collapse;
			};

			// 二级目录展开折叠切换
			scope.secondlevelToggle = function(firstindex, secondindex) {
				scope.firstlevels[firstindex].secondlevels[secondindex].expand = !scope.firstlevels[firstindex].secondlevels[secondindex].expand;
			};

			// 查找tag项的firstindex/secondindex/thirdindex
			function findByTag(tag) {
				for (var i = 0; i < scope.firstlevels.length; i++) {
					var seclevels = scope.firstlevels[i].secondlevels;
					for (var j = 0; j < seclevels.length; j++) {
						var temptag = seclevels[j].tag;
						if ( (typeof(temptag) !== 'undefined') && (temptag === tag) ) {
							return [i, j];
						}
						if (typeof(seclevels[j].thirdlevels) !== 'undefined') {
							for (var k = 0; k < seclevels[j].thirdlevels.length; k++) {
								temptag = seclevels[j].thirdlevels[k].tag;
								if ( (typeof(temptag) !== 'undefined') && (temptag === tag) ) {
									return [i, j, k];
								}
							}
						}
					}
				}
			}

			// 改变侧边栏初始选择项样式
			// if (yy_sidebarActiveVar === '') {
			// 	yy_sidebarActiveVar = 'console';
			// }
			// var index = findByTag(yy_sidebarActiveVar);
			// if (index.length == 3) { // 三级目录
			// 	scope.onSelect(index[0] + '-' + index[1] + '-' + index[2]);
			// 	scope.secondlevelToggle(index[0], index[1]);
			// } else {
			// 	scope.onSelect(index[0] + '-' + index[1]);
			// }
			// yy_sidebarActiveVar = '';
		}
	};
}]);
