// 后台模块路由配置

'use strict';

/**
 * title         网页title
 * route         路由名，如'back.myConsole'				注：用于<a ui-sref="back.myConsole">
 * url           浏览器显示路径，如'/myConsole'	        注：省略字段
 * templateUrl   对应html文件
 * template      <div ui-view></div>
 * hasOwnControl 是否拥有对应的ctrl，默认为true
 * abstract      禁用路由，默认为false
 * lazyLoad      按需加载其他组件，默认为空数组'[]'
 * data          数据加载完才进入router
 */
var yy_route_back = [
	{
		route			: "",
		templateUrl		: 'main.html',
		hasOwnControl	: false,
		abstract		: true,
		idata			: "idata",
		lazyLoad		: [
			{
				name: 'yunyunApp',
				files: [
					// header头部栏
					'scripts/directives/header/header.js',
					// iframe内容高度
					// 'scripts/directives/hresize/iframeHeight.js',
					// 左侧导航栏
					'scripts/directives/sidebar/sidebar.js',
					//中间主题块
					'scripts/directives/content/content.js',
					'scripts/directives/filemodel/filemodel.js',
					'scripts/directives/smart-table-plugin/csSelect.js',
					'scripts/directives/smart-table-plugin/lrInfiniteScroll.js',
					'scripts/directives/ng-thumb/ngThumb.js'
				]
			},
			// {
			// 	name: 'ngAnimate',
			// 	files: ['bower_components/angular/angular-animate.min.js']
			// },
			{
				name: 'ui.select',
				files: ['bower_components/ui-select/dist/select.min.css','bower_components/ui-select/dist/select.min.js']
			},
			{
				name: 'ngSanitize',
				files: ['bower_components/angular/angular-sanitize.min.js']
			},
			// {
			// 	name: 'treeControl',
			// 	files: [
			// 		'bower_components/angular-tree-control/angular-tree-control.js',
			// 		'bower_components/angular-tree-control/css/tree-control.css',
			// 		'bower_components/angular-tree-control/css/tree-control-attribute.css'
			// 	]
			// },
			{
				name: 'smart-table',
				files: [
					'bower_components/angular-smart-table/dist/smart-table.css',
					'bower_components/angular-smart-table/dist/smart-table.min.js'
				]
			},
			{
				name: 'angularFileUpload',
				files: [
					'bower_components/angular-file-upload/angular-file-upload.min.js'
				]
			},
			// {
			// 	name: 'chart.js',
			// 	files: [
			// 		'bower_components/angular-chart.js/dist/angular-chart.min.js',
			// 		'bower_components/angular-chart.js/dist/angular-chart.css'
			// 	]
			// }
		]
	},
	{
		title			: "花型列表",
		route			: "design",
		templateUrl		: "design/design.html"
	},
	{
		title			: "颜色列表",
		route			: "color",
		templateUrl		: "color/color.html"
	},
	{
		title			: "花型颜色",
		route			: "designColor",
		templateUrl		: "designColor/designColor.html",
		params          : {"designId": null, "designName": null}
	},
	{
		title			: "客户列表",
		route			: "customer",
		templateUrl		: "customer/customer.html"
	},
	{
		title			: "供应商列表",
		route			: "supplier",
		templateUrl		: "supplier/supplier.html"
	},
	{
		title			: "新增入库单",
		route			: "inboundAdd",
		templateUrl		: "inbound/inboundAdd.html"
	},
	{
		title			: "新增出库单",
		route			: "outboundAdd",
		templateUrl		: "outbound/outboundAdd.html"
	},
	{
		title			: "入库单",
		route			: "inbound",
		templateUrl		: "inbound/inbound.html"
	},
	{
		title			: "出库单",
		route			: "outbound",
		templateUrl		: "outbound/outbound.html"
	},
	{
		title			: "排单",
		route			: "scheduling",
		templateUrl		: "scheduling/scheduling.html"
	}
];
